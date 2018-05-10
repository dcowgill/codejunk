
def lit(s):       return lambda Ns: set([s]) if len(s) in Ns else null
def alt(x, y):    return lambda Ns: x(Ns) | y(Ns)
def star(x):      return lambda Ns: opt(plus(x))(Ns)
def plus(x):      return lambda Ns: genseq(x, star(x), Ns, startx=1)
def oneof(chars): return lambda Ns: set(chars) if 1 in Ns else null
def seq(x, y):    return lambda Ns: genseq(x, y, Ns)
def opt(x):       return alt(epsilon, x)
dot = oneof('?')
epsilon = lit('')

null = frozenset()

def genseq(x, y, Ns, startx=0):
    if not Ns or Ns == set([0]):
        return null
    return set(a + b
               for a in x(range(max(Ns) + 1))
               for b in y([n - len(a) for n in Ns if n >= len(a)])
               if len(a) + len(b) in Ns)
