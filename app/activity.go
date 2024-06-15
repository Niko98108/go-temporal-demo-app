package app

import (
	"context"
	"fmt"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello, %s! Welcome to temporal, ", name)
	return greeting, nil
}
