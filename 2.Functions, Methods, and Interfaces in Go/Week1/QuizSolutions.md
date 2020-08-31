## Week 1 Quiz

1. Using functions in code has which of the following impacts?

![](https://user-images.githubusercontent.com/49638680/91728935-16ecc100-eba4-11ea-9c08-32e31fae7359.png)

The use of functions makes your code more understandable, able to accomodate abstractions and reusable.

2. Which of these statements is true?

![](https://user-images.githubusercontent.com/49638680/91728967-24a24680-eba4-11ea-9a01-5d6c5783ee29.png)

You can specify different types for different arguments.

3. Let’s say that you are writing a function which takes a structure as an argument. What is a good reason to rewrite the function to take a pointer to the structure as an argument, instead of the structure itself?

![](https://user-images.githubusercontent.com/49638680/91728977-266c0a00-eba4-11ea-92ca-0127c005ac1b.png)

If you pass the structure itself, the function modifies only the copy of the structure, and when the function execution is over, the memory space will be deallocated. Hence, you end up with an unmodified structure.

4. Which of the features of functions listed below improve code understandability?

![](https://user-images.githubusercontent.com/49638680/91728992-2bc95480-eba4-11ea-8183-fe63fe81032d.png)

Understandability of the code improves when one follows the guide lines of having _low number of arguments_, _single task functions_ and _simple control-flow structures_.

5. What is a difference between passing a slice as an argument and passing an array as an argument?

![](https://user-images.githubusercontent.com/49638680/91729002-2ec44500-eba4-11ea-99d1-56ba72984f3a.png)

When you pass an array, the whole array is copied. It could be memory consuming.
