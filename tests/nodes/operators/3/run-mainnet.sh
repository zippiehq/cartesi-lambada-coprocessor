#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password 2ivNSRtq5Q0ooKpDewWu \
    --ecdsa-password VNUfEcM5qTiWiCpfhemk \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=2ivNSRtq5Q0ooKpDewWu OPERATOR_ECDSA_KEY_PASSWORD=VNUfEcM5qTiWiCpfhemk operator --config tests/nodes/operators/3/config-mainnet.yaml