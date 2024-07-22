#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password 2ivNSRtq5Q0ooKpDewWu \
    --ecdsa-password VNUfEcM5qTiWiCpfhemk \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=2ivNSRtq5Q0ooKpDewWu OPERATOR_ECDSA_KEY_PASSWORD=VNUfEcM5qTiWiCpfhemk operator --config tests/nodes/operators/3/config-devnet.yaml