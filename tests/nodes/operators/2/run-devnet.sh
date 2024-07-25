#!/bin/bash

cli setup-operator \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

operator \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --config tests/nodes/operators/2/config-devnet.yaml