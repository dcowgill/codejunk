/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

An XOR linked list is a more memory efficient doubly linked list. Instead of
each node holding next and prev fields, it holds a field named both, which is an
XOR of the next node and the previous node. Implement an XOR linked list; it has
an add(element) which adds the element to the end, and a get(index) which
returns the node at index.

If using a language that has no pointers (such as Python), you can assume you
have access to get_pointer and dereference_pointer functions that converts
between nodes and memory addresses.

*/

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

struct node {
    int value;
    uintptr_t both;
};

typedef struct node node;

/* Returns the node after curr. prev must be the node before curr. */
node* next_node(node* curr, node* prev) {
    return (node*) (curr->both ^ ((uintptr_t) prev));
}

/* Sets curr->both using prev and next. */
void set_links(node *curr, node* prev, node* next) {
    curr->both = ((uintptr_t) next) ^ ((uintptr_t) prev);
}

/* Recursive implementation of add(). */
node* add_rec(node* curr, node* prev, int value) {
    if (curr == 0) {
        curr = (node*) malloc(sizeof(node));
        curr->value = value;
        set_links(curr, 0, 0);
        return curr;
    }
    node *next = next_node(curr, prev);
    if (next != 0) {
        return add_rec(next, curr, value);
    }
    next = (node*) malloc(sizeof(node));
    next->value = value;
    set_links(next, curr, 0);
    set_links(curr, prev, next);
    return next;
}

/* Appends a node to the list with the given value. */
node* add(node* head, int value) {
    return add_rec(head, 0, value);
}

/* Recursive implementation of get(). */
node* get_rec(node* curr, node* prev, int i) {
    if (i == 0 || curr == 0) {
        return curr;
    }
    return get_rec(next_node(curr, prev), curr, i-1);
}

/* Returns the ith node in the list, or null if it does not exist. */
node* get(node* head, int i) {
    assert(i >= 0);
    return get_rec(head, 0, i);
}

/* Calls fn for each node in the list, starting with curr. */
void walk(node* curr, node* prev, void (*fn)(node*)) {
    if (curr != 0) {
        fn(curr);
        walk(next_node(curr, prev), curr, fn);
    }
}

/* Prints the node to stdout. For debugging. */
void print_node(node *n) {
    printf("node->value = %5d, node->both = %lu\n", n->value, n->both);
}

int main() {
    node *head = 0;
    node *p;
    int num_nodes = 10;
    int i;

    /* Append several nodes to the list. */
    head = add(0, 0);
    for (i = 1; i < num_nodes; i++) {
        add(head, i*i);
    }

    /* Print the list contents. */
    walk(head, 0, print_node);

    /* Verify that the nodes can be retrieved in order. */
    for (i = 0; i < num_nodes; i++) {
        p = get(head, i);
        assert(p != 0);
        assert(p->value == i*i);
    }
    assert(get(head, num_nodes) == 0);

    return 0;
}
