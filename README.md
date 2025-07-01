# ginpacifier

A simple Gin middleware for Go that gracefully recovers from panics, ensuring your Gin server does not crash and returns a 500 error if a panic occurs.

## Features
- Recovers from panics in handlers and middleware
- Returns a JSON 500 error if the response is not already written
- Does nothing if the response is already written
- Easy to use and test

## Installation

```
go get github.com/neighbor/ginpacifier
```

## Usage

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/neighbor/ginpacifier"
)

func main() {
    r := gin.New()
    r.Use(ginpacifier.PanicRecovery())
    // ... your routes
    r.Run()
}
```

## Testing

Run all tests with:

```
go test -v
```