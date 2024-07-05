#!/bin/bash

sleep 0

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xc48d425C747a468Ca3810BDa267DC4749E6036E3


cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0x7e86ed15e215f54a6fc578297554c477280b0d4faa537d15b4a8c2cdb48969e0 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xc48d425C747a468Ca3810BDa267DC4749E6036E3 10    


cli setup-operator \
    --config tests/nodes/operators/1/config.yaml \
    --bls-password LuqUSZmEeOSdUJipJHF3 \
    --ecdsa-password aK07ZKg6Ufc3JtAgkMsl \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=LuqUSZmEeOSdUJipJHF3 OPERATOR_ECDSA_KEY_PASSWORD=aK07ZKg6Ufc3JtAgkMsl operator --config tests/nodes/operators/1/config.yaml