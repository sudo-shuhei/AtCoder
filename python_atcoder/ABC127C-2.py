N,M=map(int,input().split())
list1=[list(map(int,input().split())) for i in range(M)]
start=[]
end=[]
for i in range(M):
    start.append(list1[i][0])
    end.append(list1[i][1])
stmax=max(start)
#print(stmax)
enmax=min(end)
#print(enmax)
if enmax>=stmax:
    print(enmax-stmax+1)
else:
    print(0)
