include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
SERVER=hedera-dev

############### LOCAL BUILD #################
# dev-mode
.phony: dev
dev: 
	$(COMPILEDAEMON_PATH) -command="./$(SERVER)"

# local run
run:
	go run .