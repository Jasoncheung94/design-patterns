## Momento

The memento pattern is a software design pattern that exposes the private internal state of an object. One example of how this can be used is to restore an object to its previous state (undo via rollback), another is versioning, another is custom serialization.

The memento pattern is implemented with three objects: the _originator_, a _caretaker_ and a _memento_. The originator is some object that has an internal state. The caretaker is going to do something to the originator, but wants to be able to undo the change. The caretaker first asks the originator for a memento object. Then it does whatever operation (or sequence of operations) it was going to do. To roll back to the state before the operations, it returns the memento object to the originator. The memento object itself is an opaque object (one which the caretaker cannot, or should not, change). When using this pattern, care should be taken if the originator may change other objects or resources—the memento pattern operates on a single object.

Memento is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.

### Additional Resources

- [Momento Wikipedia](https://en.wikipedia.org/wiki/Memento_pattern)

### Examples

| Language   | Description    | Status | Code Examples                              |
| ---------- | -------------- | ------ | ------------------------------------------ |
| Golang(Go) | Simple Example | ✅     | [main.go](./simple-example/golang/main.go) |
| Python     | Wiki Example   | ✅     | [main.py](./simple-example/python/main.py) |
