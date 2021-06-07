package config2

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

const (
	usageTag = "usage"
)

// Check config to be a pointer to struct.
func checkConfig(config interface{}) {
	configType := reflect.TypeOf(config)
	if configType.Kind() != reflect.Ptr {
		log.Fatal("Config not a pointer")
	}

	configValueType := configType.Elem()
	if configValueType.Kind() != reflect.Struct {
		log.Fatal("Config not a pointer to struct")
	}
}

// Parse command line arguments.
func ParseCommandLine(args []string, config interface{}) *flag.FlagSet {
	checkConfig(config)

	configType := reflect.TypeOf(config).Elem()
	configValue := reflect.ValueOf(config).Elem()

	flagSet := flag.NewFlagSet(configType.Name(), flag.ExitOnError)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		value := configValue.Field(i)
		usage := field.Tag.Get(usageTag)

		fv := &fieldValue{
			Kind:  field.Type.Kind(),
			Value: value,
		}

		flagSet.Var(fv, field.Name, usage)
	}

	flagSet.Parse(args)

	return flagSet
}

// Parse the environment variables.
func ParseEnvironmentVariables(prefix string, config interface{}) {
	checkConfig(config)

	configType := reflect.TypeOf(config).Elem()
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)

		value, ok := os.LookupEnv(prefix + field.Name)
		if ok {
			fv := &fieldValue{
				Kind:  field.Type.Kind(),
				Value: configValue.Field(i),
			}

			setFieldValue(fv, value)
		}
	}
}

// Parse configuration variables from JSON file.
func ParseJson(fileName string, config interface{}) error {
	checkConfig(config)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, config)
}

// Parse the JSON file if exists, parse the environment variables, and parse the command line arguments.
func ParseAll(fileName string, prefix string, config interface{}) {
	ParseJson(fileName, config)
	ParseEnvironmentVariables(prefix, config)
	ParseCommandLine(os.Args, config)
}
