N,M=map(int,input().split()) #N...ken M..hon
A=[list(map(int,input().split())) for i in range(N)]
A=sorted(A) #A...en B..hon
#print(A)
price=0
num=0
for i in range(N):
    if num+A[i][1]<M:
        price+=A[i][0]*A[i][1]
        num+=A[i][1]
    elif num<M and num+A[i][1]>=M:
        price+=A[i][0]*(M-num)
        break
    else:
        break
print(price)
