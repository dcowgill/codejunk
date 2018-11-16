/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Two Sigma.

Using a function rand5() that returns an integer from 1 to 5 (inclusive) with
uniform probability, implement a function rand7() that returns an integer from 1
to 7 (inclusive).

*/
package dcp045

import "math/rand"

func rand5() int { return 1 + rand.Intn(5) }
func rand7() int { return (rand5()+rand5()+rand5()+rand5()+rand5()+rand5()+rand5())%7 + 1 }
