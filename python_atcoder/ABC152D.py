N = int(input())

cnt = [[0]*10 for i in range(10)]
# print(cnt)

for i in range(1,N+1):
    s = int(str(i)[0])
    e = int(str(i)[-1])
    cnt[s][e] += 1

result = 0
for j in range(10):
    for k in range(10):
        result += cnt[j][k] * cnt[k][j]

print(result)
