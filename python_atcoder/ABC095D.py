import copy

N,C = map(int,input().split())
x = [0]*N
v = [0]*N
for i in range(N):
    x[i], v[i] = map(int,input().split())

def rec(i,x,v,position, total_cal):
    pos, cal = x.pop(i),v.pop(i)
    total_cal += cal - abs(pos - position)
    xr,xl = copy.copy(x),copy.copy(x)
    vr,vl = copy.copy(v),copy.copy(v)
    rec(i+1,xr,vr,pos,total_cal)
    rec(i-1,xl,vl,pos,total_cal)

Class susi():
    def __init__(self,x,cal):
        self.position = x
        self.cal = cal
        self.right = None
        self.left = None
