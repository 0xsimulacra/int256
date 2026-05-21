// Package int256: Fixed-size signed integer math library represented as a 256-bit absolute value plus a sign bit.
// Copyright (c) 2023 Trịnh Đức Bảo Linh(Kevin)
// Copyright (c) 2026 0xsimulacra
// SPDX-License-Identifier: MIT AND BSD-3-Clause
package int256

func (z *Int) Int64() int64 {
	absUint64 := z.abs.Uint64()
	if z.neg {
		return -int64(absUint64)
	}
	return int64(absUint64)
}
