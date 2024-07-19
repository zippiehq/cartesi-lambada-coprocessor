#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password NXJPGvJudjdVFI8YfOFI \
    --ecdsa-password YBggWv4Xc2BqDK0QYqql \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=NXJPGvJudjdVFI8YfOFI OPERATOR_ECDSA_KEY_PASSWORD=YBggWv4Xc2BqDK0QYqql operator --config tests/nodes/operators/3/config-devnet.yaml