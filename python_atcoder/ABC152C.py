N = int(input())
P = list(map(int,input().split()))

min = P[0]
# print(N,P,min)
result = 0
for pi in P:
    if pi <= min:
        result += 1
        min = pi
print(result)
