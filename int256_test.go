// Package int256: Fixed-size signed integer math library represented as a 256-bit absolute value plus a sign bit.
// Copyright 2018-2020 uint256 Authors
// Copyright (c) 2023 Trịnh Đức Bảo Linh(Kevin)
// Copyright (c) 2026 0xsimulacra
// SPDX-License-Identifier: MIT AND BSD-3-Clause
package int256

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	"github.com/holiman/uint256"
)

func TestInt_Add(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct when add two positive numbers",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
				},
				y: &Int{
					abs: *uint256.NewInt(7),
				},
			},
			want: &Int{
				abs: *uint256.NewInt(17),
				neg: false,
			},
		},
		{
			name: "Should return correct when add two negative numbers",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(17),
				neg: true,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|>|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(7),
					neg: false,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
				neg: true,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|<|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
				neg: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Add(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sub(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct when sub two positive numbers |a|>|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
				},
				y: &Int{
					abs: *uint256.NewInt(7),
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
				neg: false,
			},
		},
		{
			name: "Should return correct when sub two positive numbers |a|<|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(7),
				},
				y: &Int{
					abs: *uint256.NewInt(10),
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
				neg: true,
			},
		},
		{
			name: "Should return correct when add two negative numbers |a|>|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
				neg: true,
			},
		},
		{
			name: "Should return correct when add two negative numbers |a|<|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(7),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
				neg: false,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|>|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(7),
					neg: false,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(17),
				neg: true,
			},
		},
		{
			name: "Should return correct when add two numbers with a is negative, b is positive and |a|<|b|",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(7),
					neg: true,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(17),
				neg: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Sub(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Mul(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when multiple for two positive numbers",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(3),
				},
				y: &Int{
					abs: *uint256.NewInt(5),
				},
			},
			want: &Int{
				abs: *uint256.NewInt(15),
			},
		},
		{
			name: "Should return correct value when multiple for two negative numbers",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(5),
					neg: true,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(15),
				neg: false,
			},
		},
		{
			name: "Should return correct value when multiple for a negative number and a positive number",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(5),
					neg: false,
				},
			},
			want: &Int{
				abs: *uint256.NewInt(15),
				neg: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Mul(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_MulPanic(t *testing.T) {
	i256 := new(Int)

	defer func() {
		if r := recover(); r != nil {
			t.Error("should not have paniced", r)
		}
	}()

	res := i256.Mul(NewInt(1), NewInt(2))
	if res.Cmp(NewInt(2)) != 0 {
		t.Errorf("want: 2, got: %v", res)
	}
}

func TestInt_Sqrt(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when performing for positive number",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(9),
				},
			},
			want: &Int{
				abs: *uint256.NewInt(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Sqrt(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetString(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		s string
	}
	big1, _ := new(big.Int).SetString("-10a", 16)

	MaxUint256, _ := new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	big, _ := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when parsing correct string value",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "10",
			},
			want: &Int{
				abs: *uint256.NewInt(10),
				neg: false,
			},
			wantErr: false,
		},
		{
			name: "Should return correct value when parsing correct string value",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "-10",
			},
			want: &Int{
				abs: *uint256.NewInt(10),
				neg: true,
			},
			wantErr: false,
		},
		{
			name: "Should return error value when parsing incorrect string value",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "-10a",
			},
			want:    MustFromBig(big1),
			wantErr: false,
		},
		{
			name: "Should return error value when parsing correct string value",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "1461446703485210103287273052203988822378723970342",
			},
			want:    MustFromBig(big),
			wantErr: false,
		},
		{
			name: "Should return error value when parsing correct string value",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: false,
			},
			args: args{
				s: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			},
			want:    MustFromBig(MaxUint256),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			got, err := z.SetString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.SetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetInt64(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when ",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: 24,
			},
			want: &Int{
				abs: *uint256.NewInt(24),
				neg: false,
			},
		},
		{
			name: "Should return correct value when ",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: -24,
			},
			want: &Int{
				abs: *uint256.NewInt(24),
				neg: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.SetInt64(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Rsh(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		n uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when perform positive number",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: NewInt(10),
				n: 4,
			},
			want: MustFromBig(new(big.Int).Rsh(big.NewInt(10), 4)),
		},
		{
			name: "Should return correct value when perform negative number",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: NewInt(-10),
				n: 4,
			},
			want: MustFromBig(new(big.Int).Rsh(big.NewInt(-10), 4)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Rsh(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rsh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func int257TestValues() []*Int {
	values := []*Int{
		NewInt(0),
		NewInt(1),
		NewInt(-1),
		NewInt(2),
		NewInt(-2),
		NewInt(3),
		NewInt(-3),
		NewInt(10),
		NewInt(-10),
	}

	for _, bit := range []uint{1, 2, 63, 64, 65, 127, 128, 129, 191, 192, 193, 255} {
		pow := new(big.Int).Lsh(big.NewInt(1), bit)
		powMinusOne := new(big.Int).Sub(new(big.Int).Set(pow), big.NewInt(1))
		powPlusOne := new(big.Int).Add(new(big.Int).Set(pow), big.NewInt(1))
		for _, x := range []*big.Int{powMinusOne, pow, powPlusOne} {
			values = append(values, MustFromBig(x))
			values = append(values, MustFromBig(new(big.Int).Neg(x)))
		}
	}

	maxUint := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	values = append(values, MustFromBig(maxUint))
	values = append(values, MustFromBig(new(big.Int).Neg(maxUint)))

	return values
}

func TestInt_RshMatchesBigInt(t *testing.T) {
	values := int257TestValues()

	shifts := []uint{0, 1, 2, 3, 63, 64, 65, 127, 128, 129, 191, 192, 193, 254, 255, 256, 257, 300}

	for _, x := range values {
		xBig := x.ToBig()
		for _, shift := range shifts {
			t.Run(fmt.Sprintf("%s>>%d", xBig, shift), func(t *testing.T) {
				got := new(Int).Rsh(x, shift)
				want := MustFromBig(new(big.Int).Rsh(new(big.Int).Set(xBig), shift))
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("Rsh() = %v, want %v", got, want)
				}

				gotAlias := x.Clone()
				gotAlias.Rsh(gotAlias, shift)
				if !reflect.DeepEqual(gotAlias, want) {
					t.Fatalf("alias Rsh() = %v, want %v", gotAlias, want)
				}
			})
		}
	}
}

func TestInt_Rem(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when performing for two positive numbers",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
				},
				y: &Int{
					abs: *uint256.NewInt(3),
				},
			},
			want: &Int{
				abs: *uint256.NewInt(1),
			},
		},
		{
			name: "Should return correct value when performing for two negative numbers",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Rem(big.NewInt(-10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when performing for a negative and a numbers",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Rem(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when performing for a negative and a numbers",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Rem(big.NewInt(10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Rem(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Exp(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
		m *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when perform x,y,m is positive number",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(3), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(-3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(-3), big.NewInt(3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},

			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(3), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(10), big.NewInt(-3), big.NewInt(-3))),
		},
		{
			name: "Should return correct value when perform x,y is positive number and m is negative",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
				m: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Exp(big.NewInt(-10), big.NewInt(-3), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Exp(tt.args.x, tt.args.y, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_ExpOptimizedPathsMatchBig(t *testing.T) {
	tests := []struct {
		name string
		x    *big.Int
		y    *big.Int
		m    *big.Int
	}{
		{
			name: "positive with large uint64 modulus",
			x:    big.NewInt(123456789),
			y:    big.NewInt(17),
			m:    big.NewInt(8589934591),
		},
		{
			name: "negative base odd exponent with large uint64 modulus",
			x:    big.NewInt(-123456789),
			y:    big.NewInt(17),
			m:    big.NewInt(8589934591),
		},
		{
			name: "negative base even exponent without modulus",
			x:    big.NewInt(-3),
			y:    big.NewInt(4),
		},
		{
			name: "negative base odd exponent without modulus",
			x:    big.NewInt(-3),
			y:    big.NewInt(5),
		},
		{
			name: "negative base odd exponent with zero modulus",
			x:    big.NewInt(-3),
			y:    big.NewInt(5),
			m:    new(big.Int),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := MustFromBig(tt.x)
			y := MustFromBig(tt.y)
			var m *Int
			var mBig *big.Int
			if tt.m != nil {
				m = MustFromBig(tt.m)
				mBig = new(big.Int).Set(tt.m)
			}

			got := new(Int).Exp(x, y, m)
			want := MustFromBig(new(big.Int).Exp(new(big.Int).Set(tt.x), new(big.Int).Set(tt.y), mBig))
			if !got.Eq(want) {
				t.Fatalf("Exp() = %v, want %v", got, want)
			}
		})
	}
}

func TestInt_ModifiedOpsRandomizedMatchBig(t *testing.T) {
	rng := rand.New(rand.NewSource(0x1A2B3C4D))
	shifts := []uint{0, 1, 2, 3, 7, 31, 63, 64, 65, 95, 127, 128, 129, 159, 191, 192, 193, 223, 254, 255, 256, 257, 300}

	for i := 0; i < 256; i++ {
		xBig := randomSignedInt256(rng)
		yBig := randomSignedInt256(rng)
		x := MustFromBig(xBig)
		y := MustFromBig(yBig)

		assertBigOpResult(t, "Or", func() *Int { return new(Int).Or(x, y) },
			new(big.Int).Or(new(big.Int).Set(xBig), yBig))
		assertBigOpResult(t, "And", func() *Int { return new(Int).And(x, y) },
			new(big.Int).And(new(big.Int).Set(xBig), yBig))

		gotAliasX := x.Clone()
		assertBigOpResult(t, "alias x Or", func() *Int { return gotAliasX.Or(gotAliasX, y) },
			new(big.Int).Or(new(big.Int).Set(xBig), yBig))

		gotAliasY := y.Clone()
		assertBigOpResult(t, "alias y And", func() *Int { return gotAliasY.And(x, gotAliasY) },
			new(big.Int).And(new(big.Int).Set(xBig), yBig))

		for _, shift := range shifts {
			assertBigOpResult(t, fmt.Sprintf("Rsh/%d", shift), func() *Int {
				return new(Int).Rsh(x, shift)
			}, new(big.Int).Rsh(new(big.Int).Set(xBig), shift))

			assertBigOpResult(t, fmt.Sprintf("Lsh/%d", shift), func() *Int {
				return new(Int).Lsh(x, shift)
			}, new(big.Int).Lsh(new(big.Int).Set(xBig), shift))
		}
	}
}

func TestInt_ExpRandomizedMatchesBig(t *testing.T) {
	rng := rand.New(rand.NewSource(0x5EED1234))

	for i := 0; i < 256; i++ {
		xBig := big.NewInt(rng.Int63n(41) - 20)
		yBig := big.NewInt(rng.Int63n(32))
		modChoice := rng.Intn(5)

		var mBig *big.Int
		switch modChoice {
		case 0:
			mBig = nil
		case 1:
			mBig = new(big.Int)
		case 2:
			mBig = big.NewInt(rng.Int63n(1<<32-1) + 1)
		case 3:
			mBig = new(big.Int).SetUint64(rng.Uint64() | 1<<40)
		default:
			mBig = new(big.Int).Neg(new(big.Int).SetUint64(rng.Uint64() | 1))
		}

		x := MustFromBig(xBig)
		y := MustFromBig(yBig)
		var m *Int
		if mBig != nil {
			m = MustFromBig(mBig)
		}

		wantBig := new(big.Int).Exp(new(big.Int).Set(xBig), new(big.Int).Set(yBig), mBig)
		want := MustFromBig(wantBig)

		got := new(Int).Exp(x, y, m)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Exp(%s, %s, %v) = %v, want %v", xBig, yBig, mBig, got, want)
		}

		gotAlias := x.Clone()
		gotAlias.Exp(gotAlias, y, m)
		if !reflect.DeepEqual(gotAlias, want) {
			t.Fatalf("alias Exp(%s, %s, %v) = %v, want %v", xBig, yBig, mBig, gotAlias, want)
		}
	}
}

func TestInt_ExpNegativeExponentMutatesReceiver(t *testing.T) {
	z := NewInt(123)
	got := z.Exp(NewInt(2), NewInt(-1), NewInt(5))
	want := NewInt(3)

	if got != z {
		t.Fatalf("Exp() returned %p, want receiver %p", got, z)
	}
	if !reflect.DeepEqual(z, want) {
		t.Fatalf("receiver after Exp() = %v, want %v", z, want)
	}
}

func TestInt_ExpNegativeExponentNotInvertibleLeavesReceiver(t *testing.T) {
	z := NewInt(123)
	want := z.Clone()

	got := z.Exp(NewInt(2), NewInt(-1), NewInt(4))
	if got != nil {
		t.Fatalf("Exp() = %v, want nil", got)
	}
	if !reflect.DeepEqual(z, want) {
		t.Fatalf("receiver after failed Exp() = %v, want %v", z, want)
	}
}

func randomSignedInt256(rng *rand.Rand) *big.Int {
	z := new(big.Int)
	for i := 0; i < 4; i++ {
		z.Lsh(z, 64)
		z.Or(z, new(big.Int).SetUint64(rng.Uint64()))
	}
	if rng.Intn(2) == 0 {
		z.Neg(z)
	}
	return z
}

func assertBigOpResult(t *testing.T, name string, op func() *Int, wantBig *big.Int) {
	t.Helper()

	want, overflow := FromBig(wantBig)
	if overflow {
		if !causesPanic(func() { op() }) {
			t.Fatalf("%s should panic for result %s", name, wantBig)
		}
		return
	}

	got := op()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s = %v, want %v", name, got, want)
	}
}

func TestInt_Lsh(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		n uint
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				n: 3,
			},
			want: MustFromBig(new(big.Int).Lsh(big.NewInt(10), 3)),
		},
		{
			name: "Should return correct value when process negative number and n is odd",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				n: 3,
			},
			want: MustFromBig(new(big.Int).Lsh(big.NewInt(-10), 3)),
		},
		{
			name: "Should return correct value when process negative number and n is even",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				n: 4,
			},
			want: MustFromBig(new(big.Int).Lsh(big.NewInt(-10), 4)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Lsh(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Lsh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_LshMatchesBigInt(t *testing.T) {
	values := int257TestValues()
	shifts := []uint{0, 1, 2, 3, 63, 64, 65, 127, 128, 129, 191, 192, 193, 254, 255, 256, 257, 300}

	for _, x := range values {
		xBig := x.ToBig()
		for _, shift := range shifts {
			if bitLen := x.abs.BitLen(); bitLen != 0 && shift > uint(256-bitLen) {
				continue
			}
			t.Run(fmt.Sprintf("%s<<%d", xBig, shift), func(t *testing.T) {
				want := MustFromBig(new(big.Int).Lsh(new(big.Int).Set(xBig), shift))

				got := new(Int).Lsh(x, shift)
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("Lsh() = %v, want %v", got, want)
				}

				gotAlias := x.Clone()
				gotAlias.Lsh(gotAlias, shift)
				if !reflect.DeepEqual(gotAlias, want) {
					t.Fatalf("alias Lsh() = %v, want %v", gotAlias, want)
				}
			})
		}
	}
}

func TestInt_Or(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(-10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Or(big.NewInt(-10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Or(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_And(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(-10), big.NewInt(3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).And(big.NewInt(-10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.And(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_BitwiseMatchesBigInt(t *testing.T) {
	values := int257TestValues()

	for _, x := range values {
		xBig := x.ToBig()
		for _, y := range values {
			yBig := y.ToBig()

			t.Run(fmt.Sprintf("%s|%s", xBig, yBig), func(t *testing.T) {
				want, overflow := FromBig(new(big.Int).Or(new(big.Int).Set(xBig), yBig))
				if overflow {
					if !causesPanic(func() { new(Int).Or(x, y) }) {
						t.Fatalf("Or() should panic on overflow")
					}
					return
				}

				got := new(Int).Or(x, y)
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("Or() = %v, want %v", got, want)
				}

				gotAliasX := x.Clone()
				gotAliasX.Or(gotAliasX, y)
				if !reflect.DeepEqual(gotAliasX, want) {
					t.Fatalf("alias x Or() = %v, want %v", gotAliasX, want)
				}

				gotAliasY := y.Clone()
				gotAliasY.Or(x, gotAliasY)
				if !reflect.DeepEqual(gotAliasY, want) {
					t.Fatalf("alias y Or() = %v, want %v", gotAliasY, want)
				}
			})

			t.Run(fmt.Sprintf("%s&%s", xBig, yBig), func(t *testing.T) {
				want, overflow := FromBig(new(big.Int).And(new(big.Int).Set(xBig), yBig))
				if overflow {
					if !causesPanic(func() { new(Int).And(x, y) }) {
						t.Fatalf("And() should panic on overflow")
					}
					return
				}

				got := new(Int).And(x, y)
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("And() = %v, want %v", got, want)
				}

				gotAliasX := x.Clone()
				gotAliasX.And(gotAliasX, y)
				if !reflect.DeepEqual(gotAliasX, want) {
					t.Fatalf("alias x And() = %v, want %v", gotAliasX, want)
				}

				gotAliasY := y.Clone()
				gotAliasY.And(x, gotAliasY)
				if !reflect.DeepEqual(gotAliasY, want) {
					t.Fatalf("alias y And() = %v, want %v", gotAliasY, want)
				}
			})
		}
	}
}

func TestInt_Quo(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value when perform two positive numbers",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(10), big.NewInt(3))),
		},
		{
			name: "Should return correct value perform a < 0 and b > 0",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: false,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(-10), big.NewInt(3))),
		},
		{
			name: "Should return correct value perform a >0 and b < 0",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: false,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(10), big.NewInt(-3))),
		},
		{
			name: "Should return correct value perform two negative numbers",
			fields: fields{
				abs: uint256.Int{},
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
					neg: true,
				},
				y: &Int{
					abs: *uint256.NewInt(3),
					neg: true,
				},
			},
			want: MustFromBig(new(big.Int).Quo(big.NewInt(-10), big.NewInt(-3))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if got := z.Quo(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Quo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Cmp(t *testing.T) {
	type fields struct {
		abs uint256.Int
		neg bool
	}
	type args struct {
		x *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  int
	}{
		// TODO: Add test cases.
		{
			name: "Should return correct value",
			fields: fields{
				abs: *uint256.NewInt(10),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
				},
			},
			wantR: 0,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: *uint256.NewInt(4),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(10),
				},
			},
			wantR: -1,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: *uint256.NewInt(10),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(4),
				},
			},
			wantR: 1,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: *uint256.NewInt(10),
				neg: false,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(0),
				},
			},
			wantR: 1,
		},
		{
			name: "Should return correct value",
			fields: fields{
				abs: *uint256.NewInt(0),
				neg: true,
			},
			args: args{
				x: &Int{
					abs: *uint256.NewInt(0),
				},
			},
			wantR: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				abs: tt.fields.abs,
				neg: tt.fields.neg,
			}
			if gotR := z.Cmp(tt.args.x); gotR != tt.wantR {
				t.Errorf("Int.Cmp() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
