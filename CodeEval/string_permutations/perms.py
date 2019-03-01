import sys

def permutations(s):
    if s:
        for i, c in enumerate(s):
            for p in permutations(s[:i] + s[i+1:]):
                yield [c] + p
    else:
        yield []

def main():
    with open(sys.argv[1], "r") as fp:
        for testcase in fp:
            s = testcase.strip()
            if s:
                print(",".join("".join(p) for p in permutations(sorted(s))))

if __name__ == "__main__":
    main()
