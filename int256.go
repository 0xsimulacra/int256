// Package int256: Fixed-size signed integer math library represented as a 256-bit absolute value plus a sign bit.
// Copyright 2018-2020 uint256 Authors
// Copyright (c) 2023 Trịnh Đức Bảo Linh(Kevin)
// Copyright (c) 2026 0xsimulacra
// SPDX-License-Identifier: MIT AND BSD-3-Clause
package int256

import (
	"math/big"
	"math/bits"

	"github.com/holiman/uint256"
)

// Int
// Fixed-size signed integer math library represented as a 256-bit absolute value plus a sign bit.
//
// This is not Solidity/EVM int256 two's-complement arithmetic: bit 255 is not the sign bit,
// values are not interpreted as 256-bit EVM words, and the sign is stored separately in neg.
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
		z.neg = false
		return z
	}
	if n == 0 {
		z.abs.Set(&x.abs)
		z.neg = !z.abs.IsZero()
		return z
	}

	return z.rshNegative(x, n)
}

func (z *Int) rshNegative(x *Int, n uint) *Int {
	var rem bool
	switch {
	case n < 64:
		rem = x.abs[0]&((uint64(1)<<n)-1) != 0
		z.abs[0] = x.abs[0]>>n | x.abs[1]<<(64-n)
		z.abs[1] = x.abs[1]>>n | x.abs[2]<<(64-n)
		z.abs[2] = x.abs[2]>>n | x.abs[3]<<(64-n)
		z.abs[3] = x.abs[3] >> n
	case n == 64:
		rem = x.abs[0] != 0
		z.abs[0], z.abs[1], z.abs[2], z.abs[3] = x.abs[1], x.abs[2], x.abs[3], 0
	case n < 128:
		shift := n - 64
		rem = x.abs[0] != 0 || x.abs[1]&((uint64(1)<<shift)-1) != 0
		z.abs[0] = x.abs[1]>>shift | x.abs[2]<<(64-shift)
		z.abs[1] = x.abs[2]>>shift | x.abs[3]<<(64-shift)
		z.abs[2] = x.abs[3] >> shift
		z.abs[3] = 0
	case n == 128:
		rem = (x.abs[0] | x.abs[1]) != 0
		z.abs[0], z.abs[1], z.abs[2], z.abs[3] = x.abs[2], x.abs[3], 0, 0
	case n < 192:
		shift := n - 128
		rem = (x.abs[0]|x.abs[1]) != 0 || x.abs[2]&((uint64(1)<<shift)-1) != 0
		z.abs[0] = x.abs[2]>>shift | x.abs[3]<<(64-shift)
		z.abs[1] = x.abs[3] >> shift
		z.abs[2], z.abs[3] = 0, 0
	case n == 192:
		rem = (x.abs[0] | x.abs[1] | x.abs[2]) != 0
		z.abs[0], z.abs[1], z.abs[2], z.abs[3] = x.abs[3], 0, 0, 0
	case n < 256:
		shift := n - 192
		rem = (x.abs[0]|x.abs[1]|x.abs[2]) != 0 || x.abs[3]&((uint64(1)<<shift)-1) != 0
		z.abs[0] = x.abs[3] >> shift
		z.abs[1], z.abs[2], z.abs[3] = 0, 0, 0
	default:
		rem = (x.abs[0] | x.abs[1] | x.abs[2] | x.abs[3]) != 0
		z.abs[0], z.abs[1], z.abs[2], z.abs[3] = 0, 0, 0, 0
	}

	if rem {
		var carry uint64
		z.abs[0], carry = bits.Add64(z.abs[0], 1, 0)
		z.abs[1], carry = bits.Add64(z.abs[1], 0, carry)
		z.abs[2], carry = bits.Add64(z.abs[2], 0, carry)
		z.abs[3], _ = bits.Add64(z.abs[3], 0, carry)
	}
	z.neg = !z.abs.IsZero()
	return z
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
	return z.exp(x, y, m)
}

func (z *Int) exp(x, y, m *Int) *Int {
	if !y.neg {
		if m == nil || m.abs.IsZero() {
			z.abs.Exp(&x.abs, &y.abs)
			z.neg = !z.abs.IsZero() && x.neg && y.abs[0]&1 == 1
			return z
		}

		if x.abs.IsUint64() && m.abs.IsUint64() && m.abs[0] <= 1<<32-1 {
			mod := m.abs[0]
			base := x.abs[0] % mod
			if x.neg && y.abs[0]&1 == 1 && base != 0 {
				base = mod - base
			}
			result := uint64(1 % mod)

			expBitLen := y.abs.BitLen()
			curBit := 0
			for wordIndex := 0; wordIndex < 4 && curBit < expBitLen; wordIndex++ {
				word := y.abs[wordIndex]
				for ; curBit < expBitLen && curBit < (wordIndex+1)*64; curBit++ {
					if word&1 == 1 {
						result = (result * base) % mod
					}
					word >>= 1
					if curBit+1 < expBitLen {
						base = (base * base) % mod
					}
				}
			}

			z.abs.SetUint64(result)
			z.neg = false
			return z
		}

		var base, result uint256.Int
		base.Mod(&x.abs, &m.abs)
		if x.neg && y.abs[0]&1 == 1 && !base.IsZero() {
			base.Sub(&m.abs, &base)
		}

		result.SetUint64(1)
		result.Mod(&result, &m.abs)

		expBitLen := y.abs.BitLen()
		curBit := 0
		for wordIndex := 0; wordIndex < 4 && curBit < expBitLen; wordIndex++ {
			word := y.abs[wordIndex]
			for ; curBit < expBitLen && curBit < (wordIndex+1)*64; curBit++ {
				if word&1 == 1 {
					result.MulMod(&result, &base, &m.abs)
				}
				word >>= 1
				if curBit+1 < expBitLen {
					base.MulMod(&base, &base, &m.abs)
				}
			}
		}

		z.abs.Set(&result)
		z.neg = false
		return z
	}
	// TODO: implement
	var mBigInt *big.Int
	if m != nil {
		mBigInt = m.ToBig()
	}
	var iBig big.Int
	if iBig.Exp(x.ToBig(), y.ToBig(), mBigInt) == nil {
		return nil
	}
	result, _ := FromBig(&iBig)
	*z = *result
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
			x0, borrow := bits.Sub64(x.abs[0], 1, 0)
			x1, borrow := bits.Sub64(x.abs[1], 0, borrow)
			x2, borrow := bits.Sub64(x.abs[2], 0, borrow)
			x3, _ := bits.Sub64(x.abs[3], 0, borrow)

			y0, borrow := bits.Sub64(y.abs[0], 1, 0)
			y1, borrow := bits.Sub64(y.abs[1], 0, borrow)
			y2, borrow := bits.Sub64(y.abs[2], 0, borrow)
			y3, _ := bits.Sub64(y.abs[3], 0, borrow)

			z.abs[0], z.abs[1], z.abs[2], z.abs[3] = x0&y0, x1&y1, x2&y2, x3&y3
			z.abs[0], borrow = bits.Add64(z.abs[0], 1, 0)
			z.abs[1], borrow = bits.Add64(z.abs[1], 0, borrow)
			z.abs[2], borrow = bits.Add64(z.abs[2], 0, borrow)
			z.abs[3], _ = bits.Add64(z.abs[3], 0, borrow)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x | y == x | y
		z.abs[0] = x.abs[0] | y.abs[0]
		z.abs[1] = x.abs[1] | y.abs[1]
		z.abs[2] = x.abs[2] | y.abs[2]
		z.abs[3] = x.abs[3] | y.abs[3]
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // | is symmetric
	}

	// x | (-y) == x | ^(y-1) == ^((y-1) &^ x) == -(^((y-1) &^ x) + 1)
	y0, borrow := bits.Sub64(y.abs[0], 1, 0)
	y1, borrow := bits.Sub64(y.abs[1], 0, borrow)
	y2, borrow := bits.Sub64(y.abs[2], 0, borrow)
	y3, _ := bits.Sub64(y.abs[3], 0, borrow)

	z.abs[0], z.abs[1], z.abs[2], z.abs[3] = y0&^x.abs[0], y1&^x.abs[1], y2&^x.abs[2], y3&^x.abs[3]
	z.abs[0], borrow = bits.Add64(z.abs[0], 1, 0)
	z.abs[1], borrow = bits.Add64(z.abs[1], 0, borrow)
	z.abs[2], borrow = bits.Add64(z.abs[2], 0, borrow)
	z.abs[3], _ = bits.Add64(z.abs[3], 0, borrow)
	z.neg = true // z cannot be zero if one of x or y is negative

	return z
}

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) & (-y) == ^(x-1) & ^(y-1) == ^((x-1) | (y-1)) == -(((x-1) | (y-1)) + 1)
			x0, borrow := bits.Sub64(x.abs[0], 1, 0)
			x1, borrow := bits.Sub64(x.abs[1], 0, borrow)
			x2, borrow := bits.Sub64(x.abs[2], 0, borrow)
			x3, _ := bits.Sub64(x.abs[3], 0, borrow)

			y0, borrow := bits.Sub64(y.abs[0], 1, 0)
			y1, borrow := bits.Sub64(y.abs[1], 0, borrow)
			y2, borrow := bits.Sub64(y.abs[2], 0, borrow)
			y3, _ := bits.Sub64(y.abs[3], 0, borrow)

			z.abs[0], z.abs[1], z.abs[2], z.abs[3] = x0|y0, x1|y1, x2|y2, x3|y3
			if z.abs[0]&z.abs[1]&z.abs[2]&z.abs[3] == ^uint64(0) {
				panic("overflow")
			}
			z.abs[0], borrow = bits.Add64(z.abs[0], 1, 0)
			z.abs[1], borrow = bits.Add64(z.abs[1], 0, borrow)
			z.abs[2], borrow = bits.Add64(z.abs[2], 0, borrow)
			z.abs[3], _ = bits.Add64(z.abs[3], 0, borrow)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x & y == x & y
		z.abs[0] = x.abs[0] & y.abs[0]
		z.abs[1] = x.abs[1] & y.abs[1]
		z.abs[2] = x.abs[2] & y.abs[2]
		z.abs[3] = x.abs[3] & y.abs[3]
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // & is symmetric
	}

	// x & (-y) == x & ^(y-1) == x &^ (y-1)
	y0, borrow := bits.Sub64(y.abs[0], 1, 0)
	y1, borrow := bits.Sub64(y.abs[1], 0, borrow)
	y2, borrow := bits.Sub64(y.abs[2], 0, borrow)
	y3, _ := bits.Sub64(y.abs[3], 0, borrow)

	z.abs[0], z.abs[1], z.abs[2], z.abs[3] = x.abs[0]&^y0, x.abs[1]&^y1, x.abs[2]&^y2, x.abs[3]&^y3
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
