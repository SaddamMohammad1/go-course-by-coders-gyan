package main

import (
	"fmt"
	"net/url"
)

func main() {

	/* Parse() breaks URL into scheme, host, path, query, fragment etc. */
	myurl := "https://example.com:3000/learn?course=golang&level=beginner"

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", result.Scheme)      // https
	fmt.Println("Host:", result.Host)          // example.com:3000
	fmt.Println("Path:", result.Path)          // /learn
	fmt.Println("Port:", result.Port())        // 3000
	fmt.Println("Raw Query:", result.RawQuery) // course=golang&level=beginner

	/* Query() converts the query part into a map */
	queryParams := result.Query()
	fmt.Println("Course:", queryParams.Get("course")) // Course: golang
}
