package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	eth := new(ethclient.Client)

	fmt.Printf("%v", eth)

}
