#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-devnet.yaml \
    --bls-password XuCFvAtotd6f9BokY5C4 \
    --ecdsa-password tnMtegPwqKSiRZkrXB2Y \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=XuCFvAtotd6f9BokY5C4 OPERATOR_ECDSA_KEY_PASSWORD=tnMtegPwqKSiRZkrXB2Y operator --config tests/nodes/operators/1/config-devnet.yaml