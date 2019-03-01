import sys

cache = {}

def memoize(fn):
    def wrapper(*args):
        if args in cache:
            return cache[args]
        value = fn(*args)
        cache[args] = value
        return value
    return wrapper

@memoize
def lcs(s, t):
    if not s or not t:
        return ""
    if s[-1] == t[-1]:
        return lcs(s[:-1], t[:-1]) + s[-1]
    a = lcs(s, t[:-1])
    b = lcs(s[:-1], t)
    return a if len(a) > len(b) else b

def main():
    with open(sys.argv[1], "r") as fp:
        for testcase in fp:
            testcase = testcase.strip()
            if testcase:
                s, t = testcase.split(";")
                print(lcs(s, t))

if __name__ == "__main__":
    main()
