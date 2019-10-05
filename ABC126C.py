N,K=map(int,input().split())
if N>=K:
    ans=(N-K+1)/N
else:
    ans=0
#while 2**(i-1)<K:
    #ans+=1/2**i
    #
for i in range(K-1):
    j=1
    a=0
    while a<K:
        j+=1
        a=2**(j-1)*i
    ans+=(1/N)*1/(2**j)


print(ans)
