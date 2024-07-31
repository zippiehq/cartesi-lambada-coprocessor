#!/bin/bash

cli setup-operator \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

operator \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --config tests/nodes/operators/3/config-holesky.yaml