from itertools import combinations

N = int(input())
syogens = []
for i in range(N):
    a = int(input())
    syogen = []
    for j in range(a):
        x,y= map(int,input().split())
        syogen.append((x,y))
    syogens.append(syogen)
#人Aiの証言はsyogens[i]
#j番目の証言は　syogens[i][j]
# print(syogens)

for k in reversed(range(N)): #kにんの正直者がいると仮定
    # print(k)
    comb = combinations(range(N),k+1)
    # print(*comb)
    for pattern in comb: #(0,3,5,6)
        l = [-1]*N
        b=1
        syogen = []
        for i in range(N):
            if i in pattern:
                l[i] = 1
            else:
                l[i] = 0

        # print(l)
        #発言を検証していく
        for person in pattern:
            syogen = syogens[person]
            for s in syogen:
                # print(s)
                target = s[0]
                value = s[1]
                if value != l[target-1]:
                    b = 0
        if b == 1:
            print(k+1)
            exit()

print(0)
