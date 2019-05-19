// chainstak testing chalenge

//## geth node address = https://nd-986-703-606.p2pify.com
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	// gethtypes "github.com/ethereum/go-ethereum/core/types"
)

//const crpcadd = "https://nd-986-703-606.p2pify.com"
const crpcadd = "http://127.0.0.1"
const crpcport = "8545"

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

///   main function start
func main() {

	var err error

	gethserver := crpcadd + ":" + crpcport

	ctx := context.TODO()

	rpcc, err := ethrpc.DialContext(ctx, gethserver)

	if err != nil {
		fmt.Printf("straight connection issue\n")
		fmt.Printf("%v", err)
		panic(err)
	} else {
		fmt.Printf("!rpc server connected!\n")
	}

	gethprc, err := ownSyncProgress(ctx, rpcc)
	if err != nil {
		fmt.Print("\n Node is not on Sync mode")
	}

	fmt.Printf("\n%v\n", gethprc)

}

func ownSyncProgress(ctx context.Context, ec *ethrpc.Client) (*gethinfo, error) { // from ethclient/rpc

	var rawdata json.RawMessage

	if err := ec.CallContext(ctx, &rawdata, "eth_syncing"); err != nil {
		return nil, err
	}
	// Handle the possible response types
	var syncing bool
	if err := json.Unmarshal(rawdata, &syncing); err == nil {
		return nil, nil
	}
	var progress *rpcProgress
	if err := json.Unmarshal(rawdata, &progress); err != nil {
		return nil, err
	}

	if err := ec.CallContext(ctx, &rawdata, "net_peerCount"); err != nil {

		return nil, err
	}

	msg := fmt.Sprintf("%s", rawdata)
	num, _ := strconv.Unquote(msg)
	numc, _ := hexutil.DecodeUint64(num)

	fmt.Printf("%s    %s", numc, num)

	return &gethinfo{

		peersnum:        uint64(numc),
		curBlocknum:     uint64(progress.CurrentBlock),
		highestBlocknum: uint64(progress.HighestBlock),
		pulledStatesnum: uint64(progress.PulledStates),
		knownStatesnum:  uint64(progress.KnownStates),
	}, nil

}
