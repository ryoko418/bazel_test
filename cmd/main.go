package main

import (
    "log"

    "github.com/ryoko418/bazel_test/uuid"
)

func main() {
    id, err := uuid.Generate()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(id)
}
