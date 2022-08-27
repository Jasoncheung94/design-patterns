# Multiton

Multiton pattern is a design pattern which generalizes the singleton pattern. Whereas the singleton allows only one instance of a class to be created, the multiton pattern allows for the controlled creation of multiple instances, which it manages through the use of a map or some data structure.

Rather than having a single instance per application (e.g. the java.lang.Runtime object in the Java programming language) the multiton pattern instead ensures a single instance per key.

## Drawbacks

This pattern, like the Singleton pattern, makes unit testing far more difficult,[1] as it introduces global state into an application.

With garbage collected languages it may become a source of memory leaks as it introduces global strong references to the objects.

### Extra Resources

- [Multiton Wikipedia](https://en.wikipedia.org/wiki/Multiton_pattern)

### Examples

| Language   | Description    | Status | Code Examples                              |
| ---------- | -------------- | ------ | ------------------------------------------ |
| Golang(Go) | Simple Example | ✅     | [main.go](./simple-example/golang/main.go) |
