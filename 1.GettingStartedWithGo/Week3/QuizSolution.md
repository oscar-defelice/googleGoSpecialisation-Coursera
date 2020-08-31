## Week 3 Quiz

This quiz is all about composite data type. Questions are all relative to the results of some little piece of code.
One can copy/paste this code in [GOlang online compiler](https://play.golang.org/) and make the code run to find the correct answer.

However, since in any question, the code is quite easy to understand, you are supposed to find the answer by just reading at the code, not making it run.

1. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/79245982-2aa7bd00-7e79-11ea-99db-dd48f9a1a866.png)

The for cycle substitutes the value of `y` with the value inside the slice anytime the latter is greater than the former.
Hence $8$.

2. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/79246033-3eebba00-7e79-11ea-9e7f-ee88cd25a261.png)

The slices are modified, and so the underlying array, that is `x` in this case.
Thus, `[1 8 3]`.

3. What is printed when the following program is executed?

![]("https://user-images.githubusercontent.com/49638680/79246042-41e6aa80-7e79-11ea-8132-76512212cba9.png)

This code prints lengths and capacities of slices `y` and `z`, having the same underlying array `x`.

4. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/79246060-44e19b00-7e79-11ea-903e-65d63515e5e8.png)

We have a map definition here.
We loop through the map and we print key/value pair when the key is equal to "harris". The corresponding value is $2$.
Hence the answer is `harris2`.

5. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/79246077-49a64f00-7e79-11ea-9d29-5696ed6f4341.png)

In this code a for loop on the array `a` of P structures runs.
It compares the `y` field of the `struc` in the loop with the `b` one defined outside.
If the latter is smaller than the former it reassign `b` to this structure.
Hence, at the end `fmt.Printf(b)` would return `P{"a", 10}` and so the answer is `a`.

6. What is printed when the following program is executed?

![](https://user-images.githubusercontent.com/49638680/79246102-5034c680-7e79-11ea-883a-2ff672c9e8de.png)

This is an append procedure to a slice.
Initially `s` has zero length and capacity $3$. After the append the underlying array has been modified such that it has $100$ in its first position now and `s` has length $1$.
Capacity is still untouched since we do not have reached the size of the array.
