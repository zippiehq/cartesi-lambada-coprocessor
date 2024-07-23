#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=hG6zoZiGJOf6okxSO1G5 OPERATOR_ECDSA_KEY_PASSWORD=KRnd3uPtKuCl1NkDWzdE operator --config tests/nodes/operators/1/config-holesky.yaml