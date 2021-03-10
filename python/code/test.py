s = ('exe')
if '' in s:
    print("in",s)#ok
else:
    print("not in",s)

t = ['exe']
if '' in t:
    print("in",t)
else:
    print("not in",t)#ok

t = ('exe',)
if '' in t:
    print("in",t)
else:
    print("not in",t)#ok

s1 = 'exe'
if '' in s1:
    print("in",s1)#ok
else:
    print("not in",s1)
