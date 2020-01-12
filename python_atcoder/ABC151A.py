C = input()
# print(C+1)
numc = lambda c: ord(c) - ord('a') + 1
num2alpha = lambda c: chr(c)
# print(numc(C))
# print(num2alpha(numc(C)+1))
print(chr(ord(C)+1))
