SHELL := /bin/bash

PROJECT_DIR=$(shell pwd)
GRAMMAR_DIR=$(PROJECT_DIR)/internal/grammar
PARSER_DIR=$(PROJECT_DIR)/internal/parser

.PHONY: lexer

lexer: $(PARSER_DIR)/mini_c_lexer.go

$(PARSER_DIR)/mini_c_lexer.go: $(GRAMMAR_DIR)/MiniCLexer.g4
	@echo "Generate Lexer code"
	@cd $(GRAMMAR_DIR) && \
		antlr -Dlanguage=Go -o $(PARSER_DIR) -package parser MiniCLexer.g4