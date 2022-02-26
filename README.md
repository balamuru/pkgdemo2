# pkgdemo
Demonstration of local packaging in Golang

## Layout
```
.
├── driver              package containing invoking program
│   ├── driver.go       invoking program that referemces the greetings module
│   └── go.mod          module file containing module reference and "replace" directive that points to local version
├── greetings           package containing greeting module
    ├── go.mod          module file that describes the greeting module
    └── greetings.go    greeting module
```
## create `pkgdemo` module
```
mkdir pkgdemo
cd pkgdemo
```

## Create `greetings` Go module
```
mkdir greetings
cd greetings
go mod init 
// create and code up greetings.go - see source
```

## Invoke the `driver` module locally (does not require publication to github)
```
cd ..
mkdir driver
cd driver
go mod init
//create driver program - see source
//go mod edit -replace github.com/balamuru/pkgdemo/greetings=../greetings //replace github module as needed
go mod tidy
```

## Run the `driver` program (thereby invoking the `greetings` module)
```
go run driver.go
```

## References
* https://go.dev/doc/tutorial/create-module
* https://go.dev/doc/tutorial/call-module-code