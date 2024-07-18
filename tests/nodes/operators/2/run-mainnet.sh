#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password kySHxby2AAGcMLPLpKRw \
    --ecdsa-password U2kjFhz6DEeWYSXBBltt \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=kySHxby2AAGcMLPLpKRw OPERATOR_ECDSA_KEY_PASSWORD=U2kjFhz6DEeWYSXBBltt operator --config tests/nodes/operators/2/config-mainnet.yaml