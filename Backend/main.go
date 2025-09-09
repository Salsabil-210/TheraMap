package main

import (
    "Backend/infrastructure"  // MUST match "backend" from go.mod
)

func main() {
    infrastructure.ConnectDB()
}