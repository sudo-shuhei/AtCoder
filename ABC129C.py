N,M=map(int,input().split())
A=[int(input()) for i in range(M)]
#print(N)
#print(M)
#root=[0]*N
def root(i): #0からiにいく組み合わせ
    if i==0:
        root[0]=1
        return 1
    elif i==1:
        root[1]=1
        return 1
    elif i==2:
        root[2]=2
        return 2

    elif i<0:
        print(0)
        exit()
    else:

        return root(i-1)+root(i-2)
root=[0]*N
for i in range(N+1):
    if i==0:
        root[0]=1
    elif i==1:
        root[1]=1
    elif i==2:
        root[2]=2
    else:
        root[i]=root[i-1]+root[i-2]

if M==0:
    print(root[N])
    exit()
#print(root(5))
#print(A)
ans=0
for i in range(M):
    if i==0:
        ans=root[A[0]-1]
    else:
        if A[i]-A[i-1]-2<0:
            print(0)
            exit()
        else:
            ans*=root[A[i]-A[i-1]-2]
    #print(ans)
    #print(i)
ans*=root[N-A[M-1]-1]
print(ans%1000000007)
