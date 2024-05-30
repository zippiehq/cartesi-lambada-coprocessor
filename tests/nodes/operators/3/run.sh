#!/bin/bash

sleep 2

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xCb07803b31671Be2978D2695D1DF900fb40601D5


cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0x23dab62bf34d2a306288a01ebfb89a640eeb45cd834627757d4ac3aeb9734f52 \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xCb07803b31671Be2978D2695D1DF900fb40601D5 10    


cli setup-operator \
    --config tests/nodes/operators/3/config.yaml \
    --bls-password Tj0YQji8LeQLqnZBDOBe \
    --ecdsa-password niqBH1FHT4bqtDarUwQG \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=Tj0YQji8LeQLqnZBDOBe OPERATOR_ECDSA_KEY_PASSWORD=niqBH1FHT4bqtDarUwQG operator --config tests/nodes/operators/3/config.yaml