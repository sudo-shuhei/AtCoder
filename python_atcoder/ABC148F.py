from collections import deque

def shortest(x,y,arr): #スタックと隣接行列で最短距離を求める
    d = deque()
    d.append((u,0))
    flag = [0]*len(arr)
    # flag[u] = 1
    while True:
        # print(d)
        di = d.pop()
        if flag[di[0]] ==1:
            continue
        flag[di[0]] =1
        if di[0] == y:
            return di[1]
        near_list = arr[di[0]]
        for i in range(len(near_list)):
            # print(i)
            if flag[i] == 1:
                continue
            if near_list[i] == 1:
                d.append((i,di[1]+1))


N,u,v = map(int,input().split())

arr = [[0]*N for i in range(N)]
# print(arr)

for i in range(N-1):
    ai, bi = map(int,input().split())
    ai -= 1
    bi -= 1
    arr[ai][bi] = 1
    arr[bi][ai] = 1
# print(arr)
u -=1
v -=1
print(u,v)
print(shortest(u,v,arr))
