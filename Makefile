SOURSE_FILE := $(wildcard *.go) $(wildcard command/*.go)
TARGET := insgo

ifeq ($(OS), Windows_NT)
	EXTEXE := .exe
endif

.PHONY: debug
debug: insgo.exe
	@echo $(OS) $(EXTEXE)
	@./insgo install command/test.toml && echo "success"

$(TARGET)$(EXTEXE): $(SOURSE_FILE)
	go build
