#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password 2Vef5SPHZIUoVDyVJztv \
    --ecdsa-password ddKID9eipR6LSoWSjBsM \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=2Vef5SPHZIUoVDyVJztv OPERATOR_ECDSA_KEY_PASSWORD=ddKID9eipR6LSoWSjBsM operator --config tests/nodes/operators/2/config-mainnet.yaml