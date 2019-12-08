from collections import defaultdict

def fib(n):
    memo = defaultdict(int)
    if n <= 1:
        return n
    if memo[n] != 0:
        return memo[n]
    else:
        # print(n)
        memo[n] = fib(n-1) + fib(n-2)
    return memo[n]

print(fib(40))
