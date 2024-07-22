#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password ISqfKXLLeyUmTU2NtpXh \
    --ecdsa-password dNQd1EcAPK5rOsRs4BqK \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=ISqfKXLLeyUmTU2NtpXh OPERATOR_ECDSA_KEY_PASSWORD=dNQd1EcAPK5rOsRs4BqK operator --config tests/nodes/operators/2/config-holesky.yaml