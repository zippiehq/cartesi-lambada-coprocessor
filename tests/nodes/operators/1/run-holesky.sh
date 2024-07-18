#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --bls-password aObxOqi4nj33uOZG08D4 \
    --ecdsa-password DOK45ch0NnoxcLldxaXF \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=aObxOqi4nj33uOZG08D4 OPERATOR_ECDSA_KEY_PASSWORD=DOK45ch0NnoxcLldxaXF operator --config tests/nodes/operators/1/config-holesky.yaml