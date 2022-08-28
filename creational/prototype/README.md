## Prototype

Lets you copy existing objects without making your code dependent on their classes.

It is used when the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects. This pattern is used to avoid subclasses of an object creator in the client application, like the factory method pattern does and to avoid the inherent cost of creating a new object in the standard way (e.g., using the 'new' keyword) when it is prohibitively expensive for a given application.

### Extra Resources

- [Prototype Wikipedia](https://en.wikipedia.org/wiki/Prototype_pattern)

### Examples

| Language   | Description    | Status | Code Examples                              |
| ---------- | -------------- | ------ | ------------------------------------------ |
| Golang(Go) | Simple Example | ✅     | [main.go](./simple-example/golang/main.go) |
| Python     | Simple Example | ✅     | [main.py](./simple-example/python/main.py) |
