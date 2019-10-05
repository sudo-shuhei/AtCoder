s=input()
ans=[]
I=[0]*len(s)
for i in range(len(s)):
    if (s[i]=="A" or s[i]=="T" or s[i]=="G" or s[i]=="C"):
        I[i]=1

    else:
        I[i]=0
for i in range(len(I)):
    if (I[i]==1 and I[i-1]==0) or (I[i]==1 and i==0):
        ans.append(1)
    elif (I[i]==1 and I[i-1]==1):
        ans[-1]+=1
if ans==[]:
    print(0)
else:
    print(max(ans))
