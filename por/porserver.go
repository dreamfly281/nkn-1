package por

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/core/transaction"
	"github.com/nknorg/nkn/util/log"
	"github.com/nknorg/nkn/vault"
)

const (
	sigChainTxnCacheExpiration      = 300 * time.Second
	sigChainTxnCacheCleanupInterval = 10 * time.Second
)

type PorServer struct {
	account          *vault.Account
	id               []byte
	sigChainTxnCache common.Cache

	sync.RWMutex
	miningPorPackage map[uint32]*PorPackage
}

var porServer *PorServer

func NewPorServer(account *vault.Account, id []byte) *PorServer {
	ps := &PorServer{
		account:          account,
		id:               id,
		sigChainTxnCache: common.NewGoCache(sigChainTxnCacheExpiration, sigChainTxnCacheCleanupInterval),
		miningPorPackage: make(map[uint32]*PorPackage),
	}
	return ps
}

func InitPorServer(account *vault.Account, id []byte) error {
	if porServer != nil {
		return errors.New("PorServer already initialized")
	}
	if id == nil || len(id) == 0 {
		return errors.New("ID is empty")
	}
	porServer = NewPorServer(account, id)
	return nil
}

func GetPorServer() *PorServer {
	if porServer == nil {
		log.Error("PorServer not initialized")
		panic("PorServer not initialized")
	}
	return porServer
}

func (ps *PorServer) Sign(sc *SigChain, nextPubkey []byte, mining bool) error {
	dcPk, err := ps.account.PubKey().EncodePoint(true)
	if err != nil {
		log.Error("Get account public key error:", err)
		return err
	}

	nxPk, err := sc.GetLastPubkey()
	if err != nil {
		log.Error("Get last public key error:", err)
		return err
	}

	if !common.IsEqualBytes(dcPk, nxPk) {
		return errors.New("it's not the right signer")
	}

	err = sc.Sign(ps.id, nextPubkey, mining, ps.account)
	if err != nil {
		log.Error("Signature chain signing error:", err)
		return err
	}

	return nil
}

func (ps *PorServer) Verify(sc *SigChain) error {
	if err := sc.Verify(); err != nil {
		return err
	}

	return nil
}

// func (ps *PorServer) CreateSigChain(dataSize uint32, dataHash, blockHash *common.Uint256, srcID,
// 	destPubkey, nextPubkey []byte, mining bool) (*SigChain, error) {
// 	return NewSigChain(ps.account, dataSize, dataHash[:], blockHash[:], srcID, destPubkey, nextPubkey, mining)
// }

func (ps *PorServer) CreateSigChainForClient(dataSize uint32, dataHash, blockHash *common.Uint256, srcID,
	srcPubkey, destPubkey, signature []byte, sigAlgo SigAlgo) (*SigChain, error) {
	pubKey, err := ps.account.PubKey().EncodePoint(true)
	if err != nil {
		log.Error("Get account public key error:", err)
		return nil, err
	}
	sigChain, err := NewSigChainWithSignature(dataSize, dataHash[:], blockHash[:],
		srcID, srcPubkey, destPubkey, pubKey, signature, sigAlgo, false)
	if err != nil {
		log.Error("New signature chain with signature error:", err)
		return nil, err
	}
	return sigChain, nil
}

func (ps *PorServer) IsFinal(sc *SigChain) bool {
	return sc.IsFinal()
}

func (ps *PorServer) IsSatisfyThreshold() bool {
	//TODO need to add codes
	return true
}

func (ps *PorServer) GetSignature(sc *SigChain) ([]byte, error) {
	return sc.GetSignature()
}

func (ps *PorServer) LenOfSigChain(sc *SigChain) int {
	return sc.Length()
}

func (ps *PorServer) GetMiningSigChain(height uint32, timeOut bool) (*SigChain, *transaction.Transaction, error) {
	var porPkg *PorPackage
	ps.RLock()
	defer ps.RUnlock()

	if timeOut == false {
		porPkg = ps.miningPorPackage[height]
	} else {
		// Fixme the inital height should be height - 1?
		for h := height; h < (height + SigChainMiningHeightOffset + SigChainBlockHeightOffset); h++ {
			porPkg = ps.miningPorPackage[h]
			if porPkg != nil {
				log.Infof("Timeout and choose the proposal SigChain from %d", h)
				break
			}
		}
	}

	if porPkg == nil {
			return nil, nil, nil
	}

	v, ok := ps.sigChainTxnCache.Get(porPkg.TxnHash);
	if !ok {
		return nil, nil, fmt.Errorf("sigchain txn %s not found", porPkg.TxnHash)
	}

	txn, ok := v.(*transaction.Transaction)
	if !ok {
		return nil, nil, fmt.Errorf("convert to sigchain txn %s error", porPkg.TxnHash)
	}

	return porPkg.GetSigChain(), txn, nil
}

func (ps *PorServer) AddSigChainFromTxn(txn *transaction.Transaction, currHgt uint32) (bool, error) {
	porPkg, err := NewPorPackage(txn)
	if err != nil {
		return false, err
	}

	voteHgt := porPkg.GetVoteHeight()
	// FixME the definination of 2?
	if voteHgt < currHgt + 2 {
		return false, fmt.Errorf("sigchain vote for height %d is less than %d", voteHgt, currHgt)
	}

	ps.Lock()
	defer ps.Unlock()

	// TODO add aging time for signature chain to avoid the dead node send a txn without proposal block
	if ps.miningPorPackage[voteHgt] != nil && bytes.Compare(porPkg.SigHash, ps.miningPorPackage[voteHgt].SigHash) >= 0 {
		return false, nil
	}

	err = ps.sigChainTxnCache.Add(porPkg.TxnHash, txn)
	if err != nil {
		return false, err
	}

	ps.miningPorPackage[voteHgt] = porPkg

	return true, nil
}

func (ps *PorServer) GetThreshold() common.Uint256 {
	//TODO get from block
	return common.Uint256{}
}

func (ps *PorServer) UpdateThreshold() common.Uint256 {
	//TODO used for new block
	return common.Uint256{}
}
