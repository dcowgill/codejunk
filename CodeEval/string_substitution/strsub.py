import sys

def substitue_strings(s, pairs):
    return "".join(s for s, _ in substitute_chain([[s, False]], pairs))

# Given a chain of [string, modified?] pairs, and a list of (from, to)
# tuples to replace, return a new chain with all replacements applied to
# any unmodified links in the chain.
def substitute_chain(chain, pairs):
    if not pairs:
        return chain
    newchain = []
    for link in chain:
        s, modified = link
        if not modified:
            newchain.extend(substitute(link[0], pairs[0]))
        else:
            newchain.append(link)
    return substitute_chain(newchain, pairs[1:])

# Apply a single (from, to) replacement to the given string, returning a
# chain of [string, modified?] pairs.
def substitute(value, replacement):
    s, t = replacement
    chain = []
    while value:
        i = value.find(s)
        if i < 0:
            chain.append([value, False])
            return chain
        if i > 0:
            chain.append([value[:i], False])
        chain.append([t, True])
        value = value[i + len(s):]
    return chain

def main():
    with open(sys.argv[1], "r") as fp:
        for testcase in fp:
            testcase = testcase.strip()
            if testcase:
                s, fr = testcase.split(";")
                fr = fr.split(",")
                pairs = [fr[i:i+2] for i in range(0, len(fr), 2)]
                print(substitue_strings(s, pairs))

if __name__ == "__main__":
    main()
