A = int(input())
B = int(input())
s = set()
s.add(A)
s.add(B)
for i in (1,2,3):
    if i not in s:
        print(i)
        exit()
