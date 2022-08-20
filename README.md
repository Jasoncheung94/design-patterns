# design-patterns

Curated list of design patterns and idioms using Go (mostly) and other languages such as Python.

[Design patterns](https://en.wikipedia.org/wiki/Software_design_pattern)

## Creational Patterns

| Pattern                                       | Description                                                                                                                                           | Status |
| --------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------ |
| Abstract Factory                              | Provides an interface for creating families of related objects without specifying their concrete classes                                              | ❌     |
| [Builder](./creational/builder/README.md)     | Builds a complex object using simple objects. Allows you to produce different types and representations of an object using the same construction code | ✅     |
| Chain of Responsibility                       | Allows you to pass requests along a chain of objects until one of them can handle the request.                                                        | ❌     |
| Factory Method                                | Provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created                   | ❌     |
| Prototype                                     | allows cloning objects, even complex ones, without coupling to their specific classes                                                                 | ❌     |
| [Singleton](./creational/singleton/README.md) | Only one object of its kind exists and provides a single point of access to it for any other code ( aka Global variable )                             | ✅     |

## Structural Patterns

| Pattern   | Description                                                                                  | Status |
| --------- | -------------------------------------------------------------------------------------------- | ------ |
| Adapter   | Allows objects with incompatible interfaces to collaborate.                                  | ❌     |
| Bridge    | Separates an abstraction from its implementation so that the two can vary independently.     | ❌     |
| Composite | Allows you to compose objects into tree structures to represent part-whole hierarchies.      | ❌     |
| Decorator | Allows you to attach additional responsibilities to an object dynamically.                   | ❌     |
| Facade    | Provides a unified interface to a set of interfaces in a subsystem.                          | ❌     |
| Flyweight | An object that minimizes memory usage by sharing some of its data with other similar objects | ❌     |
| Proxy     | Provides a substitute or placeholder for another object to control access to it.             | ❌     |

## Behavioral Patterns

| Pattern                 | Description                                                                                                               | Status |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------ |
| Chain of Responsibility | Allows you to pass requests along a chain of objects until one of them can handle the request.                            | ❌     |
| Command                 | Allows you to encapsulate a request as an object, thereby letting you parameterize other objects with different requests. | ❌     |
| Interpreter             | Allows you to implement an interpreter pattern.                                                                           | ❌     |
| Iterator                | Allows you to traverse a set of objects without keeping them in memory.                                                   | ❌     |
| Mediator                | Allows you to decouple components that communicate with each other.                                                       | ❌     |
| Memento                 | Allows you to save and restore the state of an object without breaking its encapsulation.                                 | ❌     |
| Observer                | Allows you to decouple the subscriber from the publisher (sender) by implementing the publish/subscribe pattern.          | ❌     |
| State                   | Allows you to encapsulate the current state of an object in an object.                                                    | ❌     |
| Strategy                | Allows you to define a family of algorithms, put each one in a separate class, and make their objects interchangeable.    | ❌     |
| Template Method         | Allows you to define the skeleton of an algorithm in an operation, deferring some steps to subclasses.                    | ❌     |
| Visitor                 | Allows you to separate an algorithm from the data structures that support it.                                             | ❌     |


## Concurrency Patterns (TODO) - Find common idioms and patterns with different languages

| Pattern                   | Description | Status |
| ------------------------- | ----------- | ------ |
| Active Object             |             | ❌     |
| Balking pattern           |             | ❌     |
| Barrier                   |             | ❌     |
| Double-checked locking    |             | ❌     |
| Guarded suspension        |             | ❌     |
| Leaders/followers pattern |             | ❌     |
| Monitor Object            |             | ❌     |
| Nuclear reaction          |             | ❌     |
| Reactor pattern           |             | ❌     |
| Read write lock pattern   |             | ❌     |
| Scheduler pattern         |             | ❌     |
| Thread pool pattern       |             | ❌     |
| Thread-local storage      |             | ❌     |
