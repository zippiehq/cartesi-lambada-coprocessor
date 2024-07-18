#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password buM47eslEETBTZVhp5DW \
    --ecdsa-password tGDaHRFp0LKNGcobHUYE \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=buM47eslEETBTZVhp5DW OPERATOR_ECDSA_KEY_PASSWORD=tGDaHRFp0LKNGcobHUYE operator --config tests/nodes/operators/2/config-devnet.yaml