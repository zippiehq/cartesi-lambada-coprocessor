#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password ISqfKXLLeyUmTU2NtpXh \
    --ecdsa-password dNQd1EcAPK5rOsRs4BqK \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=ISqfKXLLeyUmTU2NtpXh OPERATOR_ECDSA_KEY_PASSWORD=dNQd1EcAPK5rOsRs4BqK operator --config tests/nodes/operators/2/config-devnet.yaml