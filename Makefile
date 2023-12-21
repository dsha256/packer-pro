.DEFAULT_GOAL = help

ent-init: ## Inits Ent schemas.
	go run -mod=mod entgo.io/ent/cmd/ent new \
	--target internal/entity/schema \
	$(schema)

generate: ## Generates all scenarios.
	go generate ./...

help: ## Prints this message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: ALL
