# TRANSITIONING FROM JAVA TO GOLANG

## SETUP GIT

Download and install GIT.

Set your global username:

```powershell
git config --global user.name "Your Name"
```

Set your global email address:

```powershell
git config --global user.email "your.email@example.com"
```


Confirm/verify your user.name and email configuration have been set correctly (optional), you can use the following commands: 

```powershell
git config --global user.name
git config --global user.email
```

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

## DECLARING VARIABLE

In Java we declare variables in the following syntax (for examole, type string):
 
```java
String greeting = "";
```

The same declaration is Goland will use the following syntax:

```golang
var greeting string
```

Variables can be declared and assigned a value at the same time:

```java
String greeting = "Hello, world!";
```

```golang
var greeting = "Hello, world!"
```

The preferred way to initialize a variable in Golang is by using the shorthand declaration (with the colon and equal sign ":="):

``golang
greeting := "Hello, world!"
```

By default names of variables functions, constants, and types are package scope. It's difficult comparing it to Java access modifiers, since Golang has no class inheritance. The closest will be the `package private` or `default` default modifier. 

We can make a name available outside of package scope, making it public to others importing the package, by simply capitalizing the name of the variable. This is known as exporting the name.  accessable A name can be exported by using a capital letter, making it a available outside the package scope. 

**NOTE:** When declaring a variable in a program we must use it or get an error about a varialbe that is declared and not used. 

We will change the "hello.go" program in another program "hello-var.go" and it will have the following code:

```golang
package main

import "fmt"

func main() {
	greeting := "Hello, LinkedIn!"
   fmt.Println(greeting)
}
```

We will compile and run the program using the command `go run hello-var.go`.

If we will not use the variable `greeting` as the argument to the `fmt.Println` function, then during compilation we will get an error about a variable `greeting` that is declared and not used.
