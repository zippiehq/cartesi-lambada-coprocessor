# Cartesi Lambada Coprocessor

This is proof of concept of verifiable computing system using  [Cartesi virtual machine](https://github.com/zippiehq/cartesi-lambada) and [Eigenlayer restaking platform](https://github.com/Layr-Labs/eigenlayer-contracts).

## Contract development

After changing contract source code in `contracts/src`, run `make deploy-avs`. This will update `tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json` by deploying new version of avs contracts and registering new set of operators from scratch (`tests/nodes/operators` will be updated accordingly).

## Testing

1. In first termnial run `docker compose up`
2. In second termnial run `make testing-integration`
3. Wait for tests to complete without any errors
4. Check docker logs for unexpected errors

## Configuring number of operators

Do not change number of operators in `main` branch.

To run avs with different number of operators, in separate branch:

1. Update number of generated operators [here](https://github.com/zippiehq/cartesi-lambada-coprocessor/blob/main/tests/anvil/deploy-avs-save-anvil-state.sh#L26)
2. Run `make deploy-avs`

## Troubleshooting

`tests/nodes/operators` contains mounted docker volumes, which are reused across multiple runs of docker compose. Sometimes it causes `permission denied` errors while running `docker compose build`, `make deploy-avs` and `make tests-integration`. To fix this issue - run `chomd -R 777 tests/nodes/operators`.
