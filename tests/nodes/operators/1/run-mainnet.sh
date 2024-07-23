#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=hG6zoZiGJOf6okxSO1G5 OPERATOR_ECDSA_KEY_PASSWORD=KRnd3uPtKuCl1NkDWzdE operator --config tests/nodes/operators/1/config-mainnet.yaml