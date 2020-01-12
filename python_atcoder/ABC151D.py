# import numpy as np
from itertools import combinations
import copy
from collections import deque

H, W = map(int,input().split())
# print(H,W)

# S = np.zeros((H,W))
S = [[0]*W for i in range(H)]
# print(S)
for i in range(H):
    scan =  input()
    # print(scan)
    for j,s in enumerate(scan):
        if s == ".":
            S[i][j] = 0
        else:
            S[i][j] = 1
# print(S)
positions = []
for i in range(H):
    for j in range(W):
        positions.append((i,j))
# print(positions)
c = combinations(positions,2)
result = 0
for pat in c:
    start = pat[0]
    stop = pat[1]
    distance = 0
    S1 = copy.deepcopy(S)
    q = deque()
    # d = np.full((H,W),np.inf)
    d = [[float('inf')]*W for i in range(H)]

    q.append(start)
    d[start[0]][start[1]] = 0
    while len(q) > 0:
        p = q.popleft()
        # print(p)
        i,j = p[0],p[1]
        four = [(i-1,j),(i+1,j),(i,j-1),(i,j+1)]
        if p == stop:
            break
        for f in four:
            i1,j1 = f[0],f[1]
            if 0 <= i1 < H and 0<= j1 < W and S1[i1][j1] == 0 and d[i1][j1] == float('inf'):
                q.append((i1,j1))
                d[i1][j1] = d[i][j] +1
    distance = d[stop[0]][stop[1]]
    # print(distance)
    if distance != float('inf') and distance > result:
        result = int(distance)

print(result)
