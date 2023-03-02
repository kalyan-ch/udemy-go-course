# udemy-go-course
Central repo for all projects from udemy go course - includes notes

## Packages, Imports and Functions

* packages are more like a workspace - collection of similar or common source code files 
* packages can be reusable or executable
* executable file should be of package `main`
* executable package should contain a function `main`
* any other package name means it is of a reusable type
* `import` keyword to include a package into your current code
* functions are blocks of callable, reusable code 
* functions in go start with keyword - `func`
* specify return type using `func func_name() return_type {}`
* files in the same package can use functions in other files without needing import

## Variables, Arrays and Loops

* variables in go are statically typed
* but go compiler can infer the type based on code
* full syntax for variable declaration `var card string = "AS"`
* alternative syntax - `card := "AS"` - := only used to declare variables
* can declare variables outside a function but can't assign a value to it
* array - list of elements of fixed length
* slice - list of elements that can grow and shrink, every element should be of sam type
* slices are immutable
* add elements to slice - `cards = append(cards, "6S")`
* syntax for for loop over a slice - i - index, card - element

```go
    for i, card := range cards {
        fmt.Println(i, card)
    }
```

## Misc

* Go is not an Object Oriented programming language
* custom type can have a receiver function - method that belongs to an "instance"
* any variable of type deck can access the receiver function defined in it
* call receiver function by calling - `variable.func_name()`
* receiver function syntax - t is the copy of the type variable avaiable

```go
func (t type) func_name() {

}
```
