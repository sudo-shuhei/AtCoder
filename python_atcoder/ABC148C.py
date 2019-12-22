# import math
a, b= map(int,input().split())
def lcm(x, y):
    return (x * y) // gcd(x, y)

def gcd(a,b):
  while b!=0:
    a,b=b,a%b
  return a
 # print(gcd(a,b))

print(lcm(a,b))
