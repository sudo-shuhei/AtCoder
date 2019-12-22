N = int(input())
S, T = input().split()
# print(N,S,T)
result = ""
for s, t in zip(S,T):
    result += s
    result += t
print(result)
