#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password kySHxby2AAGcMLPLpKRw \
    --ecdsa-password U2kjFhz6DEeWYSXBBltt \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=kySHxby2AAGcMLPLpKRw OPERATOR_ECDSA_KEY_PASSWORD=U2kjFhz6DEeWYSXBBltt operator --config tests/nodes/operators/2/config-devnet.yaml