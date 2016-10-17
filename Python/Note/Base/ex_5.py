#!/usr/bin/env python
# encoding: utf-8

import sys

print """\
    syntax errors and exceptions in python
"""

while True:
    try:
        x = int(raw_input("Please enter a number: "))
        break
    except ValueError: # (RuntimeError, TypeError, NameError):
        print 'No valid number. Try again...'

try:
    f = open("ex_1.py", 'rb')
    s = f.readline()
    i = int(s.strip())
except IOError as e:
    print "I/O error({0}): {1}".format(e.errno, e.strerror)
except ValueError:
    print "Could not convert data to an integer."
except: # * 通配其他所有情况
    print "Unexcepted error:", sys.exc_info()[0]
    raise
# else: # else 子句, 未发生异常时执行
finally: # 无论异常是否发生都会执行
    print 'Leave...'
f.close()
