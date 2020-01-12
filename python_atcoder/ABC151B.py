N,K,M = map(int, input().split())
A = list(map(int,input().split()))
# print(N,K,M,A)

need_sum = N * M
now_sum = sum(A)
need_score = need_sum - now_sum
if need_score > K:
    print(-1)
    exit()
if need_score < 0:
    print(0)
    exit()
print(need_score)
