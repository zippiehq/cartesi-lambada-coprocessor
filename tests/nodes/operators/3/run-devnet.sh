#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password NiaSIGSl45GL6JqlsCbg \
    --ecdsa-password f3KugsTmV2vxkXoQsQ8M \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=NiaSIGSl45GL6JqlsCbg OPERATOR_ECDSA_KEY_PASSWORD=f3KugsTmV2vxkXoQsQ8M operator --config tests/nodes/operators/3/config-devnet.yaml