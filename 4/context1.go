package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("jane","abc123")
}

func ProcessRequest(userID string, authToken string) {
	ctx := context.WithValue(context.Background(),"userID",userID)
	ctx = context.WithValue(ctx, "authToken",authToken)
	HandleRespone(ctx)
}

func HandleRespone(ctx context.Context) {
	fmt.Println("Handling respons for ",ctx.Value("userID"),ctx.Value("authToken"))
}
