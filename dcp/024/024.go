/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Implement locking in a binary tree. A binary tree node can be locked or unlocked
only if all of its descendants or ancestors are not locked.

Design a binary tree node class with the following methods:

is_locked, which returns whether the node is locked

lock, which attempts to lock the node. If it cannot be locked, then it should
return false. Otherwise, it should lock it and return true.

unlock, which unlocks the node. If it cannot be unlocked, then it should return
false. Otherwise, it should unlock it and return true.

You may augment the node to add parent pointers or any other property you would
like. You may assume the class is used in a single-threaded program, so there is
no need for actual locks or mutexes. Each method should run in O(h), where h is
the height of the tree.

*/
package dcp024

//
// Node is a binary tree that can be locked.
//
// The problem requirements aren't entirely clear to me; "a binary tree node can
// be locked or unlocked only if all of its descendants or ancestors are not
// locked" could be interpreted to mean either of the following:
//
// (1) *either* all ancestors *or* all descendants must not be locked; or
// (2) locking *any* node in the tree effectively locks the *entire* tree.
//
// In the latter case, it is sufficient to lock the tree's root node, which
// means all operations are O(height). Let's do that.
//
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	locked bool
}

// IsLocked reports whether the tree is locked.
func (n *Node) IsLocked() bool {
	if n.parent != nil {
		return n.parent.IsLocked()
	}
	return n.locked
}

// Lock tries to lock the entire tree. Returns true if successful.
func (n *Node) Lock() bool {
	if n.parent != nil {
		return n.parent.Lock()
	}
	if n.locked {
		return false
	}
	n.locked = true
	return true
}

// Unlock tries to unlock the entire tree. Returns true if successful.
func (n *Node) Unlock() bool {
	if n.parent != nil {
		return n.parent.Unlock()
	}
	if !n.locked {
		return false
	}
	n.locked = false
	return true
}
