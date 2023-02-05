# How To Contribute

This is a library to share functions in Go across packages. Function that
are generic enough should fit right in here. If you're not sure what that is
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

---

[What to put into a package]: https://go.dev/blog/organizing-go-code#what-to-put-into-a-package