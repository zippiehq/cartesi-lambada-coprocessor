#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password buM47eslEETBTZVhp5DW \
    --ecdsa-password tGDaHRFp0LKNGcobHUYE \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=buM47eslEETBTZVhp5DW OPERATOR_ECDSA_KEY_PASSWORD=tGDaHRFp0LKNGcobHUYE operator --config tests/nodes/operators/2/config-mainnet.yaml