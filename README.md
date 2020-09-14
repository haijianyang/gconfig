## Gconfig
Go Application Configuration. Simple and practical configuration management.

Inspired by [node-config](https://github.com/lorenwest/node-config) and [viper](https://github.com/spf13/viper), Gconfig supports some features of both: parse different configuration files and environment variable according to ENV. Currently only supports JSON file。

* Support both configuration files and environment variables
* Support multiple environments：development、staging、production etc
* Support multiple configuration instances

## Install

```console
go get github.com/haijianyang/gconfig
```

## Quick Start

```go
type Config struct {
}

var config Config

err := gconfig.Unmarshal(&config)
if err != nil {
	fmt.Println(err)
}

fmt.Println(config)
```

## Document

### Setting Overrides

```go
gconfig.SetDefault(key, value) // Set default configuration

gconfig.SetDefault("Env", "development") // Set Go running environment
gconfig.SetDefault("Folder", "./config") // Set the configuration file folder
gconfig.SetDefault("FileType", ".json") // Set the type of configuration file
gconfig.SetDefault("DefaultFile", "default") // Set default configuration file name
```

### Working with Environment Variables

Gconfig uses ENV and GO_ENV environment variables to select configuration files by default. For example: ENV=test will use test.json。

Environment variables take precedence over file variables, and the value of environment variables will override the value of file variables.

Gconfig parses environment variables according to the env tag in struct.

```go
type EnvConfig struct {
	token bool `json:"token" env:"TOKEN"`
}
```

## Best Practices
Create a config folder in the project root directory, and put the configuration file and config.go in the config folder.

```js
- project
  - config
    - config.go
    - default.json
    - development.json
    - production.json
    - staging.json
    - test.json
```

default.json

```js
{
  "server": {
    "port": 8000
  }
}
```

test.json
```js
{
  "server": {
    "port": 8000
  }
}
```

config.go

```go
package config

import (
	"fmt"
	"github.com/haijianyang/gconfig"
)

type ServerConfig struct {
	Port int `json:"port"`
}

type Config struct {
	Server ServerConfig
}

var Kvs Config

func init() {
	err := gconfig.Unmarshal(&Kvs)
	if err != nil {
		fmt.Println(err)
	}
}
```

code.go

```go
import (
	"fmt"

	"github.com/haijianyang/project/config"
)

fmt.Println(config.Kvs.Server)
```
