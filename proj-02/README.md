# Problem 2 - 

# Description

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be: 

$$ 1,2,3,5,8,13,21,34,55,89,... $$

Find the sum of all the even-valued terms in the sequence which do not exceed four million. 

## Solution

Totally subsiquent term of the Fibonacci sequence based on the following formula:
> $$ term_{n} = term_{n-1} + term_{n-2} $$

I solve this problem using two methods. First one is generally a mathematical way, but the second one has a more programmatic solution due to algorithms.

1. Using a total formula[^1]:

> $$ F_n = { {1\over\sqrt5}({1+\sqrt5\over2})^n } - { {1\over\sqrt5}({1-\sqrt5\over2})^n } $$

2. Using the Fibonacci algorithm.


| Language       | Path (link)                  |
| :------------- | :--------------------------- |
| Go (math)      | [`formula.go`](./formula.go) |
| Go (algorithm) | [`regular.py`](./regular.go) |

<details>
  <summary>Final result</summary>
  
  ### 4613732

</details>

[^1]: https://en.wikipedia.org/wiki/Fibonacci_number