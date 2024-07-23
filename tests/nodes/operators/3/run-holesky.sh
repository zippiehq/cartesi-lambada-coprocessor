#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=TiIE2hx6gxuKoS1NOru2 OPERATOR_ECDSA_KEY_PASSWORD=BriyubIIlIUrOipru4Pn operator --config tests/nodes/operators/3/config-holesky.yaml