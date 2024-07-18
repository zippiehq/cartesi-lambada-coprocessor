#!/bin/bash

rm /cartesi-lambada-coprocessor/contracts/script/output/lambada_coprocessor_deployment_output_holesky.json

ANVIL_FORK_URL=`cat /run/secrets/anvil_fork_url_holesky`
anvil --fork-url $ANVIL_FORK_URL --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x0c84ffF0167E5197499DbcEaD25B9c2fB7C89FA8

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xfcef29595390333a70fd43862ab2ba66021c5d96131b8372d9e56ae65a66012d \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x8674cdDC2D31C9bA13f38889ae866D0051cB2a3C

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x52956691667302e45d9c84dff635e811a34e50d6603cd3b91c46f3fc3b043814 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x51b6f9Bd94055EE166AD27834102929e058f61C0

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x1da692825f17f894caa8170de6f260907216b11ea2f51690a5dcbb972f904c6c \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
