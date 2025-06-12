# Go Command

A wrapper library for Cobra CLI framework.

## Install

```bash
go get -u github.com/InkShaStudio/go-command
```

## Usage

```go
package main

import (
  "fmt"
  "os"

  "github.com/InkShaStudio/go-command"
)

func main() {
  world := command.
    NewCommand("world").
    ChangeDescription("hello world").
    RegisterHandler(func(cmd *command.SCommand) {
      fmt.Println("Hello World")
    })

  cmd := command.RegisterCommand(
    command.
      NewCommand("hello").
      ChangeDescription("hello world").
      RegisterHandler(func(cmd *command.SCommand) {
      }).
      AddSubCommand(
        world,
      ),
  )
  if err := cmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
```
