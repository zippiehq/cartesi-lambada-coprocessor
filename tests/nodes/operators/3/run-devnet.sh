#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=TiIE2hx6gxuKoS1NOru2 OPERATOR_ECDSA_KEY_PASSWORD=BriyubIIlIUrOipru4Pn operator --config tests/nodes/operators/3/config-devnet.yaml