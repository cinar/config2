# Config2 Go

Config2 is a lightweight Golang module for managing and populating application configuration from JSON, command line, and environment variables.

## Usage

Install package as shown below.

```bash
go get github.com/cinar/config2
```

Import Config2 as shown below.

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

Config2 provides the following field tags for you to be able to define additional information for each configuration variable.

Tag | Description | Example
--- | --- | ---
usage | Usage for command line help. | `usage:"Server hostname"

Define an instance of the configuration structure.

```Golang
config := &Config{}
```

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

Use the ParseCommandLine function to parse the command line arguments as shown below.

```Golang
flagSet := ParseCommandLine(os.Args, config)
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

Use the ParseEnvironmentVariables function to parse the environment variables. The function takes a prefix for the environment variables. Please set it to empty string oif no prefix needed. 

```Golang
ParseEnvironmentVariables("test_", config)
```

## License

The source code is provided under [MIT License](LICENSE).
