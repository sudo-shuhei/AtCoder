import copy
x,y = [int(s) for s in input().split()]
# print(x,y)
def replace_zahyo(list, sum):
    print(list,sum)
    list1 = copy.deepcopy(list)
    if list == []: #リスtがからで終わり
        return sum
    for first in list1:
        if first == (0,0): #0,0にたどり着いたら加算
            sum+=1
            list.remove(first)
            return replace_zahyo(list,sum)
        elif first[0] <= 0 or first[1] <= 0: #外に出たら終わり
            list.remove(first)
            return replace_zahyo(list,sum)
        else:#普通
            list.append((first[0]-1, first[1]-2))
            list.append((first[0]-2, first[1]-1))
            list.remove(first)
            return replace_zahyo(list,sum)

def replace_dict(dict,sum):
    # print(dict)
    if dict == {}:
        return sum
    # print(dict)
    # print(dict.keys())
    first_key = sorted(dict.keys())[0]
    first_value = dict.pop(first_key)
    # print(first_key,first_value)
    if first_key == (0,0):
        sum += first_value
        return replace_dict(dict,sum)
    elif first_key[0] <= 0 or first_key[1] <= 0:
        return replace_dict(dict,sum)
    else:
        dict.setdefault((first_key[0]-1,first_key[1]-2),0)
        dict.setdefault((first_key[0]-2,first_key[1]-1),0)
        dict[(first_key[0]-1,first_key[1]-2)]+= first_value
        dict[(first_key[0]-2,first_key[1]-1)]+= first_value
        return replace_dict(dict,sum)


print(replace_dict({(x,y):1},0)%(10**9+7))
