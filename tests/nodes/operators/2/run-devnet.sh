#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password 2Vef5SPHZIUoVDyVJztv \
    --ecdsa-password ddKID9eipR6LSoWSjBsM \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=2Vef5SPHZIUoVDyVJztv OPERATOR_ECDSA_KEY_PASSWORD=ddKID9eipR6LSoWSjBsM operator --config tests/nodes/operators/2/config-devnet.yaml