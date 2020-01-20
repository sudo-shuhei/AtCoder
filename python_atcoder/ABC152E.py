import fractions
from functools import reduce

def gcd_list(list1):
    ans=list1[0]
    for i in range(len(list1)):
        ans=fractions.gcd(ans,list1[i])
    return ans
#リスト内の最大公約数を求める

def lcm_base(x,y):
    return ((x*y) // fractions.gcd(x,y))

def lcm_list(numbers):
    return reduce(lcm_base, numbers, 1)

large_num = 10**9+7
N = int(input())
A = list(map(int,input().split()))
ans = A[0]
for i in range(1,N):
    ans = ans * A[i] // fractions.gcd(ans,A[i])
lcm_all = ans
result = 0
for num in A:
    result += lcm_all//num
print(result%large_num)
