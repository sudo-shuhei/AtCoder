N,M=map(int,input().split())
list1=[list(map(int,input().split())) for i in range(M)]
list2=list(range(N+1))
ans=set(range(N+1))
for i in range(M):
    lr=set(list2[list1[i][0]:list1[i][1]+1])
    #print(lr)
    ans=ans&l
#print(ans)
print(len(ans))
