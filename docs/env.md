# env

This package contains functions for working with the environment.

## Functions

**Get(key, def string) string** - Get An environment variable falling back to
a default value when not set.

```go
package main

import "github.com/kohirens/stdlib/env"

func main() {
	// Are we running in a GitHub Action or some other CI
    env.Get("GITHUB_ACTIONS", "false")
}
```