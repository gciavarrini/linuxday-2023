package function

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */

	fmt.Println("Received request")
	fmt.Println(prettyPrint(req))      // echo to local output
	fmt.Fprintf(res, prettyPrint(req)) // echo to caller
}

func prettyPrint(req *http.Request) string {
	// Generate a random number
	randomNumber := rand.Intn(42)

	return fmt.Sprintf("Here it is a random number for you %d", randomNumber)
}
