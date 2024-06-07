anvil --fork-url https://holesky.infura.io/v3/7d2631bb88d545bc9959384210bae054

forge script script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky --rpc-url http://127.0.0.1:8545 --broadcast -v --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

Aggregator (10th anvil dev account)

go run aggregator/cmd/main.go --config tests/nodes/aggregator/aggregator.yaml --credible-squaring-deployment contracts/script/output/lambada_coprocessor_deployment_output.holesky.json  --ecdsa-private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

Operator 1:

go run cli/main.go --config=tests/nodes/operators/configs/operator1.yaml setup-operator --bls-password fZEOWNaTwe9pBkxAQASW --ecdsa-password TurBzPaKwxXdCplqgiS6 --deployment-parameters contracts/script/input/parameters.holesky.json -deposit-amount-weth 5 --devnet

go run operator/cmd/main.go --config=tests/nodes/operators/configs/operator1.yaml

Operator 2:

Operator 3:
