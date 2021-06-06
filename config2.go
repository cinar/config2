package config2

import (
	"flag"
	"log"
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
