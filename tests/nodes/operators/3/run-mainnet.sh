#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-mainnet.yaml \
    --bls-password FMPLCKxcmt2Od3dtjn48 \
    --ecdsa-password cYpXQmvEdA3PAs91qfqe \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=FMPLCKxcmt2Od3dtjn48 OPERATOR_ECDSA_KEY_PASSWORD=cYpXQmvEdA3PAs91qfqe operator --config tests/nodes/operators/3/config-mainnet.yaml