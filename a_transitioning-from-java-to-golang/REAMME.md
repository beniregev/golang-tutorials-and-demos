# TRANSITIONING FROM JAVA TO GOLANG

## Installation

In your browser navigate to [Golang Installation Documentation](https://go.dev/doc/install), follow and complete the instruction steps for your operating system (OS).

Open a command-line terminal window and run `go version` to check and verify your Golang installation. You should get a response/output similar to the following:

      C:\development\golang-projects>go version
      go version go1.24.5 windows/amd64

      C:\development\golang-projects>

Now we can start out first Golang program, "Hello World" (what else?).

## THE BASICS

### FIRST PROGRAM -- `Hello World`

Open your IDE, I'm using VS Code. Create a folder for your first program(s), inside the folder create a file "hello.go".

**NOTE:** Filenames that start with a dot (.) or underscore (_) are ignored during compilation, e.g., ".hello.go", "_hello.go", etc.

Copy the following code into your "hello.go" file:

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello world!")
}
```

### TOOLS USED IN DEVELOPMENT

- `go run` - compile code and creates an executable in the "/10" folder, and then run it. It's a shortcut command that is used to run small programs, and is not recommended for running big programs in production.
- go build - Takes all the code in a defined package and its dependencies and creates an executable ready for running. This `go build` tool is very interesting and you can read about it outside of this tutorial. It's using caching under the hood for optimization and to make sure your builds are as quick as possible, and also allow for cross-compilation.
- `go test` - our "one stop shop" for running tests, including coverage and benchmarking. We will explore the `go test` tool in later sections of this tutorial, as test are so vital for development.

The "GO" toolchain has other interesting commands, which you can research and read up on, but these are the main commands that we will be using in this tutorial. Let's try out these commands in practice.

### RUNNING THE PROGRAM

Navigate to the folder where your program is located, in my example that's `/a_transitioning-from-java-to-golang/0_hello-world.go`.

Run command `ls` to verify you are in the directory of the `hello.go` file.

compile the program using the tool "go build" and the command:
```powershell
go build hello.go
```

Run command `ls -lh` and see that you have a new filw called "hello" that is around 2 MB in size - that's the compiled Golang program ready to run.

You can run the program in 2 ways
- Run the commnad `./hello` --> this will run the compiled program, IT WILL NOT COMPILE THE PROGRAM AGAIN.
- Run the command `go run hello.go` --> this will compiled the program and run it.

ß