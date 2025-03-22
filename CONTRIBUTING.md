# How To Contribute

This is a library to share functions in Go across packages. Function that
are generic enough should fit right in here. If you're not sure what that is,
take a look around, hopefully it will become apparent.

## When To Make A Sub-Package

What is meant by sub-package here is putting related functionality into
subdirectories.

We do not want to fracture the code into too many packages, nor make one big
monolith. Try and group related functionality. If your not sure, then just
place them in the `stdlib` package. These things tend to work out over time.
Functionality can always be moved later. Yes, moving them into another package
later on will force a semantic major version increase. But don't worry
about that, just get it done, and we'll figure it out.

See [What to put into a package] for further assistance.

## Setup Local Development.

You can use a Docker environment to get going if you have Docker on your
computer. In fact, there is no documentation for any other way in this reading.

### Run Docker

1. Clone this repository.
2. In a command prompt, run: `docker compose up`
3. In another command prompt, login to the container:
   `docker exec -it stdlib_dev_1 sh`
4. Execute a command such as `go test`
   ```output
   ~/src/github.com/kohirens/stdlib $ go test
   PASS
   ok      github.com/kohirens/stdlib      0.004s
   ```

### Run with VS Code

1. Install the VS code extension "Remote Container".
2. Clone this repository locally.
3. Open the project in VS Code, which should ask to open the folder in a
   remote container.
4. Open a terminal in VS Code and type `go test`.
5. Now got to "Run and Debug" in the left nav and run one of the launch
   configurations.

---

[What to put into a package]: https://go.dev/blog/organizing-go-code#what-to-put-into-a-package