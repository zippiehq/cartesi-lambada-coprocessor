#!/bin/bash

rm /cartesi-lambada-coprocessor/contracts/script/output/lambada_coprocessor_deployment_output_mainnet.json

ANVIL_FORK_URL=`cat /run/secrets/anvil_fork_url_mainnet`
anvil --fork-url $ANVIL_FORK_URL --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545

cast rpc \
    --rpc-url http://0.0.0.0:8545 \
    anvil_impersonateAccount 0x3ad1b118813e71a6b2683FCb2044122fe195AC36


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x1eB68Fa77eB7d263320304B8f6A3C08AE9481469

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0x1eB68Fa77eB7d263320304B8f6A3C08AE9481469 10

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x6F53985155e7204D364e97c2AF41cD2eD02c0dAD

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0x6F53985155e7204D364e97c2AF41cD2eD02c0dAD 10

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xe81a23F9fDa51999E689b0A959ff7C24c9CF8918

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0xe81a23F9fDa51999E689b0A959ff7C24c9CF8918 10


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerMainnet.s.sol:LambadaCoprocessorDeployerMainnet \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
