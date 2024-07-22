#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --bls-password 7YNvInegxN5GRFPUh4Ju \
    --ecdsa-password F6rZFySHCcyzou3e4W2Z \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=7YNvInegxN5GRFPUh4Ju OPERATOR_ECDSA_KEY_PASSWORD=F6rZFySHCcyzou3e4W2Z operator --config tests/nodes/operators/1/config-holesky.yaml