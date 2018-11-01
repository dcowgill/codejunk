#!/usr/bin/env python3

def flatten(m, acc=None, prefix=""):
    acc = acc or {}
    for k, v in m.items():
        if isinstance(v, dict):
            flatten(v, acc, prefix+k+".")
        else:
            acc[prefix+k]=v
    return acc

def main():
    m = {"key": 3, "foo": {"a": 5, "bar": {"baz": 8}}}
    print(repr(flatten(m)))

if __name__ == '__main__':
    main()
