# vars - Config Variable Management
**vars** is used to manage configuration using environment variables. This project includes methods for retrieving mandatory and optional environment variables, which provide a solid solution for configuration management in your Go applications.

## Mandatory Variables
Functions that return the value of a specified environment variable. If the environment variable is not set, these functions will return an error.

### Usage
An example of how to use the `MandatoryString` function:

```go
package main

import (
    "fmt"
    "log"
    "github.com/loopinnovates/vars" 
)

func main() {
    varName := "MY_MANDATORY_VARIABLE"
    result, err := vars.MandatoryString(varName)
    if err != nil {
        log.Fatalf("Failed to get the value of %s: %v", varName, err)
    }
    fmt.Println(result)
}
```

## Optional Variables
In this, Functions will return a default value if the specified environment variable is not set.

### Usage
An example of how to use the `OptionalString` function:

```go
package main

import (
    "fmt"
    "github.com/loopinnovates/vars"
)

func main() {
    varName1 := "MY_STRING_VARIABLE"
    defaultValue1 := "default"
    result1 := vars.OptionalString(varName1, defaultValue1)
    fmt.Println(result1)
}
```