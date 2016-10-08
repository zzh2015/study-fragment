# coding: utf-8

print """\
	python常用的数据结构
"""

# list
ll = [1, 'eggs', '3.14', 2048]
print ll
ll_2 = ll[:]
print ll_2

ll.pop()
ll.remove(1)
del ll[0]
print ll
print ll_2

# filter(), map(), reduce()
def f(x): return x % 3 == 0 or x % 5 == 0
print filter(f, range(2, 25))

def cube(x): return x**3
print map(cube, range(1, 11))

def add(x, y): return x+y
print reduce(add, range(1, 101))

# list推导式
ss = [x**2 for x in range(10)]
print 'ss: ', ss
print [(x, y) for x in range(1, 4) for y in [3, 1, 4] if x != y]

# tuple 由数个逗号分隔的值组成
t = 12345, 3.14, 'Hello!'
print 'tuple: ', t

# set 无序不重复元素的集
basket = ['apple', 'orange', 'pear']
fruit = set(basket)
print fruit

# dict 字典-值对集合
tel = {'jack': 4098, 'sape': 4139}
print tel['jack']
