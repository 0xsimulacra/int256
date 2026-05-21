// Package int256: Fixed-size signed integer math library represented as a 256-bit absolute value plus a sign bit.
// Copyright (c) 2023 Trịnh Đức Bảo Linh(Kevin)
// Copyright 2018-2020 uint256 Authors
// Copyright (c) 2026 0xsimulacra
// SPDX-License-Identifier: MIT AND BSD-3-Clause
package int256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var zero = big.NewInt(0)

func (z *Int) ToBig() *big.Int {
	b := z.abs.ToBig()
	if z.neg {
		return b.Neg(b)
	}
	return b
}

func MustFromBig(x *big.Int) *Int {
	iBig, overflow := FromBig(x)
	if overflow {
		panic("cannot parsing from big.Int")
	}
	return iBig
}

func FromBig(x *big.Int) (*Int, bool) {
	num := x
	neg := false
	if x.Cmp(zero) == -1 {
		num = new(big.Int).Neg(x)
		neg = true
	}
	abs, overflow := uint256.FromBig(num)
	/*
		type Int [4]uint64
		Currently, uint26 has maxWord is 4


		bigInt has len(words) that can great than 4. So we can receive overflow error.

		words := b.Bits()
		overflow := len(words) > maxWords
		ref from uint256 code
		https://github.com/holiman/uint256/blob/master/conversion.go#L202
	*/
	if overflow {
		abs, err := uint256.FromDecimal(x.String())
		if err != nil {
			return nil, overflow
		}
		neg = x.Sign() < 0
		return &Int{
			abs: *abs,
			neg: neg,
		}, true
	}
	return &Int{
		abs: *abs,
		neg: neg,
	}, false
}
