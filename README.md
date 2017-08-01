A simple example that shows how to use build tags to avoid including test dependencies in production code. 

Here, we used a mock clock in a file with `// +build test` and the built in methods in `// +build !test`.

When compiled with `go build`/`go install`, only the default stdlib code is included.
When running tests, use `go test -tags test` to compile with the test flag. 

# Example:

```
go build
./timebuild

2017-08-01 11:47:34.711856687 -0400 EDT

----

go test -tags test

2015-03-07 11:12:39 +0000 PST
PASS
ok  	github.com/groob/timebuild	0.015s

```


