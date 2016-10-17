#!/usr/bin/env python
# encoding: utf-8

print """\
    class in python
"""

class MyClass:
    i = 1024 # 共享变量
    def __init__(self):
        self.name = 'MyClass' # 实例变量
    def f(self):
        return 'hello world'
    def cname(self):
        print self.name
# class DerivedClassName(modname.BaseClassName):
x = MyClass()
print x.f()
x.cname()
# 私有变量和类本地引用
# 以一个下划线开头的命名处理为API非公开的部分

# iter()
class Reverse:
    """def iter()"""
    def __init__(self, data):
        self.data = data
        self.index = len(data)
    def __iter__(self):
        return self
    def __next__(self):
        if self.index == 0:
            raise StopIteration
        self.index = self.index - 1
        return self.data[self.index]

rev = Reverse('spam')
iter(rev)
for ch in rev:
    print(ch)
