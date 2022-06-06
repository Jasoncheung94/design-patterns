## Factory Pattern

Note: Go isn't an Object Oriented Programming language. Due to it's lack of features such as classes, inheritance a Simple Factory pattern will be added. This is a basic version of the pattern. When other languages are added I'll update this doc and add proper examples to give the overall idea. 


### Simple Factory

```go
type Dog Struct {
    Name string
    Weight float64
}

func NewDog(name string, weight float64) *Dog {
    return &Dog{
        Name: name,
        Weight: weight,
    }
}

```

__TODO__

### Interface Factory

### Prototype
