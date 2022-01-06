package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func AuthTOken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func main() {
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userID string, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleRespone(ctx)
}

func HandleRespone(ctx context.Context) {
	fmt.Println("Handling respons for ", UserID(ctx), AuthTOken(ctx))
}
