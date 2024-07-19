#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password NXJPGvJudjdVFI8YfOFI \
    --ecdsa-password YBggWv4Xc2BqDK0QYqql \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=NXJPGvJudjdVFI8YfOFI OPERATOR_ECDSA_KEY_PASSWORD=YBggWv4Xc2BqDK0QYqql operator --config tests/nodes/operators/3/config-holesky.yaml