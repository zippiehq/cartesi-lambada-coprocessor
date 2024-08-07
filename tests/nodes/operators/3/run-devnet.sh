#!/bin/bash

cli setup-operator \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

operator \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --config tests/nodes/operators/3/config-devnet.yaml