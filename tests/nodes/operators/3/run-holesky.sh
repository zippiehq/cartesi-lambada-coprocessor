#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-holesky.yaml \
    --bls-password NiaSIGSl45GL6JqlsCbg \
    --ecdsa-password f3KugsTmV2vxkXoQsQ8M \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=NiaSIGSl45GL6JqlsCbg OPERATOR_ECDSA_KEY_PASSWORD=f3KugsTmV2vxkXoQsQ8M operator --config tests/nodes/operators/3/config-holesky.yaml