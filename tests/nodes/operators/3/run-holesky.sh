#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password Wi2ivNSRtq5Q0ooKpDew \
    --ecdsa-password aDVNUfEcM5qTiWiCpfhe \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=Wi2ivNSRtq5Q0ooKpDew OPERATOR_ECDSA_KEY_PASSWORD=aDVNUfEcM5qTiWiCpfhe operator --config tests/nodes/operators/3/config-holesky.yaml