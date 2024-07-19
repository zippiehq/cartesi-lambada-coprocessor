#!/bin/bash

rm /cartesi-lambada-coprocessor/contracts/script/output/lambada_coprocessor_deployment_output_holesky.json

ANVIL_FORK_URL=`cat /run/secrets/anvil_fork_url_holesky`
anvil --fork-url $ANVIL_FORK_URL --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x6e54781FE5f8bfaf73B4a5e3deF6eeff2297606b

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xb0031dffe44d2a5cab5c4c935e0adb58fbc30477fbf5b214f3091e8bd468b893 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xE2C7b758068ffc9611eA1a12dE09b13Bba373448

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x895210ed12367fec4405cb055a5017ac9220dda809571e92b4d85add69a58867 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x7fC2E7a02E667CD1C7F4eD12F6D8919B1F76D3df

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x8308edbed4025dda17775621be50f8d81bd44afa129ddc237c964d8db50fa305 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
