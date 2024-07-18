#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password NiaSIGSl45GL6JqlsCbg \
    --ecdsa-password f3KugsTmV2vxkXoQsQ8M \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=NiaSIGSl45GL6JqlsCbg OPERATOR_ECDSA_KEY_PASSWORD=f3KugsTmV2vxkXoQsQ8M operator --config tests/nodes/operators/3/config-mainnet.yaml