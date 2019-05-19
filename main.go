// chainstak testing chalenge

//## geth node address = https://nd-986-703-606.p2pify.com
package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
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

///   main function start
func main() {

	var err error
	var eth *ethclient.Client
	geth := gethinfo{}
	gethserver := crpcadd + ":" + crpcport

	ctx := context.TODO()
	//defer eth.Close()

	eth, _ = ethclient.Dial(gethserver)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("connected to rpc point\n")
	}

	//curBlock, err := eth.BlockByNumber(ctx, nil)

	//eth.curBlocknum = curBlock.Header.number

	curSP, err := eth.SyncProgress(ctx)
	if err != nil {
		fmt.Print("\n syncProgress got\n")
	} else {
		fmt.Printf("\nsyncProgrerssIssue")
	}

	if curSP == nil {
		fmt.Printf(" no sync currently running")
		return
	}

	geth.highestBlocknum = curSP.HighestBlock
	geth.curBlocknum = curSP.CurrentBlock
	geth.knownStatesnum = curSP.KnownStates
	geth.pulledStatesnum = curSP.PulledStates

	fmt.Printf("\n%v\n", curSP)
	fmt.Printf("%v\n", curSP.CurrentBlock)

	fmt.Printf("\n found %v", geth)

}
