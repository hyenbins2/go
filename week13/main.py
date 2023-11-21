# python의 deep copy 설명
import copy

a = [3, -9, 11]
b = a  # 참조
c = a[:]  # 다른 메모리 공간에 copy
d = list(a)  # 다른 메모리 공간에 copy
e = a.copy()  # 다른 메모리 공간에 copy

print(a, b, c, d, e)
a[1] = 100
print(a, b, c, d, e)

print("=" * 70)

a = [3, [5, -71], 11]
b = a  # 참조
c = a[:]  # shallow copy
d = list(a)  # shallow copy
e = a.copy()  # shallow copy
f = copy.deepcopy(a)  # deep copy

print(a, b, c, d, e)
# a[1] = 100
a[1][0] = 100  
print(a, b, c, d, e, f)
