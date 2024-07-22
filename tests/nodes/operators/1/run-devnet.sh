#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-devnet.yaml \
    --bls-password 7YNvInegxN5GRFPUh4Ju \
    --ecdsa-password F6rZFySHCcyzou3e4W2Z \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=7YNvInegxN5GRFPUh4Ju OPERATOR_ECDSA_KEY_PASSWORD=F6rZFySHCcyzou3e4W2Z operator --config tests/nodes/operators/1/config-devnet.yaml