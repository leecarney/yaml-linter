package main

import (
	"errors"
	"flag"
	"os"
	lint "yaml-linter/internal"
)

var (
	data   = flag.String("data", "", "--data [data file path to validate]")
	schema = flag.String("schema", "", "--schema [schema file path to use]")
)

func main() {
	checkInputs()
	v := lint.NewValidator()
	v.DataPath, v.SchemaPath = *data, *schema
	v.Validate()
}

func checkInputs() {
	flag.Parse()
	if *data == "" || *schema == "" {
		panic("ERROR ::::: Please specify all required parameters. Schema & Data are required.")
	}
	if !exists(*data) {
		panic("ERROR ::::: Data file does not exist")
	}
	if !exists(*schema) {
		panic("ERROR ::::: Schema file does not exist")
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
