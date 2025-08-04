PLANTUML_JAR ?= plantuml.jar

DIAGRAMS_DIR := ./diagrams

OUTPUT_DIR := $(DIAGRAMS_DIR)/generated

.DEFAULT_GOAL := all

.PHONY: all diagrams clean

all: diagrams

diagrams:
	@java -jar $(PLANTUML_JAR) -tpng "$(DIAGRAMS_DIR)/*.puml" "$(DIAGRAMS_DIR)/**/*.puml"

clean:
	@rm -rf $(OUTPUT_DIR)
