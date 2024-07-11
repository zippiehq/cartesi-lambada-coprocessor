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
    --value 20ether 0x96739DD77b76f30f9c524fE233d46c57A6C72423

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0x96739DD77b76f30f9c524fE233d46c57A6C72423 10

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xeA1BEcba988444fEa429cbE9b940D02926d74E42

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0xeA1BEcba988444fEa429cbE9b940D02926d74E42 10

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xe70e092C9712Ae79021a94670b37bc42bed3B2ae

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --from 0x3ad1b118813e71a6b2683FCb2044122fe195AC36 \
    --unlocked \
    0xae78736Cd615f374D3085123A210448E74Fc6393 "transfer(address,uint256)(bool)" 0xe70e092C9712Ae79021a94670b37bc42bed3B2ae 10


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerMainnet.s.sol:LambadaCoprocessorDeployerMainnet \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
