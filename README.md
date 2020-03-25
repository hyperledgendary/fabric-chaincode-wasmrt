# fabric-chaincode-wasmrt

To run the service in a container, build a FabCar docker image:

```
docker build -t wasm-chaincode .
```

Edit the chaincode.env file before starting a Wasm chaincode container using the following command:

```
docker run -it --rm --name wasmcc.example.com --hostname wasmcc.example.com --env-file chaincode.env --network=net_test wasm-chaincode
```
