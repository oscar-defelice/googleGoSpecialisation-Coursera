## Week 2 Quiz

1. Which of the following expression does NOT compute the average of two integers a and b?

![](https://user-images.githubusercontent.com/49638680/78794006-765af200-79b3-11ea-9a09-bf9318cabc85.png)

The expression `2%(a+b)` returns the remainder of the _integer division_ between 2 and (a+b).

2. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/78794614-46f8b500-79b4-11ea-979a-ccbf8a7f0aa4.png)

The function `Atoi` in the package `strconv` converts the string value `"10"` to an integer. It is assigned to the variable `i`, then `y` is defined twice of `i`. Hence `20`.

3. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/78794955-ae166980-79b4-11ea-87c8-84b7bd12f06f.png)

The function `Replace` in the package `string` returns a modified copy of the input string (`"ianianian"`) by replacing the first two instances of the substring `"ni"` with `"in"`. Hence `iainainan`.

4. What is printed by this code?

![](https://user-images.githubusercontent.com/49638680/78795280-1f561c80-79b5-11ea-9d77-49c9d0a148ac.png)

A _switch_ statement. The variable `x` is the _tag_ and its value before entering the switch is $7$. The first case expression is evaluated `x > 3`, since this is true the code block is executed and the program exits the switch.
Hence it prints `1`.

5. What is printed by this code?

![](https://user-images.githubusercontent.com/49638680/78795599-78be4b80-79b5-11ea-9724-7c4ede0a456b.png)

In this code a for loop on the variable `x` runs from $0$ to $4$ (adding one at each step). This is an implementation of Fibonacci sum.
Indeed, when `x = 4`, we have `x2 = 8`.

In any case, one can copy/paste this code in [GOlang online compiler](https://play.golang.org/) and make the code run.

__Note__ You are supposed to find the answer by just reading at the code, not making it run.

6. This code compiles correctly. __True__ or __False__:

![](https://user-images.githubusercontent.com/49638680/78796566-cab3a100-79b6-11ea-822f-5a354abd605d.png)

`&` operator refers to a variable and returns the memory address that variable is located into.
The code attempts to assign `x` to be a pointer to a pointer, but `x` is declared to be an `int`, so this causes an error.

7. Which integer type provides higher accuracy?

![](https://user-images.githubusercontent.com/49638680/78799044-f421fc00-79b9-11ea-9cdb-9588a9869af0.png)

The accuracy is exactly the same, since we are talking about integers. The number of bits reflects the range (min, max) that can be represented.
