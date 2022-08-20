## Builder Pattern

The builder pattern is a design pattern designed to provide a flexible solution to various object creation problems in object-oriented programming. The intent of the Builder design pattern is to separate the construction of a complex object from its representation.

### Advantages

- Allows you to vary a product's internal representation.
- Encapsulates code for construction and representation.
- Provides control over steps of construction process.

### Disadvantages

- A distinct ConcreteBuilder must be created for each type of product.
- Builder classes must be mutable.
- May hamper/complicate dependency injection.

**Scenarios**:

- Objects that have complex API's, multiple constructor options, and several different representation.

### Examples

| Language   | Description    | Status | Code Examples                                    |
| ---------- | -------------- | ------ | ------------------------------------------------ |
| Golang(Go) | Simple Example | ✅     | [main.go](./simple-example/golang/main.go)       |
| Python     | Simple Example | ✅     | [builder.py](./simple-example/python/builder.py) |

### Additional Resources

- [Builder Wikipedia](https://en.wikipedia.org/wiki/Builder_pattern)
