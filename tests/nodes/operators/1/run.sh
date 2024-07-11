#!/bin/bash

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xB0Acf8523E1908CF6aa3B23dDa2E2f83f3928727


cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xd3b5d5424bdc7e9dcf85ee1c6373fa76b4f43cc3b8427c23d45dbaa2934d3ef2 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cli setup-operator \
    --config tests/nodes/operators/1/config.yaml \
    --bls-password Btr571lnI3u2GBe3u9lc \
    --ecdsa-password PbkbXsUI7JUyoK8Hpiua \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=Btr571lnI3u2GBe3u9lc OPERATOR_ECDSA_KEY_PASSWORD=PbkbXsUI7JUyoK8Hpiua operator --config tests/nodes/operators/1/config.yaml