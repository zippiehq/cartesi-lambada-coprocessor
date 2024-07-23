#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=I0k5Um80exGzslUaSE0v OPERATOR_ECDSA_KEY_PASSWORD=ui9Sa12QBtcteTNKokDV operator --config tests/nodes/operators/2/config-devnet.yaml