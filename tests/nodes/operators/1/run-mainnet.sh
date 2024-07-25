#!/bin/bash

cli setup-operator \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

operator \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --config tests/nodes/operators/1/config-mainnet.yaml