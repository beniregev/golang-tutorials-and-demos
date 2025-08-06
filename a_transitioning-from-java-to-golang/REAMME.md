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

### BASIC DATA TYPES

Golang data types are divided into "Basic" and "Composite" data-types. 

The "Basic" data-types include numbers, strings, and Boolean types. They are the basic building blocks of our data. 

The "Composite" data types consist of:

- arrays.
- structs.
- pointers.
- functions.
- interfaces.
- slices.
- maps.
- channels.

There 8 composite data-types use the basic types to build more advanced constructs and structures. 

In Golang both signed and unsigned integers are available in the 8-bit, 16-bit, 32-bit, and 64-bit formats, as shown in the following table:

#### SIGNED INTEGERS

| Type | Size | Range |
|:--- | :--- | :---: |
| int8 | 8-bit | -128 to 127 |
| int16 | 16-bit | -2<sup>15</sup> to 2<sup>15</sup> |
| int32 | 32-bit | -2<sup>31</sup> to 2<sup>31</sup> |
| int64 | 64-bit | -2<sup>63</sup> to 2<sup>63</sup> |
| int | System<br>dependent | Systen<br>dependent |

#### UNSIGNED INTEGERS

| Type | Size | Range |
|:--- | :--- | :---: |
| uint8 | 8-bit | 0 to 255 |
| uint16 | 16-bit | 0 to 2<sup>16</sup> - 1 |
| uint32 | 32-bit | 0 to 2<sup>32</sup> - 1 |
| uint64 | 64-bit | 0 to 2<sup>64</sup> - 1 |
| uint | System<br>dependent | Systen<br>dependent |

If we don't specify a format (8, 16, 32, or 64) then the integer will be either 32 or 64 bit depending on the system we are compiling for.

#### OTHER ALIAS DATA TYPES

The type `byte` is an alias for the data-type `uint8` and the type `rune` is an alias for the data-type `int32`. These 2 type represent character values. Unlike Java, Golang doesn't have a `char` data-type, so we use these 2 aliases instead.

#### FLOATING POINT DATA TYPES

Aside from integers Goland has 2 floating point types: `float32` and `float64`. By default the `float64` is inferred by the compiler. For example, when assigning a value in the following code, thje compiler with implicitly assign the `float64` type to it.

```golang
f := 123.456
```

#### BOOL DATA TYPE

The type `bool` represent a Boolean value in Golang. Boolean store `true` or `false` values, and support the AND (&&), OR (||), and teh NOT (!) operators. Also, as expected logical operators support short circuit logic. The && (AND), and || (OR) support short circuit logic, for example:

- In the expression `v1 && v2`, if `v1` is `false`, then `v2` is not evaluated (the result `false` is already known, and there's no chance for `true` in the result).

#### STRING DATA TYPE

In Golang we declare string either with double quotes (") or back-ticks (`). Strings enclosed in back ticks are raw strings where escape characters don't have any special meaning. In the following code, both statements are the same with the "Hello, world!" and "Goodbye" each on their own lines:

```golang
// Double quoted string - mew line character escaped
var greeting := "Hello, world!\nGoodbye, world!"

// Raw string (back-ticked) - can span multiple lines
var raw := `Hellow, world!
            Goodbye, world!`
```

#### ZERO VALUES

Unless explicitly initialized, variables are automatically initialized with their "zero value":

| Type | Zero Value |
| :--- | :---: |
| Integer | 0 |
| Floating Point | 0.0 |
| Boolean | false |
| String | ""<br>(Empty string, not `null`) |

This means that "Basic" data types are never `nil` (`null`) and need not to be `nil` checked.

The file `1_zero-values/zero-values.go` contains code to demonstrates zero-values.

### COMPOSITE DATA TYPES

#### POINTERS

A pointer is a composite data-type taht stores the memory address of a variable. It provides a way to point to where the memory is located and find the value saved at that address.

There are 2 pointer operators:

- The * operator is known as the dereferencing operator. It's used to declare a pointer variable, as well as access the value stored at that address.
- The & operator is known as the address operator. It's user to return the memory address of a variable, which can then be saved to a pointer variable.

Pointers are associated with a data type, and only references to that type can be stored in them. 

The syntax to declare apointer is `var name *Type`. For example, `var ptr *string` which declares a pointer to a string. 

Given the variable `a`, we initialize `var ptr *string = &a`. We initialize the pointer to the variable `a` using the address operator. The pointer will contain the memory address of the variable.

Given an initialized pointer, we retrieve the value saved at that given pointer: `value := *ptr`.

We can think of the 2 operator as a 2 oposites. 

The zero value of a pointer is the `nil` value, that's how we call `null` value in Golang. As expected from our Java experience, operations performed on `nil` values will cause a oanic or `nil` pointer error at runtime.

The following is the code of "pointers.go" file:

```golang
package main

import "fmt"

func main() {
      // Declare pointer
      var ptr *string

      // initialize a greeting
      greeting := "Hello, world!"

      // Assign greeting address to pointer
      ptr = &greeting

      // Print out our variables
      fmt.Println("Greeting: ", greeting)
      fmt.Println("Address of greeting: ", ptr)
      fmt.Println("Value stored in ptr: ", *ptr)
}
```

Let's analize the code:

- First, we declare a pointer variable which receives a `nil` value.
- Second, we declare a sting variable with a string, this variable is assigned a memory address value.
- At this point the `ptr` pointer and the `greeting` variables are still unrelated and unlinked.
- Finally, we assign the address of teh `greeting` variable to the `ptr` pointer.

In Golang, pointers can be a bit strange to deal with explicitly. We will revisit them throughout this tutorial. So don't worry if you still find them confusing after this introduction.
