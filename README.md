# pkgdemo
Demonstration o packaging in Golang backed by GitHub

## Layout
```
.
├── driver              package containing invoking program
│   ├── driver.go       invoking program that referemces the greetings module
│   └── go.mod          module file containing reference  to the greetings
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

## Push changes to Github

## Invoke the `driver` module locally (does not require publication to github)
```
cd ..
mkdir driver
cd driver
go mod init
//create driver program - see source
go mod tidy
```

## Run the `driver` program (thereby invoking the `greetings` module)
```
go run driver.go
```

## References
* https://go.dev/doc/tutorial/create-module
* https://go.dev/doc/tutorial/call-module-code