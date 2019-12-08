N=int(input())
A,B,C,D,E=[int(input()) for i in range(5)]
bn=min(A,B,C,D,E)
if N%bn==0:
    m=N//bn-1
else:
    m=N//bn
print(5+m)

#AC
