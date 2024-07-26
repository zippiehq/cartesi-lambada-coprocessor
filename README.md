# Cartesi Lambada Coprocessor

This is proof of concept of verifiable computing system using  [Cartesi virtual machine](https://github.com/zippiehq/cartesi-lambada) and [Eigenlayer restaking platform](https://github.com/Layr-Labs/eigenlayer-contracts).

## Running tests

To run tests on devnet (local deployment of Eigenlayer and Lambada Coprocessor AVS) in first terminal run `docker compose up`. Then in second terminal run `make tests-integration`.

Use `docker compose -f docker-compose-holesky.yaml` and `make tests-integration-holesky` to run tests on Holesky fork.

Use `docker compose -f docker-compose-mainnet.yaml` and `make tests-integration-mainnet` to run tests on Mainnet fork.

Wait for tests to complete without any errors and check docker compose logs for unexpected errors.

## Configuring number of operators

Run `go run ./cli generate-docker-compose --operators $NUMBER_OF_OPERATORS` to set number of operators, used by testing environment. Up to 10 operators are supported currently.

## Troubleshooting

`tests/nodes/operators` contains mounted docker volumes. Sometimes it causes `permission denied` errors while running `docker ocmpose` and `make` commands. To fix this issue run `chmod -R 777 tests/nodes/operators`.

## Contract development

If either source code of AVS contracts is changed or new verison of `eigenlayer-middleware` is used, devnet testing environment must be updated by running `make deploy-all`. This will generate new [Anvil state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json), used by Anvil docker compose service.