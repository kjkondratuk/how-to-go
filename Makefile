GO=go
PWD=$(shell pwd)
BIN_DIR=$(PWD)/bin
GO_FLAGS=-o $(BIN_DIR)/
EXAMPLES_DIR=$(PWD)/examples
ARGS=Chicago IL US

clean:
	rm -rf $(BIN_DIR)/*

make-bin-dir:
	-mkdir $(PWD)/bin

build-examples: | clean make-bin-dir
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/01_hello_world/hello_world.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/02_variable_declaration/variable_declaration.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/03_data_types/data_types.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/04_arrays_and_slices/arrays_and_slices.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/05_maps/maps.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/06_structs/structs.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/07_logic/logic.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/08_looping/looping.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/09_functions/functions.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/10_pointers/pointers.go
	$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/11_interfaces/interfaces.go
	cd $(EXAMPLES_DIR)/12_visibility_and_modules/ && \
		$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/12_visibility_and_modules/visibility_and_modules.go
	cd $(EXAMPLES_DIR)/13_testing/ && \
		$(GO) build $(GO_FLAGS) $(EXAMPLES_DIR)/13_testing/testing.go
	cd $(PWD)

build-app: | clean make-bin-dir
	$(GO) build $(GO_FLAGS) ./how-to-go.go

build-all: build-examples build-app

run-app: build-app
	$(BIN_DIR)/how-to-go $(ARGS)
