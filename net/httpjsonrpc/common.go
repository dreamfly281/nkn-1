package httpjsonrpc

import (
	. "DNA/common"
	"DNA/consensus/dbft"
	. "DNA/core/transaction"
	tx "DNA/core/transaction"
	. "DNA/net/protocol"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func init() {
	mainMux.m = make(map[string]func([]interface{}) map[string]interface{})
}

//an instance of the multiplexer
var mainMux ServeMux
var node Noder
var dBFT *dbft.DbftService

//multiplexer that keeps track of every function to be called on specific rpc call
type ServeMux struct {
	sync.RWMutex
	m               map[string]func([]interface{}) map[string]interface{}
	defaultFunction func(http.ResponseWriter, *http.Request)
}

type TxAttributeInfo struct {
	Usage TransactionAttributeUsage
	Date  string
	Size  uint32
}

type UTXOTxInputInfo struct {
	ReferTxID          string
	ReferTxOutputIndex uint16
}

type BalanceTxInputInfo struct {
	AssetID     string
	Value       Fixed64
	ProgramHash string
}

type TxoutputInfo struct {
	AssetID     string
	Value       Fixed64
	ProgramHash string
}

type TxoutputMap struct {
	Key   Uint256
	Txout []TxoutputInfo
}

type AmountMap struct {
	Key   Uint256
	Value Fixed64
}

type ProgramInfo struct {
	Code      string
	Parameter string
}

type Transactions struct {
	TxType         TransactionType
	PayloadVersion byte
	Payload        PayloadInfo
	Nonce          uint64
	Attributes     []TxAttributeInfo
	UTXOInputs     []UTXOTxInputInfo
	BalanceInputs  []BalanceTxInputInfo
	Outputs        []TxoutputInfo
	Programs       []ProgramInfo

	AssetOutputs      []TxoutputMap
	AssetInputAmount  []AmountMap
	AssetOutputAmount []AmountMap

	Hash string
}

type BlockHead struct {
	Version          uint32
	PrevBlockHash    string
	TransactionsRoot string
	Timestamp        uint32
	Height           uint32
	ConsensusData    uint64
	NextMiner        string
	Program          ProgramInfo

	Hash string
}

type BlockInfo struct {
	Hash         string
	BlockData    *BlockHead
	Transactions []*Transactions
}

type TxInfo struct {
	Hash string
	Hex  string
	Tx   *Transactions
}

type TxoutInfo struct {
	High  uint32
	Low   uint32
	Txout tx.TxOutput
}

type NodeInfo struct {
	State    uint   // node status
	Port     uint16 // The nodes's port
	ID       uint64 // The nodes's id
	Time     int64
	Version  uint32 // The network protocol the node used
	Services uint64 // The services the node supplied
	Relay    bool   // The relay capability of the node (merge into capbility flag)
	Height   uint64 // The node latest block height
	TxnCnt   uint64 // The transactions be transmit by this node
	RxTxnCnt uint64 // The transaction received by this node
}

type ConsensusInfo struct {
	// TODO
}

func RegistRpcNode(n Noder) {
	if node == nil {
		node = n
	}
}

func RegistDbftService(d *dbft.DbftService) {
	if dBFT == nil {
		dBFT = d
	}
}

//a function to register functions to be called for specific rpc calls
func HandleFunc(pattern string, handler func([]interface{}) map[string]interface{}) {
	mainMux.Lock()
	defer mainMux.Unlock()
	mainMux.m[pattern] = handler
}

//a function to be called if the request is not a HTTP JSON RPC call
func SetDefaultFunc(def func(http.ResponseWriter, *http.Request)) {
	mainMux.defaultFunction = def
}

//this is the funciton that should be called in order to answer an rpc call
//should be registered like "http.HandleFunc("/", httpjsonrpc.Handle)"
func Handle(w http.ResponseWriter, r *http.Request) {
	mainMux.RLock()
	defer mainMux.RUnlock()
	//JSON RPC commands should be POSTs
	if r.Method != "POST" {
		if mainMux.defaultFunction != nil {
			log.Printf("HTTP JSON RPC Handle - Method!=\"POST\"")
			mainMux.defaultFunction(w, r)
			return
		} else {
			log.Panicf("HTTP JSON RPC Handle - Method!=\"POST\"")
			return
		}
	}

	//check if there is Request Body to read
	if r.Body == nil {
		if mainMux.defaultFunction != nil {
			log.Printf("HTTP JSON RPC Handle - Request body is nil")
			mainMux.defaultFunction(w, r)
			return
		} else {
			log.Panicf("HTTP JSON RPC Handle - Request body is nil")
			return
		}
	}

	//read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("HTTP JSON RPC Handle - ioutil.ReadAll: %v", err)
		return
	}
	request := make(map[string]interface{})
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Fatalf("HTTP JSON RPC Handle - json.Unmarshal: %v", err)
		return
	}

	//get the corresponding function
	function, ok := mainMux.m[request["method"].(string)]
	if ok {
		response := function(request["params"].([]interface{}))
		data, err := json.Marshal(map[string]interface{}{
			"jsonpc": "2.0",
			"result": response["result"],
			"id":     request["id"],
		})
		if err != nil {
			log.Fatalf("HTTP JSON RPC Handle - json.Marshal: %v", err)
			return
		}
		w.Write(data)
	} else {
		//if the function does not exist
		log.Printf("HTTP JSON RPC Handle - No function to call for", request["method"])
		data, err := json.Marshal(map[string]interface{}{
			"result": nil,
			"error": map[string]interface{}{
				"code":    -32601,
				"message": "Method not found",
				"data":    "The called method was not found on the server",
			},
			"id": request["id"],
		})
		if err != nil {
			log.Fatalf("HTTP JSON RPC Handle - json.Marshal: %v", err)
			return
		}
		w.Write(data)
	}
}

func responsePacking(result interface{}) map[string]interface{} {
	resp := map[string]interface{}{
		"result": result,
	}
	return resp
}

// Call sends RPC request to server
func Call(address string, method string, id interface{}, params []interface{}) ([]byte, error) {
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
		return nil, err
	}
	resp, err := http.Post(address, "application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
		return nil, err
	}

	return body, nil
}
