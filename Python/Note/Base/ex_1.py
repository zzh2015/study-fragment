# coding: utf-8

# """...""" 用于生成注释文档
print """
	python 基础
	基本使用代码片段
"""
# 格式化打印 %r, 打印多个 %(ll, i_num, s_name)
ll = [12345, 'Hello', 3.14]
i_num = 1024
s_name = 'lily'

print "s_name: %s" % s_name
print "i_num: %d" % i_num
print "ll: %r" % ll
print "%r, %d, %s" % (ll, i_num, s_name)

# // floor division, ** 幂乘方
# 交互式模式中，上一次的结果保存在 _ 中.
print "6.28 // 3.14 = ", 6.28 // 3.14
print "2^10 =  ", 2 ** 10

# python中内建复数, 2-3j || 2-3J || complex(2, -3)
cc = complex(2, -3)
print cc.real, cc.imag
print "(1+2j)/(1+1j) ", (1+2j)/(1+1j)

# 屏蔽 / 转义语义
print r"C:\program\note"

# word = ['hc', 'xc', 'hc'], 索引可以为负数代表从右开始，切片
word = [3.14, 235813, 'Hello world']
print word[-1], word[:2]

# 创建unicode文本
print u"Hello\u0020Unicode!"
