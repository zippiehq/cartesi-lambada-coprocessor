#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-devnet.yaml \
    --bls-password aObxOqi4nj33uOZG08D4 \
    --ecdsa-password DOK45ch0NnoxcLldxaXF \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=aObxOqi4nj33uOZG08D4 OPERATOR_ECDSA_KEY_PASSWORD=DOK45ch0NnoxcLldxaXF operator --config tests/nodes/operators/1/config-devnet.yaml