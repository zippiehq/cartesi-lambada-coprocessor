#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password NXJPGvJudjdVFI8YfOFI \
    --ecdsa-password YBggWv4Xc2BqDK0QYqql \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=NXJPGvJudjdVFI8YfOFI OPERATOR_ECDSA_KEY_PASSWORD=YBggWv4Xc2BqDK0QYqql operator --config tests/nodes/operators/3/config-mainnet.yaml