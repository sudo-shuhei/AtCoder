N,K=map(int,input().split())
ans=0
for i in range(N):
    a=i
    j=0
    while a<K:
        j+=1
        a*=2
    ans+=(1/N)/a

print(ans)
