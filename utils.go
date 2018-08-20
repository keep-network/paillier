package paillier

import (
	"crypto/rand"
	"io"
	"math/big"
)

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var TWO = big.NewInt(2)
var FOUR = big.NewInt(4)

//  returns n! = n*(n-1)*(n-2)...3*2*1
func Factorial(n int) *big.Int {
	ret := big.NewInt(1)
	for i := 1; i <= n; i++ {
		ret = new(big.Int).Mul(ret, big.NewInt(int64(i)))
	}
	return ret
}

// Draws a non-zero, pseudorandom number from a group of integers modulo n.
//
// In modular arithmetic, the integers coprime to n from the set
// {0, 1, ..., n-1} form a group under multiplication modulo n called
// the multiplicative group if integers modulo n.
//
// Two numbers a and b are coprime (or relatively prime) if the only
// positive integer that divides both of them is 1.
func GetRandomNumberInMultiplicativeGroup(n *big.Int, random io.Reader) (*big.Int, error) {
	for {
		r, err := rand.Int(random, n)
		if err != nil {
			return nil, err
		}

		if ZERO.Cmp(r) != 0 && ONE.Cmp(new(big.Int).GCD(nil, nil, n, r)) == 0 {
			return r, nil
		}
	}
}

//  Return a random generator of RQn with high probability.  THIS METHOD
//  ONLY WORKS IF N IS THE PRODUCT OF TWO SAFE PRIMES! This heuristic is used
//  threshold signature paper in the Victor Shoup
func GetRandomGeneratorOfTheQuadraticResidue(n *big.Int, rand io.Reader) (*big.Int, error) {
	r, err := GetRandomNumberInMultiplicativeGroup(n, rand)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Mod(new(big.Int).Mul(r, r), n), nil
}
