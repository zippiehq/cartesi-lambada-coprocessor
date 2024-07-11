#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password kHQKXbk4k5y2wYlONkCm \
    --ecdsa-password 6xByJMB6elTDeSPT8ppp \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=kHQKXbk4k5y2wYlONkCm OPERATOR_ECDSA_KEY_PASSWORD=6xByJMB6elTDeSPT8ppp operator --config tests/nodes/operators/3/config-mainnet.yaml