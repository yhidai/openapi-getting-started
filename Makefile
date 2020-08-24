generate-client:
	@docker run -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate -i /local/openapi.yaml -g go -o /local/out/go/client --package-name client --git-user-id yhidai --git-repo-id openapi-getting-started/client
	sudo chown -R $(whoami):$(whoami) ./out/
	cp -rf ./client/example ./
	cp -rf ./out/go/client ./
	cp -rf ./example ./client

generate-server:
	@docker run -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate -i /local/openapi.yaml -g go-server -o /local/out/go/server --package-name server --git-user-id yhidai --git-repo-id openapi-getting-started/server
	sudo chown -R $(whoami):$(whoami) ./out/
	cp -rf ./out/go/server ./

copy-generated-server:
	cp -rf ./out/go/server ./

config-help:
	@docker run -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 config-help -g go