n = int(input())
a = list(map(int, input().split()))
# print(*A)

# print(bin(A[0]))
# print(format(A[0], "b"))
MOD = (10**9) +7
result = 0

cnt = [0] *60
for num in a:
    for j in range(60):
        cnt[j] += (num >> j) %2

ans = 0
for b in range(60):
    ans += cnt[b] * (n - cnt[b]) << b
    ans %= MOD

print(ans)
