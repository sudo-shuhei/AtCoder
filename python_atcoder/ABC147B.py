S = input()
# print(S)

rs = reversed(S)
# print(rs)

result = 0
for i, char in enumerate(rs):
    if S[i] != char:
        result += 1

print(int(result/2))
