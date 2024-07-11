#/bin/sh
rm /cartesi-lambada-coprocessor/contracts/script/output/lambada_coprocessor_deployment_output_holesky.json

anvil --fork-url https://wispy-tiniest-owl.ethereum-holesky.quiknode.pro/be5b39039e9f0d2b7819fb3f18bffc82e58ad269/ --host 0.0.0.0 &

timeout 22 bash -c 'until printf "" 2>>/dev/null >>/dev/tcp/$0/$1; do sleep 1; done' 0.0.0.0:8545

cd /cartesi-lambada-coprocessor/contracts

forge script \
    script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky \
    --rpc-url http://0.0.0.0:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

tail -f /dev/null