#!/usr/bin/env python
# encoding: utf-8

print """\
    python 流程控制
"""

# if
# x = int(raw_input("Please enter an integer: "))
x = 1024
if x < 0:
    x = 0
    print 'N changed to zero'
elif x == 0:
    print 'Zero'
else:
    print 'More'

# for, range(), 数值序列
print 'range(10): ', range(10)
print 'range(5, 10): ', range(5, 10)
print 'range(0, 10, 3)', range(0, 10, 3)

r_w = ['Mary', 'had', 'a', 'little', 'lamb']
for i in range(len(r_w)):
    print i, r_w[i]

# break和continue语句, 循环中的else子句
for n in range(2, 10):
    for nn in range(2, n):
        if n % nn == 0:
            print n, 'equals', x, '*', n/x
            break
    else:
        print n, 'is a prime num'
# pass 空语句

