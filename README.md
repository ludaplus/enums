# Enum Utility for Go

Use this utility to define and use enums in Go, just like other languages.

## Features

1. **Real Namespaces.** You can define `Http.Status`, `HttpStatus.OK`, `Result.OK` in a same package.
2. **Simplest Implement.** Just use Generics and a little Reflection.
3. **One-time Reflection.** Very fast and efficient.
4. **No Dependencies.** No third-party dependencies.
5. **No Interface.** No need to implement an interface.
6. **No Code Generator is needed.** Write, and use enums in Go directly.
7. **No Bullshit Code.** Write the name of element just one time, and use it.
8. **Extensible.** You can add more methods to the elements of the enum.

## Installation

```bash
go get github.com/ludaplus/enums
```

## Usage

### Define an Enum

Define an inner Type for the element of the enum. The inner type should embed the `enums.Element` type.

```go
type innerPostType struct {
	enums.Element
	CommentEnabled bool
}
```

Create a variable of the enum type and initialize it with the `enums.Of` function.

```go
var PostType = enums.Of(&struct {
	enums.Enum[*innerPostType]
	Unknown,
	Post,
	Page,
	Note *innerPostType
}{
	enums.Enum[*innerPostType]{},
	&innerPostType{
		CommentEnabled: false,
	},
	&innerPostType{
		CommentEnabled: true,
	},
	&innerPostType{
		CommentEnabled: false,
	},
	&innerPostType{
		CommentEnabled: false,
	},
})
```

Notes:
1. The argument of the `enums.Of` function is a pointer to a struct that contains the elements of the enum.

Suggestions:
1. Put the `Unknown` as second field (first element), then its ordinal will be ZERO.
2. Use value-base initialization for the elements, in order to avoid missing. You will get an error if you miss the initialization of an element.
   In Goland and other IDEs, you will also see a label with field name of the value defined in the struct.

### Access the Enum

Use the enum variable to access the elements.

```go
func main() {
    fmt.Println(PostType.Unknown.Name())
    
    fmt.Println(PostType.Post.CommentEnabled)
	
    fmt.Println(PostType.ValueOf("Page").Ordinal())
	
    fmt.Println(PostType.ValueOf("Note") == PostType.Note)
	
    for _, postType := range PostType.Values() {
        fmt.Println(postType.CommentEnabled())
    }
}
```

### Enum Element Methods

The `enums.Element` type has the following methods:

1. `Name() string`: Returns the name of the element.
2. `Ordinal() int`: Returns the ordinal of the element.

### Enum Methods

The enum type has the following methods:

1. `ValueOf(name string) *T`: Returns the element with the given name.
2. `Values() []*T`: Returns all elements of the enum.
3. `Names() []string`: Returns the unknown element of the enum.

### Extra

You can use the `enums.Of` function to define an enum with a different way.

```go
var PostType = enums.Of(&struct {
	enums.Enum[*innerPostType]
	Unknown *innerPostType
	Post,
	Page,
	Note *innerPostType
}{
	Unknown: &innerPostType{
		CommentEnabled: false,
	},
	Post: &innerPostType{
		CommentEnabled: true,
	},
	Page: &innerPostType{
		CommentEnabled: false,
	},
	Note: &innerPostType{
		CommentEnabled: false,
	},
})
```