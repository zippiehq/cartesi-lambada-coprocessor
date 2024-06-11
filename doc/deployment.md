To generate ECDSA and BLS keys use [egnkey tool](https://github.com/Layr-Labs/eigensdk-go/tree/master/cmd/egnkey)

All commands are need to be run from repository root (while all configuration and data for aggregator and operators are stored in the `nodes/` directory).

Update [set-deployment-constants.sh](set-deployment-constants.sh) with your own values. Then run `source doc/set-deployment-constants.sh` before executing commands belw.

1. Generate ecdsa key/address for aggregator
   `egnkey generate --key-type ecdsa --num-keys 1`
   Create [aggregator configuration file](../tests/nodes/aggregator/aggregator.yaml)

2. Generate ecdsa key/address for "admin roles" (owner, churner, ejector, confirmer)
   `egnkey generate --key-type ecdsa --num-keys 1`

3. Create [deployment conifguration file](../contracts/script/input/parameters.holesky.json) 

4. Run [deployment script](../contracts/script/LambadaCoprocessorDeployerHoleksy.s.sol)
   `cd contracts/`
   `forge script script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky --rpc-url http://127.0.0.1:8545 --broadcast -v --private-key $FUNDER_PK`
   `cd ../`

5. Fund aggregator with ETH
   `cast send --private-key $FUNDER_PK --value 20ether $AGGREGATOR_ADDRESS`
   Run aggregator
   `go run aggregator/cmd/main.go --config $AGGREGATOR_CONFIG --credible-squaring-deployment contracts/script/output/lambada_coprocessor_deployment_output.holesky.json --ecdsa-private-key $AGGREGATOR_PK`

6. Generate bls and ecdsa keys for operator:
   `egnkey generate --key-type ecdsa --num-keys 1`
   `egnkey generate --key-type bls --num-keys 1`
   
   Create [operator configruation file](../tests/nodes/operators/configs/operator1.yaml)

7. Fund operator with ETH
   `cast send --private-key $FUNDER_PK --value 20ether $OPERATOR_ADDRESS`,
   Fund operator with WETH
   `cast send --private-key $OPERATOR_1_PK --value 10ether $WETH_HOLESKY_ADDRESS 'deposit'`
   Register operator
   `go run cli/main.go --config $OPERATOR_1_CONFIG setup-operator --bls-password $OPERATOR_1_BLS_PWD --ecdsa-password $OPERATOR_1_ECDSA_PWD --deployment-parameters contracts/script/input/parameters.holesky.json -deposit-amount-weth 5`

8. Run operator machine
   `docker run -v ./nodes/operator-1/data:/data -v ./machine/data-preload:/data/preload -p 5001:5001 -p 30033:3033 -e COMPUTE_ONLY='1' ghcr.io/zippiehq/cartesi-lambada:amd64-latest`
   Run operator
   `OPERATOR_BLS_KEY_PASSWORD=$OPERATOR_1_BLS_PWD OPERATOR_ECDSA_KEY_PASSWORD=#OPERATOR_1_ECDSA_PWD go run operator/cmd/main.go --config=tests/nodes/operator-1/operator-1.yaml`
