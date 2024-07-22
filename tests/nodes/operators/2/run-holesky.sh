#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password eISqfKXLLeyUmTU2NtpX \
    --ecdsa-password vWdNQd1EcAPK5rOsRs4B \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=eISqfKXLLeyUmTU2NtpX OPERATOR_ECDSA_KEY_PASSWORD=vWdNQd1EcAPK5rOsRs4B operator --config tests/nodes/operators/2/config-holesky.yaml