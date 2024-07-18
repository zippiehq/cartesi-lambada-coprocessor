############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deploy-eigenlayer: ## Deploy eigenlayer
	./tests/anvil/deploy-eigenlayer-save-anvil-state.sh

deploy-avs: ## Deploy avs
	./tests/anvil/deploy-avs-save-anvil-state.sh

deploy-all: deploy-eigenlayer deploy-avs 

bindings: ## generates contract bindings
	cd contracts && ./generate-go-bindings.sh

tests-contract: ## runs all forge tests
	cd contracts && forge test

tests-integration: ## runs integration tests on devnet
	go test ./tests/... -v -count=1 -args --network devnet

tests-integration-holesky: ## runs integration tests on Holesky fork
	go test ./tests... -v -count=1 -args --network holesky

tests-integration-mainnet: ## runs integration tests on Mainnet fork
	go test ./tests... -v -count=1 -args --network mainnet
