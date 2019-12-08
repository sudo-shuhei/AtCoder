n = int(input())
s = input()

if n%2 == 1:
    print("No")
#print(type(n/2+1))
# print (s[0:int(n/2)] )
# print(s[int(n/2):])
elif s[0:int(n/2)] == s[int(n/2):]:
    print("Yes")
else:
    print("No")
