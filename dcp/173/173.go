/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Stripe.

Write a function to flatten a nested dictionary. Namespace the keys with a period.

For example, given the following dictionary:

{
    "key": 3,
    "foo": {
        "a": 5,
        "bar": {
            "baz": 8
        }
    }
}
it should become:

{
    "key": 3,
    "foo.a": 5,
    "foo.bar.baz": 8
}
You can assume keys do not contain dots in them, i.e. no clobbering will occur.

*/
package dcp173

import "fmt"

type M map[string]interface{}

func flatten(m M) map[string]int {
	acc := make(map[string]int)
	var helper func(prefix string, m M)
	helper = func(prefix string, m M) {
		for key, val := range m {
			switch val := val.(type) {
			case M:
				helper(prefix+key+".", val)
			case int:
				acc[prefix+key] = val
			default:
				panic(fmt.Sprintf("flatten: unexpected value type: %T", val))
			}
		}
	}
	helper("", m)
	return acc
}
