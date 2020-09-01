## Week 2 Quiz

1. Is the highlighted assignment to f in the following code a legal variable assignment?

![](https://user-images.githubusercontent.com/49638680/91836278-d8f6a800-ec4a-11ea-9994-a09cc1b592eb.png)

The variable `f` is assigned to a function, taking a string as argument and returning an integer.

2. Which of the following statements correctly declares a function whose argument is another function which takes an integer as an argument and returns a string?

![](https://user-images.githubusercontent.com/49638680/91836284-dac06b80-ec4a-11ea-87d4-b1a273ff2a3c.png)

As one can see, `fA` takes as argument a function `fB`, whose argument is an integer and returns a string.

3. What is an anonymous function?

![](https://user-images.githubusercontent.com/49638680/91836298-de53f280-ec4a-11ea-9141-7262e02a7b3b.png)

Simply a function with no explicit name.

4. Which of the following statements correctly declares a function whose return value is another function which takes a string as an argument and returns an integer?

![](https://user-images.githubusercontent.com/49638680/91837283-4ce58000-ec4c-11ea-8e5e-ad241149a749.png)

The function `fA` takes as argument the function `fB` (whose properties we can neglect for this question) and returns another function whose argument is a string and returns an integer.

5. What does the above code print on the screen?

![](https://user-images.githubusercontent.com/49638680/91836316-e14ee300-ec4a-11ea-84b6-dcad28f09bfb.png)

This question is about function environments.
Indeed, when you assign `fB` to `fA` you first initialise `i` to zero, then you return `func()` to `fB`.
Hence, you call `fB` twice, meaning calling `func()`.
Thus first returning `i = 1`, then `i = 2`.

6. What symbols are used in a function declaration to indicate that it is a variadic function?

![](https://user-images.githubusercontent.com/49638680/91836323-e3b13d00-ec4a-11ea-931f-3a7a4bdc2764.png)

Variadic functions, with a variable number of arguments are indicated by the three dots, `...`.

7. What does this routine produce?

![](https://user-images.githubusercontent.com/49638680/91836328-e57b0080-ec4a-11ea-9860-d58b6b61e699.png)

_Note the typo, the routine so written produces an error. By correcting the typo (`I` -> `i`), you get the correct answer `123`._

At line `7`, you assign the value 1 to the variable `i` and print it, so the first digit is `1`.
Then you increase the value stored in `i` by one, and you evaluate `i+1` (=3) and pass it to the deferred call of `Print`.
At this stage the `fmt.Print(i)` is executed and so the second digit is `2`.
Finally the deferred instruction can be executed, printint the last digit, _i.e._ `3`.
