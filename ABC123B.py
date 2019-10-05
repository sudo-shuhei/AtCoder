a,b,c,d,e=[int(input()) for i in range(5)]
a1=a//10
a2=a%10
if a2==0:
    a1-=1
    a2=10
b1=b//10
b2=b%10
if b2==0:
    b1-=1
    b2=10
c1=c//10
c2=c%10
if c2==0:
    c1-=1
    c2=10
d1=d//10
d2=d%10
if d2==0:
    d1-=1
    d2=10
n=e//10
m=e%10
if m==0:
    n-=1
    m=10
min1=min(a2,b2,c2,d2,m)
ans=(a1+b1+c1+d1+n+4)*10+min1
print(ans)
#AC
