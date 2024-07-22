#!/bin/bash

rm /cartesi-lambada-coprocessor/contracts/script/output/lambada_coprocessor_deployment_output_holesky.json

ANVIL_FORK_URL=`cat /run/secrets/anvil_fork_url_holesky`
anvil --fork-url $ANVIL_FORK_URL --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x56B9D78f83bd4F189b83939cE0234943de52DAe2

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x85ece1511455780875d64ee2d3d0d0de6bf8f9b44ce85ff044c6b1f83b8e883b \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x66d8eB27148CB259c0810b7AbE004F963d23e052

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x9e4f5aa6aec3fc78c6aae081ac8120c720efcd6cea84b6925e607be063716f96 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x6040D45Fa68d87d0bD28b4D11834c34491860218

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xca0823ae02d67d866ac2c4fe4a725053da119b9d4f515140a2d7239c40b45ac3 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
