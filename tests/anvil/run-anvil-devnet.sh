#!/bin/bash

rm /cartesi-lambada-coprocessor/tests/anvil/devnet-operators-ready.flag

anvil --load-state /root/.anvil/state.json --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x96739DD77b76f30f9c524fE233d46c57A6C72423

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x173a189d6ab78e2417b7c9c9934d2c661a8cff9dc92e7b8e4ab84cc49f4e7911 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0x96739DD77b76f30f9c524fE233d46c57A6C72423 10    

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xeA1BEcba988444fEa429cbE9b940D02926d74E42

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x6f379616b5a64d0e3054c272553155ef7f47f7c56240d60777711ce9d2664690 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xeA1BEcba988444fEa429cbE9b940D02926d74E42 10    

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xe70e092C9712Ae79021a94670b37bc42bed3B2ae

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x30c0c6bc35f8405118132be5560f777782baec5b09b8fa859f9f441e1804f4ef \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xe70e092C9712Ae79021a94670b37bc42bed3B2ae 10    


touch /cartesi-lambada-coprocessor/tests/anvil/devnet-operators-ready.flag

tail -f /dev/null