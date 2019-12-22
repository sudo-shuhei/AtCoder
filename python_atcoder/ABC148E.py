N = int(input())
if N %2 == 1:
    print(0)
    exit()

# print(N//10)

# result = 0
# for i in range(19):
#     print(N//(10**(i+1)))
#     result += N//(10**(i+1))

# print(result)

result = 0
for i in range(100):
    # print(N//((5**(i+1))*2))
    result += N//((5**(i+1))*2)
print(result)
