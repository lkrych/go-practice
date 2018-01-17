package main

import "context"

type contextKey struct {
	name string
}

var contextKeyAPIKey = &contextKey{"api-key"}

func main() {
}

func APIKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(contextKeyAPIKey).(string)
	return key, ok
}
