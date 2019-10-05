import math
N,M=map(int,input().split())
sw=[list(map(int,input().split())) for i in range(M)]
for i in range(len(sw)):
    del sw[i][0]
p=list(map(int,input().split()))
#print(sw)
ans=0

for i in range(2**N): #i番目のスイッチの組み合わせ
    bini=format(i,"b")
    #print(bini)
    num="0"*(len(format(2**N-1,"b"))-len(bini))+bini
    #print(num)
    #print(len(format(2**N-1,"b")))
    #print(len(bini))
    #print(format(2**N-1,"b"))

    light=[0]*M
    for j in range(M):#j番目の電球がおんかオフか
        int1=0
        #print("j="+str(j))
        for k in sw[j]:#スイッチk
            #print("k="+str(k))
            int1+=int(num[k-1])
        if int1%2==p[j]:
            light[j]=1
    if light==[1]*M:
        ans+=1
print(ans)
#AC
