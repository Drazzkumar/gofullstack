// An example of creating a new instance of a Social Media Post type.
package main

import (
	"fmt"

	"github.com/razzkumar/gofullstack/volume1/section3/socialmedia"
)

func main() {

	myPost := socialmedia.NewPost("razzkumar", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
}
