package config2

import "testing"

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
