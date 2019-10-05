import fractions
def gcd_list(list1):
    ans=list1[0]
    for i in range(len(list1)):
        ans=fractions.gcd(ans,list1[i])
    return ans
#リスト内の最大公約数を求める
