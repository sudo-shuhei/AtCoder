import numpy as np

def selection_sort(l): #選択ソート O(n**2)
    i = 0
    while i < len(l):
        min = np.inf
        mini = -1
        for j,v in enumerate(l):
            if j < i:
                continue
            if v < min:
                min = v
                mini = j
        l[i],l[mini] = min, l[i]
        i += 1
    return l

l = [2,5,1,7,10,3,6,9,4,8]
# print(selection_sort(l))

def insertion_sort(l): #挿入ソート O(n**2)
    for i in range(len(l)):
        num = l[i]
        # print(num,l[:i])
        for j in range(i):
            if l[j] >= num:
                l[j+1:i+1] = l[j:i]
                l[j] = num
                break
    return l

# print(insertion_sort(l))

def bucket_sort(l): #バケツソート　値のとる範囲が既知である必要　O(n)
    max_num = max(l)+1
    bucket = [-1] * max_num
    for num in l:
        bucket[num] = num
    result = []
    for i in bucket:
        if i != -1:
            result.append(i)
    return result

# print(bucket_sort(l))

def bubble_sort(l): #バブルソート　O(n**2)
    for i in range(len(l)):
        for j in reversed(range(i,len(l)-1)):
            if l[j] > l[j+1]:
                l[j],l[j+1] = l[j+1],l[j]
    return l

# print(bubble_sort(l))

def merge_sort(l): #マージソート　O(nlog(n))
    if len(l) < 2:
        return l
    c = len(l)//2
    return merge(merge_sort(l[:c]), merge_sort(l[c:]))

def merge(x,y):
    if len(x) < 1:
        return y
    if len(y) < 1:
        return x
    if (x[0] > y[0]):
        return [y[0]]+ merge(x, y[1:])
    else:
        return [x[0]]+ merge(x[1:], y)

print(merge_sort(l))
