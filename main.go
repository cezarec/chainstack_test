// chainstak testing chalenge

//## geth node address = https://nd-986-703-606.p2pify.com
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	// gethtypes "github.com/ethereum/go-ethereum/core/types"
)

//const crpcadd = "https://nd-986-703-606.p2pify.com"
const crpcadd = "http://127.0.0.1"
const crpcport = "8545"
const cpollduration = 100
const clistenport = "9080" // port of exporter daemon

type gethinfo struct {
	//rpcserver       string
	peersnum        uint64 // number of peers
	curBlocknum     uint64 //number if currentBlock
	highestBlocknum uint64
	knownStatesnum  uint64
	pulledStatesnum uint64
}

type rpcProgress struct {
	StartingBlock hexutil.Uint64
	CurrentBlock  hexutil.Uint64
	HighestBlock  hexutil.Uint64
	PulledStates  hexutil.Uint64
	KnownStates   hexutil.Uint64
}

var (
	rpcc    *ethrpc.Client
	gethrpc *gethinfo
)

///   main function start
func main() {

	var err error

	gethserver := crpcadd + ":" + crpcport

	ctx := context.TODO()

	rpcc, err = ethrpc.DialContext(ctx, gethserver)

	if err != nil {
		fmt.Printf("rpc connection issue\n")
		panic(err)
	} else {
		fmt.Printf("!rpc server connected!\n")
	}

	go func(ctx context.Context) {
		for {
			gethrpc, err = ownSyncProgress(ctx, rpcc)
			if err != nil {
				fmt.Print("\n Node is not on Sync mode")

			}
			fmt.Printf("info %v \n", gethrpc)
			time.Sleep(time.Duration(cpollduration) * time.Millisecond)
		}
	}(ctx)

	http.HandleFunc("/metrics", Tometricshttp)
	err = http.ListenAndServe(":"+clistenport, nil)
	if err != nil {
		panic(err)
	}

}

func ownSyncProgress(ctx context.Context, ec *ethrpc.Client) (*gethinfo, error) { // from ethclient/rpc

	var rawdata json.RawMessage

	if err := ec.CallContext(ctx, &rawdata, "eth_syncing"); err != nil {
		return nil, err
	}
	// Handle the possible response types

	var progress *rpcProgress
	if err := json.Unmarshal(rawdata, &progress); err != nil {
		return nil, err
	}

	if err := ec.CallContext(ctx, &rawdata, "net_peerCount"); err != nil {

		return nil, err
	}

	num, _ := strconv.Unquote(fmt.Sprintf("%s", rawdata))
	numc, _ := hexutil.DecodeUint64(num)
	return &gethinfo{

		peersnum:        uint64(numc),
		curBlocknum:     uint64(progress.CurrentBlock),
		highestBlocknum: uint64(progress.HighestBlock),
		pulledStatesnum: uint64(progress.PulledStates),
		knownStatesnum:  uint64(progress.KnownStates),
	}, nil

}

func Tometricshttp(hwriter http.ResponseWriter, hreq *http.Request) {

	strbuf := []string{}

	strbuf = append(strbuf, fmt.Sprintf("peers_number %v", gethrpc.peersnum))
	strbuf = append(strbuf, fmt.Sprintf("current_block %v", gethrpc.curBlocknum))
	strbuf = append(strbuf, fmt.Sprintf("highest_block %v", gethrpc.highestBlocknum))
	strbuf = append(strbuf, fmt.Sprintf("known_states %v", gethrpc.knownStatesnum))
	strbuf = append(strbuf, fmt.Sprintf("pulled_states %v", gethrpc.pulledStatesnum))

	hwriter.Write([]byte(strings.Join(strbuf, "\n")))
}
