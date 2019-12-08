N=int(input())
W=list(map(int,input().split()))
S1=0
S2=sum(W)
#print(S1)
#print(S2)
#print(W)
n=float("inf")
for i in range(N-1):
    S1=S1+W[i]
    S2=S2-W[i]
    n=min(n,abs(S1-S2))
print(n)
