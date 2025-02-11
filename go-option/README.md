# go-option

Standard functional options.

## Dependencies
None

## Installation
```shell
go get github.com/ahawker/go-stdlib/go-option
```

## Usage

The package offers functional options for managing both reference types and value types.

### Reference Types
```go
package main

import (
    "github.com/ahawker/go-stdlib/go-option"
    "os"
)

type Config struct {
    Adapter string
    RPSLimit int
}

func WithAdapter(adapter string) option.Opt[*Config] {
    return func(c *Config) error {
        c.Adapter = adapter
        return nil
    }
}

func WithRPSLimit(limit int) option.Opt[*Config] {
    return func(c *Config) error {
        c.RPSLimit = limit
        return nil
    }
}

func main() {
    // Use `Make` to create a zero value of the reference type with functional options.
    c1, err := option.Make[*Config]()
    // &Config{Adapter: "", RPSLimit: 0}
    c2, err := option.Make(WithRPSLimit(1024))
    // &Person{Adapter: "", RPSLimit: 1024}

    // Use `Apply` to mutate an existing reference type with functional options.
    c3, err := option.Apply(&Config{})
    // &Person{Name: "", Role: ""}
    c4, err := option.Apply(&Config{Adapter: "sqlite"}, WithRPSLimit(50))
    // &Person{Name: "sqlite", Role: 50}
}
```

### Value Types
```go
type Employee struct {
    Name string
    Role string
}

func WithName(name string) option.Val[Employee] {
    return func(e Employee) error {
        e.Name = name
        return e, nil
    }
}

func WithRole(role string) option.Val[Employee] {
    return func(e Employee) error {
        e.Role = role
        return e, nil
    }
}

func main() {
    // Use `New` to create a zero value of the value type with functional options.
    p1, err := option.New[Person]()
    // &Person{Name: "", Role: ""}
    p2, err := option.New(WithName("bob"))
    // &Person{Name: "bob", Role: ""}

    // Use `Copy` to copy an existing value type with functional options.
    p3, err := option.Copy(Person{})
    // &Person{Name: "", Role: ""}
    p4, err := option.Copy(Person{Role: "default"}, WithName("bob"))
    // &Person{Name: "bob", Role: "default"}
}
```


## License

[Apache 2.0](../LICENSE)

