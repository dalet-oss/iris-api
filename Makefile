BIN_DIR = bin
CMD_DIR = cmd
DOCS_DIR = docs
MODELS_DIR = models

API_CLI_DIR = cli
API_CLI_NAME = irisctl

API_CLIENT_DIR = client
API_CLIENT_NAME = iris

API_SERVER_DIR = server
API_SERVER_OPERATIONS = api
API_SERRVERCLIENT_NAME = irisd

YQ = $(BIN_DIR)/yq
YQ_VERSION = v4.31.1

GOSWAGGER = $(BIN_DIR)/goswagger
GOSWAGGER_VERSION = v0.30.4
SWAGGER_YAML_TO_HTML = $(BIN_DIR)/swagger-yaml-to-html

OPENAPI_DIR = openapi
OPENAPI_DEFINITION = swagger.generated.yml
OPENAPI_DOC = $(DOCS_DIR)/index.html

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: api ; @
	$Q echo "done"

.PHONY: bin
bin: ; $(info $(M) create local bin directory) @
	$Q mkdir -p $(BIN_DIR)

.PHONY: get-yq
get-yq: bin; $(info $(M) downloading yq…) @
	$Q test -x $(YQ) || curl -sL https://github.com/mikefarah/yq/releases/download/$(YQ_VERSION)/yq_$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m | sed 's%x86_64%amd64%') --output $(YQ)
	$Q chmod a+x $(YQ)

.PHONY: swagger-specs
swagger-specs: get-yq ; $(info $(M) merge Swagger API fragments…) @
	$Q $(YQ) ea '. as $$item ireduce ({}; . * $$item )' $(OPENAPI_DIR)/*.yml > $(OPENAPI_DEFINITION)

.PHONY: get-goswagger
get-goswagger: bin; $(info $(M) downloading go-swagger…) @
	$Q test -x $(GOSWAGGER) || curl -sL https://github.com/go-swagger/go-swagger/releases/download/$(GOSWAGGER_VERSION)/swagger_$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m | sed 's%x86_64%amd64%') --output $(GOSWAGGER)
	$Q chmod a+x $(GOSWAGGER)

.PHONY: swagger-validate
swagger-validate: get-goswagger ; $(info $(M) validate Swagger API syntax…) @
	$Q $(GOSWAGGER) validate -q $(OPENAPI_DEFINITION)

.PHONY: swagger-generate-server
swagger-generate-server: get-goswagger ; $(info $(M) generate Swagger REST API Server Go code…) @
	$Q $(GOSWAGGER) generate server -q -f $(OPENAPI_DEFINITION) -s $(API_SERVER_DIR) -a $(API_SERVER_OPERATIONS) --exclude-main --name=$(API_SERVER_NAME)

.PHONY: swagger-generate-cli
swagger-generate-cli: get-goswagger ; $(info $(M) generate Swagger Client & CLI Go code…) @
	$Q $(GOSWAGGER) generate cli -q -f $(OPENAPI_DEFINITION) -c $(API_CLIENT_DIR) --name=$(API_CLIENT_NAME) --cli-app-name=$(API_CLI_NAME)

.PHONY: get-swagger-yaml-to-html
get-swagger-yaml-to-html: bin; $(info $(M) downloading swagger-yaml-to-html…) @
	$Q test -x $(SWAGGER_YAML_TO_HTML) || curl -sL https://raw.githubusercontent.com/yousan/swagger-yaml-to-html/master/swagger-yaml-to-html.py --output $(SWAGGER_YAML_TO_HTML)
	$Q chmod a+x $(SWAGGER_YAML_TO_HTML)

.PHONY: swagger-doc
swagger-doc: get-swagger-yaml-to-html ; $(info $(M) generate Swagger API HTML documentation…) @
	$Q mkdir -p $(shell dirname $(OPENAPI_DOC))
	$Q python3 $(SWAGGER_YAML_TO_HTML) < $(OPENAPI_DEFINITION) > $(OPENAPI_DOC)

.PHONY: api
api: swagger-specs swagger-validate swagger-doc swagger-generate-server swagger-generate-cli ; @

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@
	$Q rm -rf $(BIN_DIR)
	$Q rm -rf $(CMD_DIR)
	$Q rm -rf $(OPENAPI_DOC)
	$Q rm -rf $(OPENAPI_DEFINITION)
	$Q rm -rf $(DOCS_DIR)
	$Q rm -rf $(MODELS_DIR)
	$Q rm -rf $(API_CLI_DIR)
	$Q rm -rf $(API_CLIENT_DIR)
	$Q rm -rf $(API_SERVER_DIR)
