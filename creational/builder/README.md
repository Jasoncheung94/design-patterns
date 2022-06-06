## Builder Pattern

The builder pattern is a very common creational design pattern. Its main purpose is to encapsulate and simplify the creation of complex objects, which makes it much easier to create objects that can have many different representations.


__Scenarios__: 
- Objects that have complex API's, multiple constructor options, and several different representation.

### Advantages & Disadvantages
- Encapsulates code for construction and representation.
- Provides control over steps of construction process.
- A concrete builder must be created for each type of product.
- Builder classes must be mutable.
- May complicated dependency injection.


## Examples

There is different implementations of the builder, I've linked the different versions I have discovered so far. 
The Pizza and House examples would be my preferred options to reference.
| Type of example | 
| --- |
| [Pizza](./pizzaExample/) |
| [House](./houseExample/) |
| [Notification](./notificationExample/) |
