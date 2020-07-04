## What is Go path?

The Go path is a list of directory trees containing Go source code. It is consulted to resolve imports that cannot be found in the standard Go tree. The default path is the value of the GOPATH environment variable, interpreted as a path list appropriate to the operating system (on Unix, the variable is a colon-separated string).

\$GOPATH is a enviorment variable that references the list of directory trees. It can be accessed with **go env**, a go tool which shows enviorment variables and configurations. Use the following command to see the GOPATH:

```
    go env GOPATH
```

Each directory listed in the Go path must have a prescribed structure:

```
    src pkg bin
```

The src/ directory holds source code. The path below 'src' determines the import path or executable name.

The pkg/ directory holds installed package objects.

If DIR is a directory listed in the Go path, a package with source in DIR/src/foo/bar can be imported as "foo/bar" and has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a" (or, for gccgo, "DIR/pkg/gccgo/foo/libbar.a").

The bin/ directory holds compiled commands. Each command is named for its source directory, but only using the final element, not the entire path. That is, the command with source in DIR/src/foo/quux is installed into DIR/bin/quux, not DIR/bin/foo/quux. The foo/ is stripped so that you can add DIR/bin to your PATH to get at the installed commands.

_Note: It is good practice to place your code within the src/github.com directory under a directory with your github account name as its name._

## Command go

Go is a tool for managing Go source code.

Usage:

```
    go <command> [arguments]
```

The commands are:

- **bug** start a bug report
- **build** compile packages and dependencies
- **clean** remove object files and cached files
- **doc** show documentation for package or symbol
- **env** print Go environment information
- **fix** update packages to use new APIs
- **fmt** gofmt (reformat) package sources
- **generate** generate Go files by processing source
- **get** add dependencies to current module and install them
- **install** compile and install packages and dependencies
- **list** list packages or modules
- **mod** module maintenance
- **run** compile and run Go program
- **test** test packages
- **tool** run specified go tool
- **version** print Go version
- **vet** report likely mistakes in packages

Use "go help <command>" for more information about a command.

## Packages

### What is a package?

In Go, source files are organized into system directories called packages, which enable code reusability across the Go applications. The naming convention for Go package is to use the name of the system directory where we are putting our Go source files. Within a single folder, the package name will be same for the all source files which belong to that directory.

All package files should be in the same directory and all files should belong to the same package.

example:

```
    ./main/main.go
    package main
    func main() {}

    ./main/bye.go
    package main
    func bye() {}

    ./main/hey.go
    package main
    func hey() {}
```

The Package Clause should be the first code in a file and can only appear once. The package clause in a file lets Go know to which package a file belongs to.

Each Go package has its own scope. For example, names declared outside of blocks ({}) are visible to the files that belong to the same package.

There are two kinds of packages in Go: Executable and Library.

### Excutable packages

Has a package main and a func main.

Is created for running, is non-importable, it's packge name must be main and have a func main.

### Library packages

Can have any package name diferent to main and no func main.

Is created for reusability, is non-executable, importable, and can have any name other than main and has no func main.

## SCOPES

Scope relates to visibility, who can see what. There are package, file, func and block scopes.

What is the importance of names?

A declaration declares a unique name within a scope. Scope is defined by where a name is declared. The same name cannot be declared again inside the same scope.

Every line of code can have different scope depending on their position in a Go file.

example:

```
    package main
    import "fmt" <--- File scoped, only visible in this file.

    const ok = true <--- pPackage scoped

    func main() <--- Visible to all the files that belong to the packages.
        {

        var hello = "Hello!" <--- Block scoped

        fmt.Println(hello, ok) <--- Visible after its declaration within the function block
        }
```

### Package Scope

Names declared outside blocks are visible throughout the package.

### What is a package scope?

Each Go package has its own scope. For example, declared funcs are only visible to the files belong to the same package.

Declarations which are outside of functions are visible to the files belong to the same package.

### File Scope

Names are visible throughout the file. Each Go file has its own scope. Imported packages are only visible to the importing file.

## Importing

Importing allows a file to use functionalities from a library package. Its as if you've declared what's inside the imported package's files in your own file.

Each file has to import external packages on its own.

### Renaming imported packages

Because names must be unique in the same scope, it may happen that by including multiple import declarations names could collide.

You can import packages with the same name into the same file by giving one of them an alias.

example:

```
    package main
    import "fmt"
    import (
        f "fmt" <--- now you can access the package using f as an alias.
        "runtime"
        "math
    )
```

## Statements

### What is a statement?

Statements instructs Go to execute something.

They can change the execution flow of code, like an if statement. They can also declare a name for a variable or function. Almost all the code is made of statements. Statements control the execution flow.

Go adds a semicolon between statements, behind the scenes.

example:

```
    package main
    import "fmt"

    func main() {
        fmt.Println("Hello")
        fmt.Println("Cool"); fmt.Println("Super Cool")
        if 5 > 1 {
            fmt.Println("bigger")
        }
    }
```

## Expressions

### What is an expressions?

Expressions computes one or multiple values. They calculate or evaluate and return one or more values.

example:

```
    // There's only one expression here: "Hello!" string literal.

    package main
    import "fmt"

    func main() {
        fmt.Println("Hello")
    }
```

Operators allow you to combine expressions.

example:

```
package main
import "fmt"

    func main() {
        fmt.Println("Hello" + "!")
    }
```

Some statements like func calls can also act like expressions, and can be used as expressions inside statements.

example:

```
    package main
    import (
        "fmt"
        "runtime"
    )

    func main() {
        fmt.Println(runtime.NumCPU())
    }
```

Expressions must be used with or within a statement. They can never be alone.

## Comments

How Go comments work?

    // single line comments

    /*
     multi line comments
     more than single line
    */

## What is Go Doc?

https://blog.golang.org/godoc

Godoc parses Go source code - including comments - and produces documentation as HTML or plain text. The end result is documentation tightly coupled with the code it documents.

## Exporting

Exporting allows a package to make its functionalities available to other packages.

To export a name just make its first letter an uppercase letter.
