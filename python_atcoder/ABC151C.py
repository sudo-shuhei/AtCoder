from collections import defaultdict
N, M = map(int,input().split())
submit = defaultdict(list)
for i in range(M):
    problem_num, S = input().split()
    problem_num = int(problem_num)
    if S == "AC":
        S = 1
    else:
        S = 0
    submit[problem_num].append(S)
# print(submit)

ac = 0
penalty = 0
for l in submit.values():
    if 1 in l:
        ac += 1
        for i in range(len(l)):
            if l[i] == 0:
                penalty +=1
            elif l[i] == 1:
                break
print(str(ac)+" "+str(penalty))
