// Package int256: Fixed-size signed-magnitude math library with a 256-bit absolute value.
// Copyright 2018-2020 uint256 Authors
// Copyright (c) 2026 0xsimulacra
// SPDX-License-Identifier: BSD-3-Clause
package int256

import (
	"math/big"
	"testing"
)

var (
	benchmarkIntSink    *Int
	benchmarkBigSink    *big.Int
	benchmarkBoolSink   bool
	benchmarkIntSinkVal int
	benchmarkStringSink string
	benchmarkBytesSink  []byte
	benchmarkUint8Sink  uint8
	benchmarkOverflow   bool
)

type benchmarkBinaryCase struct {
	name string
	xs   *[numSamples]Int
	ys   *[numSamples]Int
	bxs  *[numSamples]big.Int
	bys  *[numSamples]big.Int
}

var benchmarkBinaryCases = []benchmarkBinaryCase{
	{
		name: "positive_positive",
		xs:   &int256Samples,
		ys:   &int256SamplesLt,
		bxs:  &big256Samples,
		bys:  &big256SamplesLt,
	},
	{
		name: "positive_negative",
		xs:   &int256Samples,
		ys:   &int256SamplesNeg,
		bxs:  &big256Samples,
		bys:  &big256SamplesNeg,
	},
	{
		name: "negative_positive",
		xs:   &int256SamplesNeg,
		ys:   &int256Samples,
		bxs:  &big256SamplesNeg,
		bys:  &big256Samples,
	},
	{
		name: "negative_negative",
		xs:   &int256SamplesNeg,
		ys:   &int256SamplesLtNeg,
		bxs:  &big256SamplesNeg,
		bys:  &big256SamplesLtNeg,
	},
}

var benchmarkDivCases = []benchmarkBinaryCase{
	{
		name: "positive_positive",
		xs:   &int256Samples,
		ys:   &int256SamplesLt,
		bxs:  &big256Samples,
		bys:  &big256SamplesLt,
	},
	{
		name: "positive_negative",
		xs:   &int256Samples,
		ys:   &int256SamplesNeg,
		bxs:  &big256Samples,
		bys:  &big256SamplesNeg,
	},
	{
		name: "negative_positive",
		xs:   &int256SamplesNeg,
		ys:   &int256Samples,
		bxs:  &big256SamplesNeg,
		bys:  &big256Samples,
	},
	{
		name: "negative_negative",
		xs:   &int256SamplesNeg,
		ys:   &int256Samples,
		bxs:  &big256SamplesNeg,
		bys:  &big256Samples,
	},
}

var (
	benchmarkSmallPositive [numSamples]Int
	benchmarkSmallNegative [numSamples]Int
	benchmarkSmallBigPos   [numSamples]big.Int
	benchmarkSmallBigNeg   [numSamples]big.Int
	_                      = initBenchmarkSamples()
)

func initBenchmarkSamples() bool {
	for i := 0; i < numSamples; i++ {
		value := int64((i + 1) * 7919)
		benchmarkSmallPositive[i].SetInt64(value)
		benchmarkSmallNegative[i].SetInt64(-value)
		benchmarkSmallBigPos[i].SetInt64(value)
		benchmarkSmallBigNeg[i].SetInt64(-value)
	}
	return true
}

func benchmarkIntBinary(
	b *testing.B,
	xs, ys *[numSamples]Int,
	op func(z, x, y *Int) *Int,
) {
	b.ReportAllocs()
	z := New()
	b.ResetTimer()
	for n := 0; n < b.N; {
		for i := 0; i < numSamples && n < b.N; i++ {
			op(z, &xs[i], &ys[i])
			n++
		}
	}
	benchmarkIntSink = z
}

func benchmarkBigBinary(
	b *testing.B,
	xs, ys *[numSamples]big.Int,
	op func(z, x, y *big.Int) *big.Int,
) {
	b.ReportAllocs()
	z := new(big.Int)
	b.ResetTimer()
	for n := 0; n < b.N; {
		for i := 0; i < numSamples && n < b.N; i++ {
			op(z, &xs[i], &ys[i])
			n++
		}
	}
	benchmarkBigSink = z
}

func benchmarkIntUnary(
	b *testing.B,
	xs *[numSamples]Int,
	op func(z, x *Int) *Int,
) {
	b.ReportAllocs()
	z := New()
	b.ResetTimer()
	for n := 0; n < b.N; {
		for i := 0; i < numSamples && n < b.N; i++ {
			op(z, &xs[i])
			n++
		}
	}
	benchmarkIntSink = z
}

func benchmarkBigUnary(
	b *testing.B,
	xs *[numSamples]big.Int,
	op func(z, x *big.Int) *big.Int,
) {
	b.ReportAllocs()
	z := new(big.Int)
	b.ResetTimer()
	for n := 0; n < b.N; {
		for i := 0; i < numSamples && n < b.N; i++ {
			op(z, &xs[i])
			n++
		}
	}
	benchmarkBigSink = z
}

func BenchmarkAdd(b *testing.B) {
	for _, tc := range benchmarkBinaryCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Add)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Add)
		})
	}
}

func BenchmarkSub(b *testing.B) {
	for _, tc := range benchmarkBinaryCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Sub)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Sub)
		})
	}
}

func BenchmarkMul(b *testing.B) {
	for _, tc := range benchmarkBinaryCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Mul)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Mul)
		})
	}
}

func BenchmarkQuo(b *testing.B) {
	for _, tc := range benchmarkDivCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Quo)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Quo)
		})
	}
}

func BenchmarkDiv(b *testing.B) {
	for _, tc := range benchmarkDivCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Div)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Div)
		})
	}
}

func BenchmarkRem(b *testing.B) {
	for _, tc := range benchmarkDivCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Rem)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Rem)
		})
	}
}

func BenchmarkMulDivOverflow(b *testing.B) {
	for _, tc := range benchmarkDivCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			b.ReportAllocs()
			z := New()
			var overflow bool
			b.ResetTimer()
			for n := 0; n < b.N; {
				for i := 0; i < numSamples && n < b.N; i++ {
					_, overflow = z.MulDivOverflow(&tc.xs[i], &tc.xs[i], &tc.ys[i])
					n++
				}
			}
			benchmarkIntSink = z
			benchmarkOverflow = overflow
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			b.ReportAllocs()
			z := new(big.Int)
			b.ResetTimer()
			for n := 0; n < b.N; {
				for i := 0; i < numSamples && n < b.N; i++ {
					z.Mul(&tc.bxs[i], &tc.bxs[i])
					z.Quo(z, &tc.bys[i])
					n++
				}
			}
			benchmarkBigSink = z
		})
	}
}

func BenchmarkSqrt(b *testing.B) {
	b.Run("positive/int256", func(b *testing.B) {
		benchmarkIntUnary(b, &int256Samples, (*Int).Sqrt)
	})
	b.Run("positive/big", func(b *testing.B) {
		benchmarkBigUnary(b, &big256Samples, (*big.Int).Sqrt)
	})
}

func BenchmarkAnd(b *testing.B) {
	for _, tc := range benchmarkBinaryCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).And)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).And)
		})
	}
}

func BenchmarkOr(b *testing.B) {
	for _, tc := range benchmarkBinaryCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			benchmarkIntBinary(b, tc.xs, tc.ys, (*Int).Or)
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			benchmarkBigBinary(b, tc.bxs, tc.bys, (*big.Int).Or)
		})
	}
}

func BenchmarkCmp(b *testing.B) {
	for _, tc := range benchmarkBinaryCases {
		b.Run(tc.name+"/int256", func(b *testing.B) {
			b.ReportAllocs()
			var sink int
			b.ResetTimer()
			for n := 0; n < b.N; {
				for i := 0; i < numSamples && n < b.N; i++ {
					sink = tc.xs[i].Cmp(&tc.ys[i])
					n++
				}
			}
			benchmarkIntSinkVal = sink
		})
		b.Run(tc.name+"/big", func(b *testing.B) {
			b.ReportAllocs()
			var sink int
			b.ResetTimer()
			for n := 0; n < b.N; {
				for i := 0; i < numSamples && n < b.N; i++ {
					sink = tc.bxs[i].Cmp(&tc.bys[i])
					n++
				}
			}
			benchmarkIntSinkVal = sink
		})
	}
}

func BenchmarkEq(b *testing.B) {
	b.Run("equal/int256", func(b *testing.B) {
		b.ReportAllocs()
		var sink bool
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				sink = int256Samples[i].Eq(&int256Samples[i])
				n++
			}
		}
		benchmarkBoolSink = sink
	})
	b.Run("not_equal/int256", func(b *testing.B) {
		b.ReportAllocs()
		var sink bool
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				sink = int256Samples[i].Eq(&int256SamplesLt[i])
				n++
			}
		}
		benchmarkBoolSink = sink
	})
	b.Run("equal/big", func(b *testing.B) {
		b.ReportAllocs()
		var sink bool
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				sink = big256Samples[i].Cmp(&big256Samples[i]) == 0
				n++
			}
		}
		benchmarkBoolSink = sink
	})
	b.Run("not_equal/big", func(b *testing.B) {
		b.ReportAllocs()
		var sink bool
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				sink = big256Samples[i].Cmp(&big256SamplesLt[i]) == 0
				n++
			}
		}
		benchmarkBoolSink = sink
	})
}

type benchmarkShiftCase struct {
	name string
	n    uint
}

var benchmarkShiftCases = []benchmarkShiftCase{
	{name: "n_eq_0", n: 0},
	{name: "n_gt_0", n: 1},
	{name: "n_gt_64", n: 65},
	{name: "n_gt_128", n: 129},
	{name: "n_gt_192", n: 193},
}

func benchmarkIntShift(
	b *testing.B,
	samples *[numSamples]Int,
	n uint,
	op func(z, x *Int, n uint) *Int,
) {
	b.ReportAllocs()
	z := New()
	b.ResetTimer()
	for j := 0; j < b.N; {
		for i := 0; i < numSamples && j < b.N; i++ {
			op(z, &samples[i], n)
			j++
		}
	}
	benchmarkIntSink = z
}

func benchmarkBigShift(
	b *testing.B,
	samples *[numSamples]big.Int,
	n uint,
	op func(z, x *big.Int, n uint) *big.Int,
) {
	b.ReportAllocs()
	z := new(big.Int)
	b.ResetTimer()
	for j := 0; j < b.N; {
		for i := 0; i < numSamples && j < b.N; i++ {
			op(z, &samples[i], n)
			j++
		}
	}
	benchmarkBigSink = z
}

func BenchmarkLsh(b *testing.B) {
	for _, shift := range benchmarkShiftCases {
		b.Run(shift.name+"/positive/int256", func(b *testing.B) {
			benchmarkIntShift(b, &benchmarkSmallPositive, shift.n, (*Int).Lsh)
		})
		b.Run(shift.name+"/positive/big", func(b *testing.B) {
			benchmarkBigShift(b, &benchmarkSmallBigPos, shift.n, (*big.Int).Lsh)
		})
		b.Run(shift.name+"/negative/int256", func(b *testing.B) {
			benchmarkIntShift(b, &benchmarkSmallNegative, shift.n, (*Int).Lsh)
		})
		b.Run(shift.name+"/negative/big", func(b *testing.B) {
			benchmarkBigShift(b, &benchmarkSmallBigNeg, shift.n, (*big.Int).Lsh)
		})
	}
}

func BenchmarkRsh(b *testing.B) {
	for _, shift := range benchmarkShiftCases {
		b.Run(shift.name+"/positive/int256", func(b *testing.B) {
			benchmarkIntShift(b, &int256Samples, shift.n, (*Int).Rsh)
		})
		b.Run(shift.name+"/positive/big", func(b *testing.B) {
			benchmarkBigShift(b, &big256Samples, shift.n, (*big.Int).Rsh)
		})
		b.Run(shift.name+"/negative/int256", func(b *testing.B) {
			benchmarkIntShift(b, &int256SamplesNeg, shift.n, (*Int).Rsh)
		})
		b.Run(shift.name+"/negative/big", func(b *testing.B) {
			benchmarkBigShift(b, &big256SamplesNeg, shift.n, (*big.Int).Rsh)
		})
	}
}

func BenchmarkExp(b *testing.B) {
	base := NewInt(7)
	negativeBase := NewInt(-7)
	exponent := NewInt(17)
	modulus := NewInt(101)
	bigBase := big.NewInt(7)
	bigNegativeBase := big.NewInt(-7)
	bigExponent := big.NewInt(17)
	bigModulus := big.NewInt(101)

	b.Run("positive_no_mod/int256", func(b *testing.B) {
		b.ReportAllocs()
		z := New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z.Exp(base, exponent, nil)
		}
		benchmarkIntSink = z
	})
	b.Run("positive_no_mod/big", func(b *testing.B) {
		b.ReportAllocs()
		z := new(big.Int)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z.Exp(bigBase, bigExponent, nil)
		}
		benchmarkBigSink = z
	})
	b.Run("positive_mod/int256", func(b *testing.B) {
		b.ReportAllocs()
		z := New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z.Exp(base, exponent, modulus)
		}
		benchmarkIntSink = z
	})
	b.Run("positive_mod/big", func(b *testing.B) {
		b.ReportAllocs()
		z := new(big.Int)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z.Exp(bigBase, bigExponent, bigModulus)
		}
		benchmarkBigSink = z
	})
	b.Run("negative_base_mod/int256", func(b *testing.B) {
		b.ReportAllocs()
		z := New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z.Exp(negativeBase, exponent, modulus)
		}
		benchmarkIntSink = z
	})
	b.Run("negative_base_mod/big", func(b *testing.B) {
		b.ReportAllocs()
		z := new(big.Int)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z.Exp(bigNegativeBase, bigExponent, bigModulus)
		}
		benchmarkBigSink = z
	})
}

func BenchmarkSet(b *testing.B) {
	b.Run("positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		z := New()
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z.Set(&int256Samples[i])
				n++
			}
		}
		benchmarkIntSink = z
	})
	b.Run("negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		z := New()
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z.Set(&int256SamplesNeg[i])
				n++
			}
		}
		benchmarkIntSink = z
	})
	b.Run("positive/big", func(b *testing.B) {
		b.ReportAllocs()
		z := new(big.Int)
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z.Set(&big256Samples[i])
				n++
			}
		}
		benchmarkBigSink = z
	})
	b.Run("negative/big", func(b *testing.B) {
		b.ReportAllocs()
		z := new(big.Int)
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z.Set(&big256SamplesNeg[i])
				n++
			}
		}
		benchmarkBigSink = z
	})
}

func BenchmarkClone(b *testing.B) {
	b.Run("positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *Int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256Samples[i].Clone()
				n++
			}
		}
		benchmarkIntSink = z
	})
	b.Run("negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *Int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256SamplesNeg[i].Clone()
				n++
			}
		}
		benchmarkIntSink = z
	})
	b.Run("positive/big", func(b *testing.B) {
		b.ReportAllocs()
		var z *big.Int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = new(big.Int).Set(&big256Samples[i])
				n++
			}
		}
		benchmarkBigSink = z
	})
	b.Run("negative/big", func(b *testing.B) {
		b.ReportAllocs()
		var z *big.Int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = new(big.Int).Set(&big256SamplesNeg[i])
				n++
			}
		}
		benchmarkBigSink = z
	})
}

func BenchmarkToBig(b *testing.B) {
	b.Run("positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *big.Int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256Samples[i].ToBig()
				n++
			}
		}
		benchmarkBigSink = z
	})
	b.Run("negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *big.Int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256SamplesNeg[i].ToBig()
				n++
			}
		}
		benchmarkBigSink = z
	})
}

func BenchmarkEncoding(b *testing.B) {
	b.Run("hex_positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z string
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256Samples[i].Hex()
				n++
			}
		}
		benchmarkStringSink = z
	})
	b.Run("hex_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z string
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256SamplesNeg[i].Hex()
				n++
			}
		}
		benchmarkStringSink = z
	})
	b.Run("dec_positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z string
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256Samples[i].Dec()
				n++
			}
		}
		benchmarkStringSink = z
	})
	b.Run("dec_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z string
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256SamplesNeg[i].Dec()
				n++
			}
		}
		benchmarkStringSink = z
	})
	b.Run("marshal_text_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z []byte
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z, _ = int256SamplesNeg[i].MarshalText()
				n++
			}
		}
		benchmarkBytesSink = z
	})
	b.Run("marshal_json_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z []byte
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z, _ = int256SamplesNeg[i].MarshalJSON()
				n++
			}
		}
		benchmarkBytesSink = z
	})
}

func BenchmarkParsing(b *testing.B) {
	positiveDecimal := int256Samples[0].Dec()
	negativeDecimal := int256SamplesNeg[0].Dec()
	positiveHex := int256Samples[0].Hex()
	negativeHex := int256SamplesNeg[0].Hex()

	b.Run("from_decimal_positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *Int
		var err error
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z, err = FromDecimal(positiveDecimal)
			if err != nil {
				b.Fatal(err)
			}
		}
		benchmarkIntSink = z
	})
	b.Run("from_decimal_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *Int
		var err error
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z, err = FromDecimal(negativeDecimal)
			if err != nil {
				b.Fatal(err)
			}
		}
		benchmarkIntSink = z
	})
	b.Run("from_hex_positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *Int
		var err error
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z, err = FromHex(positiveHex)
			if err != nil {
				b.Fatal(err)
			}
		}
		benchmarkIntSink = z
	})
	b.Run("from_hex_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z *Int
		var err error
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			z, err = FromHex(negativeHex)
			if err != nil {
				b.Fatal(err)
			}
		}
		benchmarkIntSink = z
	})
}

func BenchmarkSmallAccessors(b *testing.B) {
	b.Run("sign_positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256Samples[i].Sign()
				n++
			}
		}
		benchmarkIntSinkVal = z
	})
	b.Run("sign_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256SamplesNeg[i].Sign()
				n++
			}
		}
		benchmarkIntSinkVal = z
	})
	b.Run("int64_positive/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int(benchmarkSmallPositive[i].Int64())
				n++
			}
		}
		benchmarkIntSinkVal = z
	})
	b.Run("int64_negative/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z int
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int(benchmarkSmallNegative[i].Int64())
				n++
			}
		}
		benchmarkIntSinkVal = z
	})
	b.Run("most_significant_bit/int256", func(b *testing.B) {
		b.ReportAllocs()
		var z uint8
		b.ResetTimer()
		for n := 0; n < b.N; {
			for i := 0; i < numSamples && n < b.N; i++ {
				z = int256Samples[i].MostSignificantBit()
				n++
			}
		}
		benchmarkUint8Sink = z
	})
}
