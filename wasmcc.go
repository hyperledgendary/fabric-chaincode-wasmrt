package main

import (
	"log"
	"github.com/jt-nti/fabric-chaincode-wasmrt/wasmruntime"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.Printf("[host] Wasm Contract runtime..")

	// TODO pass the Wasm chaincode in with an env var?
	wrt := wasmruntime.NewRuntime("/chaincode/input/src/wasmruntime/fabric_contract.wasm")

	err := shim.Start(wrt)
	check(err)

	return
}
