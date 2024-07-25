#!/bin/bash

cli setup-operator \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

operator \
    --bls-password hG6zoZiGJOf6okxSO1G5 \
    --ecdsa-password KRnd3uPtKuCl1NkDWzdE \
    --config tests/nodes/operators/1/config-holesky.yaml