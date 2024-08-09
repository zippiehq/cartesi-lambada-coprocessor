# Cartesi Lambada Coprocessor

This is proof of concept of verifiable computing system using  [Cartesi virtual machine](https://github.com/zippiehq/cartesi-lambada) and [EigenLayer restaking platform](https://github.com/Layr-Labs/eigenlayer-contracts).

## Running tests

Run `git submodule update --init --recursive`

To run tests on devnet (local deployment of EigenLayer and Lambada Coprocessor AVS) in first terminal run `docker compose up`. Then in second terminal run `make tests-integration`.

Use `docker compose -f docker-compose-holesky.yaml up` and `make tests-integration-holesky` to run tests on Holesky fork.

Use `docker compose -f docker-compose-mainnet.yaml up` and `make tests-integration-mainnet` to run tests on Mainnet fork.

Wait for tests to complete without any errors and check docker compose logs for unexpected errors.

## Configuring number of operators

Run `go run ./cli generate-docker-compose --operators $NUMBER_OF_OPERATORS` to set number of operators, used by testing environment. Up to 10 operators are supported currently.

## Troubleshooting

`tests/nodes/operators` contains mounted docker volumes. Sometimes it causes `permission denied` errors while running `docker compose` and `make` commands. To fix this issue run `chmod -R 777 tests/nodes/operators`.

## Contract development

Devnet testing environment is based on [local Anvil state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json), used by Anvil docker compose service. That state includes:

- Deployment of core EigenLayer contracts
- Deployment of AVS contracts

In case source code of AVS contracts is changed, AVS must be re-deployed to local Anvil state by running `make deploy-avs`.

In case new version of `eigenlayer-middleware` is used, appropriate version of EigenLayer contracts must be re-deployed by running `make deploy-eigenlayer` and then `make deploy-avs` (or by executing single `make deploy-all` command).