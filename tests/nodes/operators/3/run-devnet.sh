#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password Wi2ivNSRtq5Q0ooKpDew \
    --ecdsa-password aDVNUfEcM5qTiWiCpfhe \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=Wi2ivNSRtq5Q0ooKpDew OPERATOR_ECDSA_KEY_PASSWORD=aDVNUfEcM5qTiWiCpfhe operator --config tests/nodes/operators/3/config-devnet.yaml