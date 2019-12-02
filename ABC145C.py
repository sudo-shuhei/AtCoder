import math
import itertools
n = int(input())
x = [0]*n
y = [0]*n
for i in range(n):
    x[i],y[i] = [int(j) for j in input().split()]
# print(n,x,y)
dict = {}
for i in range(n):
    for j in range(n):
        if i < j:
            dict[str(i)+str(j)] = math.sqrt((x[i] - x[j])**2 + (y[i]- y[j])**2)

# print(dict)
# perm = [n for n in range(n)]
# print(perm)
sumation = 0
# print(len([v for v in itertools.permutations(perm,n)]))
# for p for perm:
#     for i in p: #それぞれのまち
#         if i == 0:
#             continue
#         else:
#             sum
# list = [v for v in dict.values()]
# # print(type(list))
# sumation = sum(list)
# # print(sumation)
# sumation = sumation * math.factorial(n)
# # print(sumation)
# for kei, length in dict.items():
#     if n != 2:
#         sumation -= length*2*math.factorial(n-2)
#     else:
#         continue
# # print(sumation)
# print(sumation/math.factorial(n))
for kei,length in dict.items():
    sumation += (n-1) * math.factorial(n-2)*length*2
print(sumation/math.factorial(n))
