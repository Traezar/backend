# Makefile for executing SQL files

DB_USER := postgres
DB_NAME := golang-gorm
SQL_FILE := book.sql
DB_PASSWORD := password123

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: all build test clean run deps mod-tidy mod-vendor
all: test build

build:
		$(GOBUILD) -o $(BINARY_NAME) -v

test: 
		 go test -v  ./...

clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)

run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

deps:
		$(GOGET) -u ./...

mod-tidy:
		$(GOMOD) tidy

mod-vendor:
		$(GOMOD) vendor




fake_books:
	 psql -U $(DB_USER) -d $(DB_NAME) -f $(SQL_FILE)


