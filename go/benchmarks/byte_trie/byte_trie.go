package byte_trie

type trie struct {
	nodes [9]*trie
}

func (t *trie) insert(v []byte) {
	if len(v) == 0 {
		return
	}
	k := key(v)
	u := t.nodes[k]
	if u == nil {
		u = &trie{}
		t.nodes[k] = u
	}
	u.insert(v[2:])
}

func (t *trie) contains(v []byte) bool {
	if len(v) == 0 {
		return true
	}
	if u := t.nodes[key(v)]; u != nil {
		return u.contains(v[2:])
	}
	return false
}

func key(v []byte) byte {
	return v[0] + 3*v[1]
}
