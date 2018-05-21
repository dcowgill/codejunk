/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Twitter.

Implement an autocomplete system. That is, given a query string s and a set of
all possible query strings, return all strings in the set that have s as a
prefix.

For example, given the query string de and the set of strings [dog, deer, deal],
return [deer, deal].

Hint: Try preprocessing the dictionary into a more efficient data structure to
speed up queries.

*/
package dcp011

// A trie of runes.
type node struct {
	children map[rune]*node
	terminal bool
}

// Adds a rune sequence to the trie rooted at n.
func (n *node) insert(a []rune) {
	if len(a) == 0 {
		n.terminal = true
		return
	}
	if n.children == nil {
		n.children = make(map[rune]*node)
	}
	r := a[0]
	c := n.children[r]
	if c == nil {
		c = &node{}
		n.children[r] = c
	}
	c.insert(a[1:])
}

// Creates a trie containing the strings in dict.
func buildTrie(dict []string) *node {
	root := &node{}
	for _, s := range dict {
		root.insert([]rune(s))
	}
	return root
}

// Returns every string in the trie that begins with the prefix.
func autocomplete(root *node, prefix string) []string {
	if len(prefix) == 0 {
		return nil // special case
	}
	return find(root, prefix, []rune(prefix))
}

// Recursive implementation of autocomplete().
func find(n *node, prefix string, rest []rune) []string {
	if len(rest) == 0 {
		var dst []string
		gather(n, prefix, nil, &dst)
		return dst
	}
	if c, ok := n.children[rest[0]]; ok {
		return find(c, prefix, rest[1:])
	}
	return nil
}

// Prepends the prefix to every string in the trie, storing the results in *dst.
func gather(n *node, prefix string, acc []rune, dst *[]string) {
	if n.terminal {
		*dst = append(*dst, prefix+string(acc))
	}
	for r, child := range n.children {
		gather(child, prefix, append(acc, r), dst)
	}
}
