#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --bls-password XuCFvAtotd6f9BokY5C4 \
    --ecdsa-password tnMtegPwqKSiRZkrXB2Y \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=XuCFvAtotd6f9BokY5C4 OPERATOR_ECDSA_KEY_PASSWORD=tnMtegPwqKSiRZkrXB2Y operator --config tests/nodes/operators/1/config-mainnet.yaml