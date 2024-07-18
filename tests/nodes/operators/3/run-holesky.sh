#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password kHQKXbk4k5y2wYlONkCm \
    --ecdsa-password 6xByJMB6elTDeSPT8ppp \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=kHQKXbk4k5y2wYlONkCm OPERATOR_ECDSA_KEY_PASSWORD=6xByJMB6elTDeSPT8ppp operator --config tests/nodes/operators/3/config-holesky.yaml