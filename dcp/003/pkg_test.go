package dcp003

import (
	"encoding/json"
	"reflect"
	"testing"
)

var serializers = []struct {
	name string
	s    func(*Node) string
	d    func(string) *Node
}{
	{"json", serializeJSON, deserializeJSON},
	{"flat", serialize, deserialize},
}

func TestAll(t *testing.T) {
	for _, serializer := range serializers {
		t.Run(serializer.name, func(t *testing.T) {
			{
				root := &Node{
					Val:   "root",
					Left:  &Node{Left: &Node{Val: "left.left"}},
					Right: &Node{Val: "right", Right: &Node{Val: "tab\tnewline\ns p a c e s\n"}},
				}
				result := serializer.d(serializer.s(root))
				if !reflect.DeepEqual(root, result) {
					t.Fatalf("got %s, want %s", dumpTree(result), dumpTree(root))
				}
			}
		})
	}
}

func dumpTree(n *Node) string {
	bs, _ := json.Marshal(n)
	return string(bs)
}
