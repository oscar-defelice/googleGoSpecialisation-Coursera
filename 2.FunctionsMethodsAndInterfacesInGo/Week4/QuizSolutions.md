## Week 4 Quiz

1. What is polymorphism?

![](https://user-images.githubusercontent.com/49638680/92223581-f58c1d80-eea0-11ea-8d29-1c0dff898d35.png)

_Polymorphism_, in this context, refers to the fact that different types can have same methods _signatures_ (name, arguments, return values) with different concrete implementations.

2. Which of the following statements is true?

![](https://user-images.githubusercontent.com/49638680/92223585-f755e100-eea0-11ea-8bb5-ea7b5c448908.png)

Inheritance and overriding enable polymorphism. Indeed with these two features a language allows polymorphic expressions, because a class can inherit a method (and its signature) and override it with a new implementation. This is often called _polymorphism for inclusion_.

3. If a type satisfies an interface, which of the following statements is true?

![](https://user-images.githubusercontent.com/49638680/92223589-f8870e00-eea0-11ea-982d-0c8666b7edf5.png)

Satisfying an interface for a type means to implement all the methods of the interface.

4. Which of the following statements is true?

![](https://user-images.githubusercontent.com/49638680/92223601-fc1a9500-eea0-11ea-987d-8096fac0f17c.png)

An interface always has a dynamic type, it can have no dynamic value, but the dynamic type is there.

5. Which of the following statements is/are true?

![](https://user-images.githubusercontent.com/49638680/92223605-fd4bc200-eea0-11ea-9edd-8d947810d516.png)

Interfaces model abstract similarities between types, _i.e._ they hide differences, so I is true.
In order to differentiate between the different types satisfying the interface, one can use _type assertions_, this makes II also true.
Finally, type assertions return two values, as many functions in Go, the first one is the underlying value of the asserted type (0 if the assertion is false), the second one is a boolean (true if the assertion is correct, false otherwise).

6. What is a use for an empty interface?

![](https://user-images.githubusercontent.com/49638680/92223615-0046b280-eea1-11ea-9fa3-7a549674f4e7.png)

An empty interface is one that is satisfied by all types. Thus, one can use it to define a function whose parameters can be of any type.

7. After executing the expression below, what is the value of err if there is no error?

![](https://user-images.githubusercontent.com/49638680/92223627-0341a300-eea1-11ea-8369-5b1d990d32ac.png)

The null value for errors in Go is `nil`.
