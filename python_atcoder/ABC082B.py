s = input()
t = input()

s = sorted(s)
t = sorted(t)

# print(s,t)
if s[0] < t[-1]:
    print('Yes')
    exit()
elif s[0] > t[-1]:
    print('No')
    exit()
elif len(set(s)) ==1 and len(set(t)) ==1 and len(s) < len(t):
    print('Yes')
    exit()
else:
    print('No')
