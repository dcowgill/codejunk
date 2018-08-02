/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Given a dictionary of words and a string made up of those words (no spaces),
return the original sentence in a list. If there is more than one possible
reconstruction, return any of them. If there is no possible reconstruction, then
return null.

For example, given the set of words 'quick', 'brown', 'the', 'fox', and the
string "thequickbrownfox", you should return ['the', 'quick', 'brown', 'fox'].

Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond', and the string
"bedbathandbeyond", return either ['bed', 'bath', 'and', 'beyond] or ['bedbath',
'and', 'beyond'].

*/
package dcp022

// Strategy:
// Create n-ary Trie containing dictionary
// Search w/ backtracking

type Trie struct {
	children map[rune]*Trie
	terminal bool // does a word end here?
}

// insert adds the word to the trie. Note that t may be nil, in which case a new
// node will be returned; it is the caller's responsibility to save it.
func (t *Trie) insert(word []rune) *Trie {
	if t == nil {
		t = &Trie{}
	}
	if len(word) == 0 {
		t.terminal = true
		return t
	}
	if t.children == nil {
		t.children = make(map[rune]*Trie)
	}
	c := t.children[word[0]]
	if c == nil {
		c = &Trie{}
		t.children[word[0]] = c
	}
	t.children[word[0]].insert(word[1:])
	return t
}

// Searches for a set of words in the trie that can spell 'term'.
func search(root *Trie, term string) []string {
	return searchRec(root, root, nil, nil, []rune(term), 0)
}

// Recursive depth-first search.
func searchRec(root, curr *Trie, acc1 []string, acc2 []rune, term []rune, pos int) []string {
	if pos == len(term) {
		return append(acc1, string(acc2))
	}
	if curr.terminal {
		if result := searchRec(root, root, append(acc1, string(acc2)), nil, term, pos); result != nil {
			return result
		}
	}
	r := term[pos]
	if child := curr.children[r]; child != nil {
		if result := searchRec(root, child, acc1, append(acc2, r), term, pos+1); result != nil {
			return result
		}
	}
	return nil
}
