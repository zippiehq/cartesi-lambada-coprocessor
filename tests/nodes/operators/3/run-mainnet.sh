#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password TiIE2hx6gxuKoS1NOru2 \
    --ecdsa-password BriyubIIlIUrOipru4Pn \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=TiIE2hx6gxuKoS1NOru2 OPERATOR_ECDSA_KEY_PASSWORD=BriyubIIlIUrOipru4Pn operator --config tests/nodes/operators/3/config-mainnet.yaml