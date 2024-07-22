#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password Wi2ivNSRtq5Q0ooKpDew \
    --ecdsa-password aDVNUfEcM5qTiWiCpfhe \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=Wi2ivNSRtq5Q0ooKpDew OPERATOR_ECDSA_KEY_PASSWORD=aDVNUfEcM5qTiWiCpfhe operator --config tests/nodes/operators/3/config-mainnet.yaml