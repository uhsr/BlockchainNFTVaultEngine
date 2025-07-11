// cmd/blockchainnftvaultengine/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftvaultengine/internal/blockchainnftvaultengine"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftvaultengine.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
