#!/bin/bash

rm /cartesi-lambada-coprocessor/contracts/script/output/coprocessor_deployment_output_holesky.json

ANVIL_FORK_URL=`cat /run/secrets/anvil_fork_url_holesky`
anvil --fork-url $ANVIL_FORK_URL --host 0.0.0.0 --block-time 12 &
timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545


cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x02C9ca5313A6E826DC05Bbe098150b3215D5F821

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xc276a0e2815b89e9a3d8b64cb5d745d5b4f6b84531306c97aad82156000a7dd7 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x71f897938C155D4569b9f8fbff8fBFC7A89069Fb

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0x850affd0f354c8b3b3176ae914dcd90cdb0f2051d1c5e31c9bbf97a732b68a07 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xEc1dc4D2a9459758DCe2bb13096F303a8FAF4c92

cast send \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xa38181ab9321e4bfdfe8dae9f99b05529483bdf0254bd5bcbcc51232f26b8c36 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/CoprocessorDeployerHoleksy.s.sol:CoprocessorDeployerHolesky \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null
