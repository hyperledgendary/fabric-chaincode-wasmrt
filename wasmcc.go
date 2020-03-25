package main

import (
	"log"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/jt-nti/fabric-chaincode-wasmrt/wasmruntime"
)

type ServerConfig struct {
	CCID     string
	Address  string
	WasmCC   string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.Printf("[host] Wasm Contract runtime..")

	// See chaincode.env.example
	config := ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
		WasmCC:  os.Getenv("CHAINCODE_WASM_FILE"),
	}

	wrt := wasmruntime.NewRuntime(config.WasmCC)

	// err := shim.Start(wrt)
	// check(err)

	server := &shim.ChaincodeServer{
					CCID: config.CCID,
					Address: config.Address,
					CC: wrt,
					TLSProps: shim.TLSProperties{
							Disabled: true,
					},
			}

	err := server.Start()
	check(err)

	return
}
