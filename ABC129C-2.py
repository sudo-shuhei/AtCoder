N,M=map(int,input().split())
A=[int(input()) for i in range(M)] #0からM-1まで


root=[0]*(N+2)   #間がiの時の行き方
for i in range(N+2):    #0からN-1まで
    if i==0:
        root[0]=0
    elif i==1:
        root[1]=1
    elif i==2:
        root[2]=1
    else:
        root[i]=root[i-1]+root[i-2]
#print(root)

ans=0
if M==0:
    print(root[N+1])
    exit()

if N==0:
    print(0)
    exit()

for i in range(M+1): #0からM
    if i==0:
        ans=root[A[0]]
    elif i==M:
        ans*=root[N-A[M-1]]
    else:
        ans*=root[A[i]-A[i-1]-1]
print(ans%(10**9+7))
