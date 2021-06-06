package config2

import (
	"os"
	"testing"
)

type Config struct {
	Host  string `usage:"Server hostname"`
	Port  int    `usage:"Server port"`
	Debug bool   `usage:"Enable debug"`
}

func TestParseCommandLine(t *testing.T) {
	args := []string{"-Debug", "-Host", "localhost", "-Port", "9090"}
	config := &Config{Port: 8080}

	flagSet := ParseCommandLine(args, config)

	if config.Host != "localhost" {
		t.Errorf("Host %s", config.Host)
	}

	if config.Port != 9090 {
		t.Errorf("Port %d", config.Port)
	}

	if !config.Debug {
		t.Errorf("Debug %t", config.Debug)
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

	if config.Host != "localhost" {
		t.Errorf("Host %s", config.Host)
	}

	if config.Port != 9090 {
		t.Errorf("Port %d", config.Port)
	}

	if !config.Debug {
		t.Errorf("Debug %t", config.Debug)
	}
}
