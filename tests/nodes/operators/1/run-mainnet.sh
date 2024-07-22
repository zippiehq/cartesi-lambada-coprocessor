#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --bls-password 7YNvInegxN5GRFPUh4Ju \
    --ecdsa-password F6rZFySHCcyzou3e4W2Z \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=7YNvInegxN5GRFPUh4Ju OPERATOR_ECDSA_KEY_PASSWORD=F6rZFySHCcyzou3e4W2Z operator --config tests/nodes/operators/1/config-mainnet.yaml