#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password 2ivNSRtq5Q0ooKpDewWu \
    --ecdsa-password VNUfEcM5qTiWiCpfhemk \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=2ivNSRtq5Q0ooKpDewWu OPERATOR_ECDSA_KEY_PASSWORD=VNUfEcM5qTiWiCpfhemk operator --config tests/nodes/operators/3/config-holesky.yaml