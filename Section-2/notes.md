What is GOPATH?

\$GOPATH is a enviorment variable that references a workspace, in it is your source code, go tools and other related things like compiled packages and executable binaries. When yow download a go project it is by defauly downloaded at the GOPATH.

The GOPATH is set by default at your user directory under go.
~/go

go env is a go tool which shows enviorment variables and configurations which Go tools use. Use the following command to see the GOPATH:
go env GOPATH

At the GOPATH the directories can be found:
bin pkg src

src directory is the source-code directory.

- Note: It is good practice to place your code within the src/github.com directory under a directory with your github account name as its name.

Compile with "go build"
Is best used for big applications or to deploy.

Run with "go run"
Compiles and runs your program at the same time, its a greate simplifier for small programs.

PACKAGES

What is a package?

In Go, source files are organized into system directories called packages, which enable code reusability across the Go applications. The naming convention for Go package is to use the name of the system directory where we are putting our Go source files. Within a single folder, the package name will be same for the all source files which belong to that directory.

All package files should be in the same directory and all files should belong to the same package.

example:

    ./main/main.go
    package main
    func main() {}

    ./main/bye.go
    package main
    func bye() {}

    ./main/hey.go
    package main
    func hey() {}

The Package Clause should be the first code in a file and can only appear once. The package clause in a file lets Go know to which package a file belongs to.

All of the files belong to the package main. Each file has a package clause as "package main" and is inside the same directory.

Each Go package has its own scope. For example, declared funcs are only visible to the files belong to the main package.

There are two kinds of packages in Go: Executable and Library.

Excutable packages
package main + func main
Is created for running, is non-importable, it's name should be main and have a func main.

Library packages
package anyName + no func main
Is created for reusability, is non-executable, importable non-importable, can have any name and has no func main.

SCOPES

Scope relates to visibility, who can see what. There are package, file, func and block scopes.

What is the importance of names?

A declaration declares a unique name within a scope. Scope is defined by where a name is declared. The same name cannot be declared again inside the same scope.

Every line of code can have different scope depending on their position in a Go file.

example:

package main

import "fmt" <--- File scoped, only visible in this file.

const ok = true <--- Package Scoped,
func main() <--- visible to all the files that belong to the packages.
{
var hello = "Hello!" <--- block scoped
fmt.Println(hello, ok) <--- visible after its declaration within the function block
}

Package SCOPES
Names are visible throughout the package.

What is a package scope?
Each Go package has its own scope. For example, declared funcs are only visible to the files belong
to the same package.

Declarations which are outside of functions are visible to the files belong to the same package.

The same names in the same packages

File Scope
Names are visible throughout the file.
Each Go file has its own scope. Imported packages are only visible to the importing file.

IMPORTING
Importing allows a file to use functionalities from a library package. It as if you've declared what's inside the imported package's files in your own file.

Each file has to import external packages on its own.

Renaming imported packages
Because names must be unique in the same scope, it may happen that by including multiple import declarations names could collide.

You can import packages with the same name into the same file by giving one of them another name.

example:
package main
import "fmt"
import f "fmt" <--- now you can acees the package using f as its name.

STATEMENTS

What is a statement?
Statements instructs Go to execute something.

They can change the execution flow of code, like an if statement. They can also declare a name for a variable or function. Almost all the code is made of statements. Statements control the execution flow.

Go adds a semicolon between statements, behind the scenes.

example:
package main
import "fmt"

    func main() {
        fmt.Println("Hello")
        fmt.Println("Cool"); fmt.Println("Super Cool")
        if 5 > 1 {
            fmt.Println("bigger")
        }
    }

EXPRESSIONS

what is an expressions?
Expressions computes one or multiple values. They calculate or evaluate and return one or more values.

example:
There's only one expression here: "Hello!" string literal.

    package main
    import "fmt"

    func main() {
        fmt.Println("Hello")
    }

Operators allow you to combine expressions.

example:
package main
import "fmt"

    func main() {
        fmt.Println("Hello" + "!")
    }

Functions can be used as expressions inside statements.

example:
package main
import (
"fmt"
"runtime"
)

    func main() {
        fmt.Println(runtime.NumCPU())
    }

Should be used with or within a statement.

Some statements like func calls can also act like expressions.

COMMENTS

How Go comments work?

    // single line comments

    /*
     multi line comments
     more than single line
    */

What is Go Doc? (https://blog.golang.org/godoc)
Godoc parses Go source code - including comments - and produces documentation as HTML or plain text. The end result is documentation tightly coupled with the code it documents.

Write a Library Package

Create your first library package

A library package doesn't have to be compiled, in can be directly imported.

EXPORTING
Exporting allows a package to make its functionalities available to other packages.

To export a name just make its first letter an uppercase letter.

How Go standard library exports?

Exporta a function from your package
