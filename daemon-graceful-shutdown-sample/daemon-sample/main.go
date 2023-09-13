package main

import (
  "os"
  "context"
  "fmt"
  "os/exec"
  "os/signal"
  "time"
)

func main() {
  ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
  defer stop()
  ctx, cancel := context.WithTimeout(ctx, time.Second*5)
  defer cancel()

  cmd := exec.CommandContext(ctx, "sleep", "10")
  err := cmd.Run()

  if err != nil {
    if ctx.Err() != nil {
      fmt.Printf("error: %v\n", ctx.Err())
    } else {
      fmt.Printf("error: %v\n", err)
    }
    os.Exit(1)
  }
}
