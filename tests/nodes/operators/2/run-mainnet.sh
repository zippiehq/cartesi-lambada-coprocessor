#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=I0k5Um80exGzslUaSE0v OPERATOR_ECDSA_KEY_PASSWORD=ui9Sa12QBtcteTNKokDV operator --config tests/nodes/operators/2/config-mainnet.yaml