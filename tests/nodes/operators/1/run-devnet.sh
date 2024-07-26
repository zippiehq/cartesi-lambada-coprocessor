#!/bin/bash

cli setup-operator \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --config tests/nodes/operators/1/config-devnet.yaml \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

operator \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --config tests/nodes/operators/1/config-devnet.yaml