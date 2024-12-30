package main

import (
	"errors"
	"fmt"
	"github.com/junkyard/internal/server"
	"net/http"
	"os"
)

func main() {

	server.RegisterHandlers()

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
