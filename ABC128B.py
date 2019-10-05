N=int(input())
info=[input().split() for i in range(N)] #mojiretu
#print(info)
from operator import itemgetter
dic1={}
for i in range(len(info)):
    dic1[i+1]=info[i]
#print(dic1)
info1=sorted(info,key=lambda x:int(x[1]),reverse=True)
info2=sorted(info1,key=itemgetter(0))
#print(info2)
for i in info2:
    for key,value in dic1.items():
        if value==i:
            print(key)
