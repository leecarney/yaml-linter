package internal

/**
 * Lints a yaml file against a json schema spec file
 *
 * 	@param DataPath    String  Yaml file
 * 	@param SchemaPath  String  Json schema file
 * 	@param Validate    Func    Compares yaml and json
 *
 * 	@return Status string
 *
 * Example:
 * 		import lint "yaml-linter/cmd"
 *
 * 		v := lint.NewValidator
 * 		v.DataPath, v.SchemaPath = *data, *schema
 * 		v.Validate()
 *
 * 		Output: SUCCESS || ERROR
 */

import (
	"errors"
	"fmt"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

type Validator struct {
	DataPath   string
	SchemaPath string
}

const (
	SUCCESSYAML      = "SUCCESS ::::: Validated %v against given schema"
	ERRORYAML        = "ERROR ::::: Structurally Incorrect Yaml. %v"
	ERRORJSON        = "ERROR ::::: Unable to load Json Schema. %v"
	ERRORLOADINGFILE = "ERROR ::::: Unable to load file %v: %v"
	ERRORSTRINGKEYS  = "ERROR ::::: Found non-string key %v"
)

func (v *Validator) Validate() string {
	initLogger()
	res := doCompareYamlToJson(v.DataPath, v.SchemaPath)
	fmt.Printf(res)
	return res
}

func NewValidator() *Validator {
	return &Validator{}
}

func initLogger() {
	log.SetPrefix("yaml-linter: ")
	log.SetFlags(0)
}

func doCompareYamlToJson(dataPath string, schemaPath string) string {
	compiler := jsonschema.NewCompiler()

	if err := compiler.AddResource("schema.json", strings.NewReader(loadFileContents(schemaPath))); err != nil {
		log.Fatalf(ERRORJSON, err)
	}

	schema, err := compiler.Compile("schema.json")
	if err != nil {
		log.Fatalf(ERRORJSON, err)
	}

	if err := schema.Validate(getYaml(dataPath)); err != nil {
		log.Fatalf(ERRORYAML, err)
	}

	status := fmt.Sprintf(SUCCESSYAML, dataPath)
	return status
}

func getYaml(dataPath string) interface{} {
	var yamlString interface{}

	err := yaml.Unmarshal([]byte(loadFileContents(dataPath)), &yamlString)
	if err != nil {
		log.Fatalf(ERRORYAML, err)
	}

	yamlString, err = toStringKeys(yamlString)
	if err != nil {
		log.Fatalf(ERRORSTRINGKEYS, err)
	}
	return yamlString
}

func loadFileContents(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf(ERRORLOADINGFILE, path, err)
	}
	fileString := string(fileBytes)
	return fileString
}

// toStringKeys decodes a map and any nested maps it contains.
// Returns an Error if any key is not a string
func toStringKeys(val interface{}) (interface{}, error) {
	var err error
	switch val := val.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range val {
			k, ok := k.(string)
			if !ok {
				return nil, errors.New(ERRORSTRINGKEYS)
			}
			m[k], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return m, nil

	case []interface{}:
		var l = make([]interface{}, len(val))
		for i, v := range val {
			l[i], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		return val, nil
	}
}
