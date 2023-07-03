# jorobin

Package jorobin provides a Go implementation of round-robin selection, a technique for selecting elements in a circular order. It is commonly used to distribute workloads across multiple servers or to select items from a list in a rotating order.

## Installation

To use `jorobin`, you can install it using `go get`:

```
go get github.com/jonathantyar/jorobin
```

## Usage

Here is an example usage of the `jorobin` package:

```go
import (
    "fmt"
    "github.com/jonathantyar/jorobin"
)

func main() {
    rr, err := jorobin.New("foo", "bar", "baz")
    if err != nil {
        log.Fatal(err)
    }

    for i := 0; i < 5; i++ {
        fmt.Println(rr.Next())
    }
}
```

In this example, we create a new `RoundRobin` object with a list of three strings ("foo", "bar", and "baz"). We then call the `Next()` method five times to select items from the list in a round-robin fashion.

The `RoundRobin` object maintains a pointer to the current item in the list, and updates the pointer each time `Next()` is called. When the end of the list is reached, the pointer is reset to the beginning of the list.

## API

The `RoundRobin` interface represents the round-robin balancing algorithm, and the `roundrobin` struct provides an implementation of this interface.

The `RoundRobin` interface has the following methods:

- `Next() string`: Returns the next item in the list in a round-robin fashion.

- `Total() int`: Returns the number of items in the list.

- `New(newItem string)`: Adds a new item to the list.

- `Remove(item string)`: Removes the specified item from the list.

- `Clear()`: Clears all items from the list.

The `New()` function creates a new `RoundRobin` object with the specified list of items.

Here's an example usage of these additional methods:

```go
import (
    "fmt"
    "github.com/jonathantyar/jorobin"
)

func main() {
    rr, err := jorobin.New("foo", "bar", "baz")
    if err != nil {
        log.Fatal(err)
    }

    rr.New("qux")
    fmt.Println(rr.Total()) // Output: 4

    rr.Remove("bar")
    fmt.Println(rr.Next()) // Output: foo

    rr.Clear()
    fmt.Println(rr.Next()) // Output: ""
}
```

In this example, we create a new `RoundRobin` object with a list of three strings ("foo", "bar", and "baz"). We then add a new string "qux" to the list using the `New()` method, and print out the total number of items in the list using the `Total()` method.

We then remove the string "bar" from the list using the `Remove()` method, and call the `Next()` method to retrieve the next item in the list, which should be "foo".

Finally, we clear all items from the list using the `Clear()` method, and call `Next()` again to confirm that the list is now empty.
