from math import factorial

def cmb(n, r, mod):
    if ( r<0 or r>n ):
        return 0
    r = min(r, n-r)
    return g1[n] * g2[r] * g2[n-r] % mod

mod = 10**9+7 #出力の制限
N = 10**4
g1 = [1, 1] # 元テーブル
g2 = [1, 1] #逆元テーブル
inverse = [0, 1] #逆元テーブル計算用テーブル

for i in range( 2, N + 1 ):
    g1.append( ( g1[-1] * i ) % mod )
    inverse.append( ( -inverse[mod % i] * (mod//i) ) % mod )
    g2.append( (g2[-1] * inverse[-1]) % mod )

# a = cmb(n,r,mod)

X,Y = [int(s) for s in input().split()]
# print(x,y)
if (X+Y) %3 != 0:
    print(0)
    exit()
n = int((2*Y - X)/3)
m = int(Y - 2*n)
# print(n,m)
# result = factorial(n+m)/factorial(n)
result = cmb(n+m, n, mod)
big = 10**9+7
print(result)
