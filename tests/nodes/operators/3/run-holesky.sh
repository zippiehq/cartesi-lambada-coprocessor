#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password FMPLCKxcmt2Od3dtjn48 \
    --ecdsa-password cYpXQmvEdA3PAs91qfqe \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=FMPLCKxcmt2Od3dtjn48 OPERATOR_ECDSA_KEY_PASSWORD=cYpXQmvEdA3PAs91qfqe operator --config tests/nodes/operators/3/config-holesky.yaml