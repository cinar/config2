# Config2 Go

Config2 is a lightweight Golang module for managing and populating application configuration from JSON, command line, and environment variables.

## Usage

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
- | - | -
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