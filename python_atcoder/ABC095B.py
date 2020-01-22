N,X = map(int,input().split())
m = [0]*N
for i in range(N):
    m[i] = int(input())
remain = X - sum(m)
r = remain // min(m)
print(r + N)
