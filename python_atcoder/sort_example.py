import numpy as np
from collections import defaultdict

l1 = [123,973,206,215,77,527,426,885,300,681,732,994,258,770,359,512,863,44,137,905]
l = [2,5,1,7,10,3,6,9,4,8]

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

def radix_sort(l): #基数ソート　m進数k桁ならO(kN)
    l_str = [str(i) for i in l]
    max_radix = 0
    for i in l_str:
        if len(i) > max_radix:
            max_radix = len(i)
    bucket = defaultdict(list)
    # print(bucket)
    array = l
    for m in range(1,max_radix+1):
        # print(m)
        r = 10**(m-1)
        for val in array:
            key = int(val/r) %10
            bucket[key].append(val)
        j = 0
        for i in range(10):
            values = bucket[i]
            for val in values:
                array[j] = val
                j+=1

        for j in range(len(bucket)):
                bucket[j] = list()
    return array

# print(radix_sort(l1))


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

# print(merge_sort(l))

def min_heapify(array, i): #取り出した後にヒープに直す
    left = 2 * i + 1
    right = 2 * i + 2
    length = len(array) - 1
    smallest = i
    if left <= length and array[i] > array[left]:
        smallest = left
    if right <= length and array[smallest] > array[right]:
        smallest = right
    if smallest != i:
        array[i], array[smallest] = array[smallest], array[i]
        min_heapify(array, smallest)

def build_min_heap(array): #最初にヒープを作る
    for i in reversed(range(len(array)//2)):
        min_heapify(array, i)

def heap_sort_asc(array): #ヒープソート　O(NlogN)
    array = array.copy()
    build_min_heap(array)
    sorted_array = []
    for _ in range(len(array)):
        array[0], array[-1] = array[-1], array[0]
        sorted_array.append(array.pop())
        min_heapify(array, 0)

    return sorted_array

# print(heap_sort_asc(l1))

def quick_sort(arr):
    left = []
    right = []
    if len(arr) <= 1:
        return arr

    # データの状態に左右されないためにrandom.choice()を用いることもある。
    # ref = random.choice(arr)
    ref = arr[0]
    ref_count = 0

    for ele in arr:
        if ele < ref:
            left.append(ele)
        elif ele > ref:
            right.append(ele)
        else:
            ref_count += 1
    left = quick_sort(left)
    right = quick_sort(right)
    return left + [ref] * ref_count + right

print(quick_sort(l1))
