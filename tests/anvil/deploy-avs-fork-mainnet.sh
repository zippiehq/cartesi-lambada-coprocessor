#/bin/sh
rm ./contracts/script/output/lambada_coprocessor_deployment_output_mainnet.json

anvil --fork-url https://holesky.infura.io/v3/7d2631bb88d545bc9959384210bae054 --host 0.0.0.0 &

timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545

cd ./contracts

forge script \
    script/LambadaCoprocessorDeployerMainnet.s.sol:LambadaCoprocessorDeployerMainnet \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v