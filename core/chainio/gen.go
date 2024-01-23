package chainio

//go:generate mockgen -destination=./mocks/avs_subscriber.go -package=mocks github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio AvsSubscriberer
//go:generate mockgen -destination=./mocks/avs_writer.go -package=mocks github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio AvsWriterer
//go:generate mockgen -destination=./mocks/avs_reader.go -package=mocks github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio AvsReaderer
