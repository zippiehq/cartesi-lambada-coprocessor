#!/bin/bash

rm /cartesi-lambada-coprocessor/tests/anvil/devnet-operators-ready.flag

anvil --load-state /root/.anvil/state.json --host 0.0.0.0 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x6e54781FE5f8bfaf73B4a5e3deF6eeff2297606b

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xb0031dffe44d2a5cab5c4c935e0adb58fbc30477fbf5b214f3091e8bd468b893 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0x6e54781FE5f8bfaf73B4a5e3deF6eeff2297606b 10    

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xE2C7b758068ffc9611eA1a12dE09b13Bba373448

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x895210ed12367fec4405cb055a5017ac9220dda809571e92b4d85add69a58867 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xE2C7b758068ffc9611eA1a12dE09b13Bba373448 10    

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x7fC2E7a02E667CD1C7F4eD12F6D8919B1F76D3df

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x8308edbed4025dda17775621be50f8d81bd44afa129ddc237c964d8db50fa305 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0x7fC2E7a02E667CD1C7F4eD12F6D8919B1F76D3df 10    


touch /cartesi-lambada-coprocessor/tests/anvil/devnet-operators-ready.flag

tail -f /dev/null