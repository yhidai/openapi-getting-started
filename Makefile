UNAME = $(shell whoami)
generate-client:
	@docker run -v ${PWD}:/local -u $(id -u):$(id -g) openapitools/openapi-generator-cli:v4.3.1 generate -i /local/openapi.yaml -g go -o /local/out/go/client --package-name client --git-user-id yhidai --git-repo-id openapi-getting-started/client
	@sudo chown -R $(UNAME):$(UNAME) ./out/
	@cp -rf ./client/example ./
	@cp -rf ./out/go/client ./
	@cp -rf ./example ./client
	@echo "client generated."

generate-server:
	@docker run -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate -i /local/openapi.yaml -g go-server -o /local/out/go/server --package-name server --git-user-id yhidai --git-repo-id openapi-getting-started/server
	@sudo chown -R $(UNAME):$(UNAME) ./out/
	@cp -rf ./out/go/server ./
	@echo "server generated."

update-server:
	@docker run -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate -i /local/openapi.yaml -g go-server -o /local/out/go/server --package-name server --git-user-id yhidai --git-repo-id openapi-getting-started/server
	@sudo chown -R $(UNAME):$(UNAME) ./out/
	@cp -rf ./out/go/server/go ./server
	@echo "server updated."

clear:
	@rm -rf ./out
	@echo "clear out directory."

config-help:
	@docker run -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 config-help -g go