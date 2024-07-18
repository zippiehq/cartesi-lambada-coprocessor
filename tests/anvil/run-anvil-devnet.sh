#!/bin/bash

rm /cartesi-lambada-coprocessor/tests/anvil/devnet-operators-ready.flag

anvil --load-state /root/.anvil/state.json --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x1eB68Fa77eB7d263320304B8f6A3C08AE9481469

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x90c285f405d1e1611b6c9dc51a3528ca7a4eb0713c3eee0cb11b209ddfa8240c \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0x1eB68Fa77eB7d263320304B8f6A3C08AE9481469 10    

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x6F53985155e7204D364e97c2AF41cD2eD02c0dAD

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xdc294635c567ada925b683fde0eec97a2b3a781fabcbdbbece51ad05f4c5a364 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0x6F53985155e7204D364e97c2AF41cD2eD02c0dAD 10    

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xe81a23F9fDa51999E689b0A959ff7C24c9CF8918

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x84473845357bd4387fddef3bcf3347657ead79e17b258f38f4c8a3f2069dfa85 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xe81a23F9fDa51999E689b0A959ff7C24c9CF8918 10    


touch /cartesi-lambada-coprocessor/tests/anvil/devnet-operators-ready.flag

tail -f /dev/null