#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --bls-password aObxOqi4nj33uOZG08D4 \
    --ecdsa-password DOK45ch0NnoxcLldxaXF \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=aObxOqi4nj33uOZG08D4 OPERATOR_ECDSA_KEY_PASSWORD=DOK45ch0NnoxcLldxaXF operator --config tests/nodes/operators/1/config-mainnet.yaml