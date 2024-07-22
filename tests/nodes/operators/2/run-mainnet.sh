#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password eISqfKXLLeyUmTU2NtpX \
    --ecdsa-password vWdNQd1EcAPK5rOsRs4B \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=eISqfKXLLeyUmTU2NtpX OPERATOR_ECDSA_KEY_PASSWORD=vWdNQd1EcAPK5rOsRs4B operator --config tests/nodes/operators/2/config-mainnet.yaml