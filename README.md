[![GoDoc](https://godoc.org/github.com/cinar/config2?status.svg)](https://godoc.org/github.com/cinar/config2)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.com/cinar/config2.svg?branch=master)](https://travis-ci.com/cinar/config2)

# Config2 Go

Config2 is a lightweight Golang module for managing and populating application configuration from JSON, command line, and environment variables.

## Usage

Install package.

```bash
go get github.com/cinar/config2
```

Import Config2.

```Golang
import (
  "github.com/cinar/config2
)
```

Define a configuration structure as shown below.

```Golang
type Config struct {
	Host  string `usage:"Server hostname"`
	Port  int    `usage:"Server port"`
	Debug bool   `usage:"Enable debug"`
}
```

Config2 provides the following field tags to define additional information for configuration variables.

Tag | Description | Example
--- | --- | ---
usage | Usage for command line help. | `usage:"Server hostname"

Set the default values for the configuration variables.

```Golang
config := &Config{Port: 8080}
```
### Read from command line arguments

Config2 can automatically generate a FlagSet for your configuration structure and parse the command line arguments.

Application can be launched with command line arguments.

```
./main -Debug -Host localhost -Port 9090
```

Use the [ParseCommandLine](https://pkg.go.dev/github.com/cinar/config2#ParseCommandLine) function to parse the command line arguments as shown below.

```Golang
flagSet := config2.ParseCommandLine(os.Args, config)
```

Validate the configuration variables. In case of an error, show the usage through the returned FlagSet as shown below.

```Golang
flagSet.PrintDefaults()
```

The command line arguments, their usage, and their defaults will be shown by the FlagSet as usual.

```
  -Debug
        Enable debug (default false)
  -Host
        Server hostname
  -Port
        Server port (default 8080)
```
### Read from environment variables

Config2 can read configuration from the environment variables.

Application can be launched with the environment variables set as shown below.

```bash
export test_Host=localhost
export test_Port=9090
export test_Debug=true

./main
```

Use the [ParseEnvironmentVariables](https://pkg.go.dev/github.com/cinar/config2#ParseEnvironmentVariables) function to parse the environment variables. The function takes a prefix for the environment variables. Please set it to empty string ("") if no prefix needed. 

```Golang
config2.ParseEnvironmentVariables("test_", config)
```

### Read from JSON file

Config2 can read configuration from a JSON file as well, such as the one below.

```JSON
{
  "Host": "localhost",
  "Port": 9090,
  "Debug": true
}
```

Use the [ParseJson](https://pkg.go.dev/github.com/cinar/config2#ParseJson) function to parse the JSON file.

```Golang
err := config2.ParseJson("test.json", config)
if err != nil {
  log.Fatal(err)
}
```

### Parse all

Config2 can parse the JSON file if exists, parse the environment variables, and parse the command line arguments.

Use the [ParseAll](https://pkg.go.dev/github.com/cinar/config2#ParseAll) function as shown below.

```Golang
config2.ParseAll("test.json", "test_", config)
```

## License

The source code is provided under MIT License.
