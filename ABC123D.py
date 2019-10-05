X,Y,Z,K=map(int,input().split())
A=list(map(int,input().split()))
B=list(map(int,input().split()))
C=list(map(int,input().split()))
a=sorted(A,reverse=True)
b=sorted(B,reverse=True)
c=sorted(C,reverse=True)
q=[]
ans=[0]*K

q.append((a[0]+b[0]+c[0],1,1,1))
#q.append((0,1,1,1))
i=0
j=0
k=0

for l in range(K):
    #ans[l]=q[0]
    print(q[0][0])
    del q[0]
    if (i+1<=X-1 and j<=Y-1 and k<=Z-1) and ((a[i+1]+b[j]+c[k],i+1,j,k) not in q) :
        q.append((a[i+1]+b[j]+c[k],i+1,j,k))
    if (i<=X-1 and j+1<=Y-1 and k<=Z-1) and ((a[i]+b[j+1]+c[k],i,j+1,k) not in q):
        q.append((a[i]+b[j+1]+c[k],i,j+1,k))
    if (i<=X-1 and j<=Y-1 and k+1<=Z-1) and ((a[i]+b[j]+c[k+1],i,j,k+1) not in q):
        q.append((a[i]+b[j]+c[k+1],i,j,k+1))
    q=sorted(q,reverse=True)
    #print(q)
    if q==[]:
        i=0
        j=0
        k=0
    else:
        i=q[0][1]
        j=q[0][2]
        k=q[0][3]
    #print(q)
#[print(ans[i][0]) for i in range(k)]
