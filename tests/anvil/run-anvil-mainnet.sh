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
    --value 20ether 0x0c84ffF0167E5197499DbcEaD25B9c2fB7C89FA8

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0x0c84ffF0167E5197499DbcEaD25B9c2fB7C89FA8 10

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x8674cdDC2D31C9bA13f38889ae866D0051cB2a3C

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0x8674cdDC2D31C9bA13f38889ae866D0051cB2a3C 10

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x51b6f9Bd94055EE166AD27834102929e058f61C0

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0x51b6f9Bd94055EE166AD27834102929e058f61C0 10


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerMainnet.s.sol:LambadaCoprocessorDeployerMainnet \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
