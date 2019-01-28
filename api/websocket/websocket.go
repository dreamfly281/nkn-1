package websocket

import (
	"bytes"
	"encoding/json"

	"github.com/nknorg/nkn/api/common"
	"github.com/nknorg/nkn/api/websocket/server"
	. "github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/core"
	"github.com/nknorg/nkn/events"
	"github.com/nknorg/nkn/net/node"
	"github.com/nknorg/nkn/pb"
	"github.com/nknorg/nkn/types"
	. "github.com/nknorg/nkn/util/config"
	"github.com/nknorg/nkn/vault"
)

var ws *server.WsServer

var (
	pushBlockFlag    bool = true
	pushRawBlockFlag bool = false
	pushBlockTxsFlag bool = false
)

func NewServer(localNode *node.LocalNode, w vault.Wallet) *server.WsServer {
	//	common.SetNode(n)
	core.DefaultLedger.Blockchain.BCEvents.Subscribe(events.EventBlockPersistCompleted, SendBlock2WSclient)
	ws = server.InitWsServer(localNode, w)
	return ws
}

func SendBlock2WSclient(v interface{}) {
	n, err := GetServer().GetNetNode()
	if err == nil && n.GetSyncState() == pb.PersistFinished {
		go func() {
			PushSigChainBlockHash(v)
		}()
		if Parameters.HttpWsPort != 0 && pushBlockFlag {
			go func() {
				PushBlock(v)
			}()
		}
		if Parameters.HttpWsPort != 0 && pushBlockTxsFlag {
			go func() {
				PushBlockTransactions(v)
			}()
		}
	}
}

func GetWsPushBlockFlag() bool {
	return pushBlockFlag
}

func SetWsPushBlockFlag(b bool) {
	pushBlockFlag = b
}

func GetPushRawBlockFlag() bool {
	return pushRawBlockFlag
}

func SetPushRawBlockFlag(b bool) {
	pushRawBlockFlag = b
}

func GetPushBlockTxsFlag() bool {
	return pushBlockTxsFlag
}

func SetPushBlockTxsFlag(b bool) {
	pushBlockTxsFlag = b
}

func SetTxHashMap(txhash string, sessionid string) {
	if ws == nil {
		return
	}
	ws.SetTxHashMap(txhash, sessionid)
}

func PushBlock(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(common.SUCCESS)
	if block, ok := v.(*types.Block); ok {
		if pushRawBlockFlag {
			w := bytes.NewBuffer(nil)
			block.Serialize(w)
			resp["Result"] = BytesToHexString(w.Bytes())
		} else {
			info, _ := block.MarshalJson()
			var x interface{}
			json.Unmarshal(info, &x)
			resp["Result"] = x
		}
		resp["Action"] = "sendRawBlock"
		ws.PushResult(resp)
	}
}

func PushBlockTransactions(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(common.SUCCESS)
	if block, ok := v.(*types.Block); ok {
		if pushBlockTxsFlag {
			resp["Result"] = common.GetBlockTransactions(block)
		}
		resp["Action"] = "sendblocktransactions"
		ws.PushResult(resp)
	}
}

func PushSigChainBlockHash(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(common.SUCCESS)
	if block, ok := v.(*types.Block); ok {
		resp["Action"] = "updateSigChainBlockHash"
		//resp["Result"] = common.GetBlockInfo(block).BlockData.PrevBlockHash
		resp["Result"] = BytesToHexString(block.Header.UnsignedHeader.PrevBlockHash)
		ws.PushResult(resp)
	}
}

func GetServer() *server.WsServer {
	return ws
}
