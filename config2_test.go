package config2

import (
	"fmt"
	"os"
	"testing"
)

type Config struct {
	Host  string `usage:"Server hostname"`
	Port  int    `usage:"Server port"`
	Debug bool   `usage:"Enable debug"`
}

func checkTestValues(config *Config) error {
	if config.Host != "localhost" {
		return fmt.Errorf("Host %s", config.Host)
	}

	if config.Port != 9090 {
		return fmt.Errorf("Port %d", config.Port)
	}

	if !config.Debug {
		return fmt.Errorf("Debug %t", config.Debug)
	}

	return nil
}

func TestParseCommandLine(t *testing.T) {
	args := []string{"-Debug", "-Host", "localhost", "-Port", "9090"}
	config := &Config{Port: 8080}

	flagSet := ParseCommandLine(args, config)

	err := checkTestValues(config)
	if err != nil {
		t.Fatal(err)
	}

	flagSet.PrintDefaults()
}

func TestParseEnvironmentVariables(t *testing.T) {
	prefix := "test_"

	hostVariable := prefix + "Host"
	portVariable := prefix + "Port"
	debugVariable := prefix + "Debug"

	os.Setenv(hostVariable, "localhost")
	defer os.Unsetenv(hostVariable)

	os.Setenv(portVariable, "9090")
	defer os.Unsetenv(portVariable)

	os.Setenv(debugVariable, "true")
	defer os.Unsetenv(debugVariable)

	config := &Config{Port: 8080}

	ParseEnvironmentVariables(prefix, config)

	err := checkTestValues(config)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseJson(t *testing.T) {
	config := &Config{Port: 8080}

	err := ParseJson("test.json", config)
	if err != nil {
		t.Fatal(err)
	}

	err = checkTestValues(config)
	if err != nil {
		t.Fatal(err)
	}
}
