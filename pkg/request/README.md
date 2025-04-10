## Request utils

Simple utility function that wraps common (for my use case) http fetch functionality for creating http requests.

### Usage 

The package exports a single `MakeRequest` function. This function takes a url, http method, a map of headers and a body in the type of a `[]byte`.

### Examples

```go
type responseBody struct {
	// ...
}

func example() {
	body, responseCode, err := MakeRequest[responseBody]("localhost:3000/home", "GET", map[string]string{}, []byte{});
	if err != nil {
		... // Handle errors
	}

	if responseCode != 200 {
		... // Handle response codes
	}
}
```
