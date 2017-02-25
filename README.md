# Go Walkthrough: Ordinals
I find I have a need for a simple command line tool that will take in a number,
and spit out the English ordinal for that number. I'm using Go 1.8.0.

# Rules
No third party libraries! We're just using the standard library here. No reason
we need anything beyond that.

# Version 1
Let's dissect this a bit:
 - `package main` declares our namespace, "main". `main` is a special package,
 in that it contains the entry point to an executable, and that it can't be
 referenced from any other package.
 - `import` is how we get other namespaces (called "packages" in go) into our
 file. Here everything we're importing is from the standard library.
  - `fmt` provides string formatting functions, like the familiar sprintf and
  friends.
  - `os` provides platform-specific stuff, we're just using it to get command
  line arguments and send exit codes.
  - `strconv` allows converting between strings and other types - in our case,
  integers.
 - `func main()` is the entry point to our executable; it's what gets called
 when we run the program. It has no arguments and no return values.
 - `if` statements, like most flow control in Go, don't use parentheses.
 - `len` is a language construct, not a function, and it gives the length of an
 array, slice, string, or map. Here we're using it to get the length of a slice.
 - A slice is a variable-length reference to some part of an underlying
 fixed-length array. Slices and arrays in Go are zero-indexed.
 - `os.Args` is referencing a global variable (`Args`) inside the `os` package.
 This variable is a slice of strings, where each element is a command-line
 argument, beginning with the command name.
 - `fmt.Fprint` writes to an output stream, here `os.Stderr`, the standard error
 output.
 - `os.Exit` exits the program with an exit code.
 - Line 14 introduces a few concepts:
  - `n, err`: here we're assigning values to two variables at once, because the
  function we're calling, `strconv.Atoi`, has two return values.
  - `:=` is the short-assignment operator, it declares a new local variable and
  sets its value at the same time. Go is statically-typed, but with a short
  assignment, it infers the type of the variable from the type of the value.
  - `strconv.Atoi` converts a string to an integer, returning an integer and
  and error. If it succeeds, the error will be `nil`.
  - `os.Args[1]` contains the first argument to our command; `os.Args[0]` is the
  name of the command itself.
 - `if err != nil` is a very common construct in Go, being the standard error
 handling construct. If the returned error is not nil, something went wrong,
 and we should deal with it; otherwise, we carry on as normal.
 - `switch n % 10` is a switch statement like any other language, without the
 parentheses. `%` is the modulo operator. So, we're switching based on the last
 digit of the input.
 - `case` statements do not fall through.

So, what does it do?
 1. Check if the user provided exactly one argument, and if not, print an
 error message and exit.
 1. Try converting the argument to an integer, and if we can't, print an error
 message and exit.
 1. If the last digit is a 1, print "st"; if it's a 2, print "nd"; if it's a 3,
 print "rd"; otherwise, print "th".
 1. Exit successfully.

# Usage
We can do a quick compile-and-run test:
```
> go run main.go 1
st
```

`go run` takes a source file, compiles it, and executes it, without generating
a binary file.

We can build it and get our binary file:
```
> go build

> ./gordinals 1
st
> ./gordinals
You must pass exactly one integer argument.
> ./gordinals potato
potato is not an integer.
```

`go build` compiles the current package and outputs an executable into the
current working directory.

We can also install it:
```
> go install

```

This compiles the current package and puts the binary in %GOPATH%/bin. If that
is in our $PATH, the executable is immediately accessible.

# Sharing
Actually, maybe other people would want to use this, too. Let's cross compile it.

```
> env GOOS=linux GOARCH=amd64 go build -o gordinals_x64
> env GOOS=linux GOARCH=386 go build -o gordinals_x86
> env GOOS=linux GOARCH=arm go build -o gordinals_arm
> env GOOS=windows GOARCH=386 go build -o gordinals32.exe
> env GOOS=windows GOARCH=amd64 go build -o gordinals64.exe
```

Now we have several binaries in our working directory, each for a different OS
and architecture combination. Go supports a long (and growing) list of operating
systems and architectures.
