#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password FMPLCKxcmt2Od3dtjn48 \
    --ecdsa-password cYpXQmvEdA3PAs91qfqe \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=FMPLCKxcmt2Od3dtjn48 OPERATOR_ECDSA_KEY_PASSWORD=cYpXQmvEdA3PAs91qfqe operator --config tests/nodes/operators/3/config-devnet.yaml