N = int(input())
a = list(map(int, input().split()))

count = 1
for i in a:
    if i == count:
        count +=1
if count-1 == 0:
    print(-1)
    exit()
# print(count-1)
print(N-count+1)
