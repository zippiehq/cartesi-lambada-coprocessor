#!/bin/bash

cli setup-operator \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

operator \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --config tests/nodes/operators/2/config-mainnet.yaml