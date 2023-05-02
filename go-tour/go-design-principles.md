## Go Design Principles
    https://youtu.be/zzAdEt3xZ1M

## SOLID
   1. Single Responsibility Principle
   2. Open / Closed Principle
   3. Liskov Substitution Principle
   4. Interface Segregation Principle
   5. Dependency Inversion Principle

### Single Responsibility Principle
   A class should have one and only one reason to change. Go has a more powerful notion of composition


### Go vs Java - Side by side comparisons
    https://youtu.be/1ZjvhGfpwJ8
- Object Oriented
    - Encapsulation  >> Yes
    - Inheritance    >> No   >> Composition
    - Polymorphism   >> No

- classes >> struct

- Constructors
    - variables initialized to default values >> empty values

- Enums  >> could use const( ) and iota

- interfaces

- generics >> newly introduced

- exceptions >> errors, panics

References
    https://medium.com/@genchilu/javas-thread-model-and-golang-goroutine-f1325ca2df0c#:~:text=How%20Goroutine,another%20goroutine%20in%20about%2010ms.
    https://medium.com/@damithadayananda/golang-vs-java-concurrency-a-comparative-study-b0aea90f5fd7
    http://txt.fliglio.com/2014/04/concurrency-in-java-and-go/
