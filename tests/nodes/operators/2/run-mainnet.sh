#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password ISqfKXLLeyUmTU2NtpXh \
    --ecdsa-password dNQd1EcAPK5rOsRs4BqK \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=ISqfKXLLeyUmTU2NtpXh OPERATOR_ECDSA_KEY_PASSWORD=dNQd1EcAPK5rOsRs4BqK operator --config tests/nodes/operators/2/config-mainnet.yaml