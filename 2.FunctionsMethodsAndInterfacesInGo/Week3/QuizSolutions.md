## Week 3 Quiz

1. What is the difference between an __object__ and a __class__?

![](https://user-images.githubusercontent.com/49638680/92139413-fd4fb180-ee0f-11ea-90f0-86ce7bba5f87.png)

This is subtle, but easy if you listened to the first video of the week.
A class has been described as a template of a structure encoding data and methods to act on these data. An object is nothing more than an instance of that template.

2. What is the difference between a struct in Go and a class in an object-oriented language?

![](https://user-images.githubusercontent.com/49638680/92139419-ffb20b80-ee0f-11ea-8014-8e304b43dd78.png)

In other OO-languages (like Java, C++, Python, etc.) data and methods come all together in the class template. While in Go, you encode only data in structures and then you build methods with a receiving type.

3. Which of the following refers to data hiding?

![](https://user-images.githubusercontent.com/49638680/92139428-02acfc00-ee10-11ea-80cf-8614d549b12d.png)

The _encapsulation_ is sometimes referred as data-hiding.

4. How do you associate a method with an arbitrary data type on Go?

![](https://user-images.githubusercontent.com/49638680/92139434-0476bf80-ee10-11ea-9bf3-58bbe04d150d.png)

To associate a method to a particular data structure, or a type, you define the receiver type for the method.

_e.g._
```Go
func (p MyNewType) Method (argument string) (result float64) {
  ...
}
```

5. In Go, how do you hide variables or functions in a package, so that functions outside of the package cannot access them?

![](https://user-images.githubusercontent.com/49638680/92139448-080a4680-ee10-11ea-8571-cb2d556ddc67.png)

In other OO-languages there are keywords (C++) or you have to start your methods with a special character (Python). In Go, private methods are the ones whose name starts with lower-case letter.

6. Say that you have defined a type `t` and you have declared an object of that type called `t1`. Assume that the type `t` is the receiver type for a method called `Foo()`. Which expression shows a proper invocation of the the method `Foo()`?

![](https://user-images.githubusercontent.com/49638680/92139456-0a6ca080-ee10-11ea-83a8-eaf8a08d4cca.png)

Methods for types are always called as the `instance_name.Method()`.

7. Assume that that the type `t` is the receiver type for a method called `Foo()`. Under what conditions would it be better to make the receiver type of `Foo()` a pointer to `t`, rather than itself?

![](https://user-images.githubusercontent.com/49638680/92139467-0d679100-ee10-11ea-8bf3-cb195249223b.png)

Methods have an implicit argument, which is the structure instance. This is passed _by value_, meaning the method gets a copy of the argument.
First, this may cause a memory overflow if the object is big.
Second, if it tries to modify the argument, it will only modify its copy, having no effect on the structure instance.
Thus, passing a pointer to the object, _i.e._ passing the argument _by reference_ would be the solution to these two issues.
