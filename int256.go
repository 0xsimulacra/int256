// Package int256: Fixed-size signed-magnitude math library with a 256-bit absolute value.
//
// Values are represented as a 256-bit absolute value plus a sign bit. This is
// not Solidity/EVM int256 two's-complement arithmetic.
//
// Copyright (c) 2023 Trịnh Đức Bảo Linh(Kevin)
// Copyright 2018-2020 uint256 Authors
// Copyright (c) 2026 0xsimulacra
// SPDX-License-Identifier: MIT AND BSD-3-Clause
package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var one = uint256.NewInt(1)
var maxUint256 = uint256.MustFromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

// Int
// Fixed-size signed-magnitude math library with a 256-bit absolute value.
//
// Values are represented as a 256-bit absolute value plus a sign bit. This is
// not Solidity/EVM int256 two's-complement arithmetic.
type Int struct {
	abs uint256.Int
	neg bool
}

// Clone method creates a deep copy of Int
func (z *Int) Clone() *Int {
	return &Int{abs: z.abs, neg: z.neg}
}

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
func (z *Int) Sign() int {
	if z.abs.IsZero() {
		return 0
	}
	if z.neg {
		return -1
	}
	return 1
}

func New() *Int {
	return &Int{}
}

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *Int {
	return New().SetInt64(x)
}

// SetInt64 sets z to x and returns z.
func (z *Int) SetInt64(x int64) *Int {
	if x >= 0 {
		z.neg = false
	} else {
		z.neg = true
		x = -x
	}
	z.abs.SetUint64(uint64(x))
	return z
}

// SetUint64 sets z to x and returns z.
func (z *Int) SetUint64(x uint64) *Int {
	z.abs.SetUint64(x)
	z.neg = false
	return z
}

// Set sets z to x and returns z
func (z *Int) Set(x *Int) *Int {
	z.abs.Set(&x.abs)
	z.neg = x.neg
	return z
}

func (z *Int) SetString(s string) (*Int, error) {
	origin := s
	neg := false
	// Remove max one leading +
	if len(s) > 0 && s[0] == '+' {
		neg = false
		s = s[1:]
	}

	if len(s) > 0 && s[0] == '-' {
		neg = true
		s = s[1:]
	}
	var (
		abs *uint256.Int
		err error
	)
	abs, err = uint256.FromDecimal(s)
	if err != nil {
		// TODO: parse base as input param
		b, ok := new(big.Int).SetString(origin, 16)
		if !ok {
			return nil, err
		}
		return MustFromBig(b), nil
	}

	return &Int{
		abs: *abs,
		neg: neg,
	}, nil
}

// // setFromScanner implements SetString given an io.ByteScanner.
// // For documentation see comments of SetString.
// func (z *Int) setFromScanner(r io.ByteScanner, base int) (*Int, bool) {
// 	if _, _, err := z.scan(r, base); err != nil {
// 		return nil, false
// 	}
// 	// entire content must have been consumed
// 	if _, err := r.ReadByte(); err != io.EOF {
// 		return nil, false
// 	}
// 	return z, true // err == io.EOF => scan consumed all content of r
// }

func (z *Int) Add(x, y *Int) *Int {
	neg := x.neg

	if x.neg == y.neg {
		// x + y == x + y
		// (-x) + (-y) == -(x + y)
		z.abs.Add(&x.abs, &y.abs)
	} else {
		// x + (-y) == x - y == -(y - x)
		// (-x) + y == y - x == -(x - y)
		if x.abs.Cmp(&y.abs) >= 0 {
			z.abs.Sub(&x.abs, &y.abs)
		} else {
			neg = !neg
			z.abs.Sub(&y.abs, &x.abs)
		}
	}
	z.neg = neg // 0 has no sign
	return z
}

// Sub sets z to the difference x-y and returns z.
func (z *Int) Sub(x, y *Int) *Int {
	neg := x.neg
	if x.neg != y.neg {
		// x - (-y) == x + y
		// (-x) - y == -(x + y)
		z.abs.Add(&x.abs, &y.abs)
	} else {
		// x - y == x - y == -(y - x)
		// (-x) - (-y) == y - x == -(x - y)
		if x.abs.Cmp(&y.abs) >= 0 {
			z.abs.Sub(&x.abs, &y.abs)
		} else {
			neg = !neg
			z.abs.Sub(&y.abs, &x.abs)
		}
	}
	z.neg = neg // 0 has no sign
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *Int) Mul(x, y *Int) *Int {
	z.abs.Mul(&x.abs, &y.abs)
	z.neg = x.neg != y.neg // 0 has no sign
	return z
}

// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
// It panics if x is negative.
func (z *Int) Sqrt(x *Int) *Int {
	if x.neg {
		panic("square root of negative number")
	}
	z.neg = false
	z.abs.Sqrt(&x.abs)
	return z
}

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int {
	if !x.neg {
		z.abs.Rsh(&x.abs, n)
		z.neg = x.neg
		return z
	}

	hasRem := hasLowerBits(&x.abs, n)
	z.abs.Rsh(&x.abs, n)
	if hasRem {
		z.abs.Add(&z.abs, one)
	}
	z.neg = !z.abs.IsZero()
	return z
}

func hasLowerBits(x *uint256.Int, n uint) bool {
	if n == 0 || x.IsZero() {
		return false
	}
	if n >= 256 {
		return true
	}

	words := n / 64
	for i := uint(0); i < words; i++ {
		if x[i] != 0 {
			return true
		}
	}

	bits := n % 64
	if bits == 0 {
		return false
	}

	mask := (uint64(1) << bits) - 1

	return x[words]&mask != 0
}

// Quo sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func (z *Int) Quo(x, y *Int) *Int {
	z.abs.Div(&x.abs, &y.abs)
	z.neg = !z.abs.IsZero() && x.neg != y.neg // 0 has no sign
	return z
}

// Rem sets z to the remainder x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func (z *Int) Rem(x, y *Int) *Int {
	z.abs.Mod(&x.abs, &y.abs)
	z.neg = !z.abs.IsZero() && x.neg // 0 has no sign
	return z
}

// Eq returns true if z == x
func (z *Int) Eq(x *Int) bool {
	if z.abs.IsZero() && x.abs.IsZero() {
		return true
	}
	return z.neg == x.neg && z.abs.Eq(&x.abs)
}

// IsZero returns true if z == 0
func (z *Int) IsZero() bool {
	return z.abs.IsZero()
}

// Cmp compares x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func (z *Int) Cmp(x *Int) (r int) {
	// x cmp y == x cmp y
	// x cmp (-y) == x
	// (-x) cmp y == y
	// (-x) cmp (-y) == -(x cmp y)
	switch {
	case z == x:
		// nothing to do
	case z.neg == x.neg:
		r = z.abs.Cmp(&x.abs)
		if z.neg {
			r = -r
		}
	case z.abs.IsZero() && x.abs.IsZero():
		r = 0
	case z.neg:
		r = -1
	default:
		r = 1
	}
	return
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m != 0, y < 0,
// and x and m are not relatively prime, z is unchanged and nil is returned.
//
// Modular exponentiation of inputs of a particular size is not a
// cryptographically constant-time operation.
func (z *Int) Exp(x, y, m *Int) *Int {
	if x == nil {
		panic("x is nil")
	}
	if !x.neg && !y.neg && m == nil {
		z.neg = false
		z.abs.Exp(&x.abs, &y.abs)
		return z
	}
	// TODO: implement
	var mBigInt *big.Int
	if m != nil {
		mBigInt = m.ToBig()
	}
	big := new(big.Int).Exp(x.ToBig(), y.ToBig(), mBigInt)
	z, _ = FromBig(big)
	return z
}

// MulDivOverflow calculates (x*y)/d with full precision, returns z and whether overflow occurred in multiply process (result does not fit to 256-bit).
// computes 512-bit multiplication and 512 by 256 division.
func (z *Int) MulDivOverflow(x, y, d *Int) (*Int, bool) {
	z.neg = (x.neg != y.neg) != d.neg

	var overflow bool
	_, overflow = z.abs.MulDivOverflow(&x.abs, &y.abs, &d.abs)

	return z, overflow
}

func (z *Int) Div(x, y *Int) *Int {
	z.abs.Div(&x.abs, &y.abs)
	if x.neg == y.neg {
		z.neg = false
	} else {
		z.neg = true
	}
	return z
}

// Lsh sets z = x << n and returns z.
func (z *Int) Lsh(x *Int, n uint) *Int {
	if bitLen := x.abs.BitLen(); bitLen != 0 && n > uint(256-bitLen) {
		panic("overflow")
	}
	z.abs.Lsh(&x.abs, n)
	z.neg = x.neg
	return z
}

// Or sets z = x | y and returns z.
func (z *Int) Or(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) | (-y) == ^(x-1) | ^(y-1) == ^((x-1) & (y-1)) == -(((x-1) & (y-1)) + 1)
			x1 := new(uint256.Int).Sub(&x.abs, one)
			y1 := new(uint256.Int).Sub(&y.abs, one)
			z.abs.Add(z.abs.And(x1, y1), one)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x | y == x | y
		z.abs.Or(&x.abs, &y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // | is symmetric
	}

	// x | (-y) == x | ^(y-1) == ^((y-1) &^ x) == -(^((y-1) &^ x) + 1)
	y1 := new(uint256.Int).Sub(&y.abs, one)
	z.abs.Add(z.abs.And(y1, new(uint256.Int).Xor(&x.abs, maxUint256)), one)
	z.neg = true // z cannot be zero if one of x or y is negative

	return z
}

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) & (-y) == ^(x-1) & ^(y-1) == ^((x-1) | (y-1)) == -(((x-1) | (y-1)) + 1)
			x1 := new(uint256.Int).Sub(&x.abs, one)
			y1 := new(uint256.Int).Sub(&y.abs, one)
			z.abs.Add(z.abs.Or(x1, y1), one)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x & y == x & y
		z.abs.And(&x.abs, &y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // & is symmetric
	}

	// x & (-y) == x & ^(y-1) == x &^ (y-1)
	y1 := new(uint256.Int).Sub(&y.abs, one)
	z.abs.And(&x.abs, new(uint256.Int).Xor(y1, maxUint256))
	z.neg = false

	return z
}

// MostSignificantBit return the most significant bit of z, ignoring the bit or the sign
func (z *Int) MostSignificantBit() uint8 {
	return uint8(z.abs.BitLen() - 1)
}

// Negate return a new int256.Int equal to the negative of z
func (z *Int) Negate() *Int {
	x := z.Clone()
	x.neg = !z.neg
	return x
}

// InPlaceNegate  transform the sign of z to its opposite
func (z *Int) InPlaceNegate() *Int {
	z.neg = !z.neg
	return z
}

// AbsInt Absolute value of z having the same type int256.Int
func (z *Int) AbsInt() *Int {
	return &Int{
		abs: z.abs,
		neg: false,
	}
}

// SignedMaxAbs take the maximum of abs of a and b and return z with that value but the sign of c
func (z *Int) SignedMaxAbs(a, b, c *Int) *Int {
	if a.abs.Cmp(&b.abs) >= 0 {
		// a is bigger or equal
		z.abs = a.abs
	} else {
		// b is bigger
		z.abs = b.abs
	}
	z.neg = c.neg
	return z
}

// Signed z with the value equal to a but the sign of b
func (z *Int) Signed(a, b *Int) *Int {
	z.abs = a.abs
	z.neg = b.neg
	return z
}

// Relu function (x > 0 ? x : 0)
func (z *Int) Relu() *Int {
	x := NewInt(0)
	if !z.neg {
		x.abs = z.abs
	}
	return x
}

// CondRef return either x or y based on the condition c (c ? x : y)
func CondRef(c bool, x, y *Int) *Int {
	if c {
		return x
	}
	return y
}
