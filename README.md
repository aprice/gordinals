# Go Walkthrough: Ordinals
I find I have a need for a simple command line tool that will take in a number,
and spit out the English ordinal for that number. I'm using Go 1.8.0.

# Rules
No third party libraries! We're just using the standard library here. No reason
we need anything beyond that.

We'll follow basic best practices, to the point that it makes sense for such a
small project, and doesn't serve to make it overly difficult to follow.

# Version 1.1.1
Looks like we aren't handling teens correctly... Let's add some unit tests. And
some benchmarks, while we're here. We'll be using some best practices that
introduce several new concepts, so don't feel bad if you have to read this
section a couple of times or reference the Go documentation to make sense of it.
I'll do my best to explain as we go. In the course of this section we'll briefly
cover structs, anonymous functions, and closures. If these are totally alien
concepts to you, it might benefit you to take a quick look at the Go
walkthrough.

## Tests
In Go, unit tests live alongside the code they test, in files named `*_test.go`.
Files named this way will only be built when running tests, they won't be
included when we build our binaries. We'll put our tests in
`ordinal/ordinal_test.go`.
 - Every test file will import `testing`, which is Go's standard testing library.
 - Tests are in functions called `Test*`, which must take a single `*testing.T`
 argument and have no return. Each such function will be run when we run tests.

We're using a very common practice called *table-driven tests*. At the start of
our test function, we build a collection of test cases, using an anonymous
struct type to define them. The syntax looks a little strange, so let's break it
down.
 - `tests := []struct{` defines a variable `tests`, which will contain a *slice
 of structs*. A slice is a variable-length array, and a struct is much like
 other languages, a custom-made data structure. Typically structs are defined
 with names, but here we're using an *anonymous* struct, which we define in
 place.
 - Lines 10 and 11 define our struct's *fields*, the data it holds. Here we
 have two fields, `in`, our integer input value, and `expected`, our string
 expected output.
 - We then provide a list of values for our slice, with each value being an
 instance of our anonymous struct. We could use `in: 0, expected: "th"`,
 explicitly giving the field names, but it's more concise to just specify them
 in the order they're given in the struct definition.
 - You'll notice on line 24 there is a trailing comma. This may seem strange,
 but **it is not optional**.
 - Having listed off our test cases, we then iterate over them with a `for`
 loop. Here we're using `range`, which lets us iterate over an array, slice,
 map, or channel, executing the loop once for each value.
 - `range` returns two values for each iteration: the index, and the value. We
 use `_, test` to *ignore* the index, taking only the value. `_` can be used
 in various places to ignore a value or import.
 - Inside our loop, we call `t.Run`, which helps to organize test results. We
 could just test each case one by one, issuing an error when one fails, but then
 execution would stop. We want to run through *all* the test cases, then find
 out which fail and which pass. So, we use `t.Run` to separate execution of each
 test case.
 - `t.Run` takes a name string, which will be used to identify this test case in
 the output, and a function to run the test, which has a familiar signature: a
 single argument of type `*testing.T`, and no return values, just like our
 `TestFor` function itself.
 - Here, we're using an *anonymous* function, which, like our anonymous struct
 above, is defined in place, right there in the call to `t.Run`. Note that this
 function calls its argument `tt`, to distinguish it from the outer function's
 `t` argument. We need to do this because...
 - ... Our anonymous function is a *closure*. Because it is defined inside
 another function, it *closes over* the variables in the outer function. Because
 of this, we get to use the `test` variable from our `for` loop inside our
 anonymous test function.
 - Our test function itself is pretty simple:
  - Call the function under test with the input from the current test case, and
  store the output in a variable called `actual`.
  - If the `actual` value is not equal to our test case's `expected` value, the
  test fails, and we output a detailed error message. By calling `tt.Errorf`
  instead of, for example, `fmt.Printf`, we not only print the error message,
  but we also tell Go's testing package that this particular test has failed.

Phew! There's a lot there, I know, but trust me: once you've written a couple of
tests, it becomes second nature, and this style of testing is incredibly useful.
It keeps tests organized and gives valuable output when tests fail. Let's try it:

```
go test ./...
?       github.com/aprice/gordinals/cmd/ordinal [no test files]
--- FAIL: TestFor (0.00s)
    --- FAIL: TestFor/n=11 (0.00s)
        ordinal_test.go:31: For(11): expected th, got st
    --- FAIL: TestFor/n=12 (0.00s)
        ordinal_test.go:31: For(12): expected th, got nd
    --- FAIL: TestFor/n=13 (0.00s)
        ordinal_test.go:31: For(13): expected th, got rd
    --- FAIL: TestFor/n=111 (0.00s)
        ordinal_test.go:31: For(111): expected th, got st
    --- FAIL: TestFor/n=1011 (0.00s)
        ordinal_test.go:31: For(1011): expected th, got st
FAIL
FAIL    github.com/aprice/gordinals/ordinal     0.092s
```

Oh no! `FAIL` doesn't sound good. You just had to put it in all caps to rub it
in, didn't you, Go? So, what is this telling us?
 - `go test ./...` says "run tests for the current directory, and every
 directory under it, recursively". Many Go commands use this `./...` syntax the
 same way.
 - `? github.com/aprice/gordinals/cmd/ordinal [no test files]` - it found a
 package, but doesn't know if it passed or failed, because there were no tests.
 Fair enough.
 - `FAIL: TestFor`: our `TestFor` test function did not pass.
 - `FAIL: TestFor/n=11`: inside `TestFor`, the test case called `n=11` (which is
 the name passed to `t.Run`) did not pass. The next line shows the specific
 error output.
 - Several more tests also failed.
 - `FAIL`: The test run overall failed.
 - `FAIL github.com/aprice/gordinals/ordinal` our package overall failed.

Well, we can't have that, now can we? Let's fix it, it `ordinal/ordinal.go`, by
adding a check for teens:
```go
mod := n % 100
if mod >= 4 && mod <= 20 {
    return "th"
}
```

We re-run our tests, and voila:
```
$ go test ./...
?       github.com/aprice/gordinals/cmd/ordinal [no test files]
ok      github.com/aprice/gordinals/ordinal
```

Oh, sure, when it doesn't pass you say `FAIL` but when it does you just say `ok`?
That's just cold, Go.

## Benchmarks
We said we'd add benchmarks too, and we will, in our same `ordinal_test.go`
file. Benchmarks are named `Benchmark*`, and take a single argument of type
`*testing.B`, with no return value. There are two basic types of benchmarks. By
default, benchmarks are run single-threaded. However, if you use `RunParallel`,
they are run - shockingly - in parallel, allowing you to test multi-threaded
performance. In our case it hardly matters, but you can at least see how it
works.

First, we wrap all our code up in an anonymous function passed to
`b.RunParallel`, which is what makes the benchmarks run in parallel. This
anonymous function takes a single argument of type `*testing.PB` (*not* just
`testing.B`), which handles Parallel Benchmarks. In a parallel benchmark, each
execution is started with `PB.Next()`. It's important that each execution runs
the same, so we pass the same argument to our `For` function every time. This is
because Go benchmarks are run with more and more iterations until the execution
time stabilizes, so if we used incremental or random input values, it might
never stabilize and the benchmark would never complete (actually, it would fail
with an error that it took too long).

Let's run our benchmark:
```
> go test -run ^$ -bench .* ./...
```

Wait, what? Sorry, let me explain. `go test` allows specifying what tests or
benchmarks to run using regular expressions. We pass in `-run ^$`, which matches
nothing, so that no tests are run; we only want to run our benchmarks right now.
We pass in `-bench .*` to run all benchmarks. The `./...` works the same as
before. Right, back to it:

```
> go test -run ^$ -bench .* ./...
?       github.com/aprice/gordinals/cmd/ordinal [no test files]
BenchmarkFor-8          1000000000               1.59 ns/op
PASS
ok      github.com/aprice/gordinals/ordinal     1.888s
```

Right, no benchmarks in the command package, that's fine. Our `BenchmarkFor` ran
with 8 threads, executed 1,000,000,000 iterations, and came up with a total time
of 1.59 nanoseconds per execution. I guess that's acceptable performance. Let's
check one more thing, though:

```
> go test -run ^$ -bench .* -cpu 1,2,4,8,16 ./...
?       github.com/aprice/gordinals/cmd/ordinal [no test files]
BenchmarkFor            200000000               10.1 ns/op
BenchmarkFor-2          200000000                5.21 ns/op
BenchmarkFor-4          1000000000               2.13 ns/op
BenchmarkFor-8          2000000000               0.80 ns/op
BenchmarkFor-16         2000000000               0.66 ns/op
PASS
ok      github.com/aprice/gordinals/ordinal     9.887s
```

Here we've passed in the `-cpu` parameter, which lets us specify how many
threads we want to use for the parallel benchmarks. By passing in a
comma-separated list of values, we get to see the benchmarks executed for each,
so we can compare performance. The results are indicated for each, with
`BenchmarkFor` being the single-threaded test, and `BenchmarkFor-n` being the
results for *n* threads.
