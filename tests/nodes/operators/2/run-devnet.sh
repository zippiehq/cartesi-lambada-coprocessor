#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password eISqfKXLLeyUmTU2NtpX \
    --ecdsa-password vWdNQd1EcAPK5rOsRs4B \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=eISqfKXLLeyUmTU2NtpX OPERATOR_ECDSA_KEY_PASSWORD=vWdNQd1EcAPK5rOsRs4B operator --config tests/nodes/operators/2/config-devnet.yaml