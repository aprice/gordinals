# Go Walkthrough: Ordinals
I find I have a need for a simple command line tool that will take in a number,
and spit out the English ordinal for that number. I'm using Go 1.8.0.

# Rules
No third party libraries! We're just using the standard library here. No reason
we need anything beyond that.

We'll follow basic best practices, to the point that it makes sense for such a
small project, and doesn't serve to make it overly difficult to follow.

# Version 1.1
What if someone wants to embed our function in another program? Let's make it
reusable as a library.

First, we set up a new directory structure:
 - `/ordinal` will hold our logic, so that people can import it as a library.
 - `/cmd` will hold our executable(s).
  - `/cmd/ordinal` will hold the main package for the command-line executable we
  created for version 1.

First, we move our logic into `/ordinal/ordinal.go`. This necessitates some
minor changes:
 - `package ordinal` puts our logic in a different namespace, called `ordinal`.
 This matches the directory it's in, which isn't required, but is generally a
 best practice. Our package will be imported by its path and referenced by its
 package name, so keeping them consistent makes it easier to use.
 - `func For(n int) string` looks pretty different:
  - `For` is the name of our function. Capitalizing the first letter *exposes*
  the function outside our package. Anything with a lower-case first letter can
  only be referenced inside the same package.
  - `n int` is our single argument, called `n`, which is an integer. Type
  declarations go *after* variable names in Go.
  - `string` is our return type; like with variable declarations, return type
  declarations in Go are put *after* the function name.
 - Instead of `fmt.Print`ing our results, we instead `return` them to the
 caller.

 Then, we move `main.go` into `/cmd/ordinal/`. It is a common practice in Go to
 have packages that can be used as libraries in the same project with one or
 more executables. In order to allow multiple executables in the same project,
 we give each executable its own directory; in order to keep things clear, we
 put those directories inside a parent directory, so they aren't confused with
 library packages to be imported.

 Little has changed:
  - We now import `github.com/aprice/gordinals/ordinal`, which contains our
  `ordinal.go` file created above. This full repository reference both serves
  to indicate that this is a third-party library (and where it came from) as
  well as allowing Go to automatically download it if it isn't found locally.
  - We now call `ordinal.For` to handle our logic. `ordinal` is the package name
  (see above), and `For` is our function name. This is exactly how our other
  references to `os`, `strconv`, and `fmt` work - in fact, you can look at the
  Go source for those packages to see, the standard library works exactly like
  our own code.

# Usage
Nothing has really changed, but to build it, we now need to build `cmd/ordinal`,
because that's where our `main` package lives now:
```
> go build ./cmd/ordinal

> ordinal 1
st
```

Also, at this point, anyone could reference our package in their own project
by importing `github.com/aprice/gordinals/ordinal`, and they could install it
by running `go get github.com/aprice/gordinals/cmd/ordinal`. The latter would
download our source, compile it, and install it under `$GOPATH/bin`.
