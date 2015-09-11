// Copyright 2015 Vladimír Chalupecký. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparse

type Matrix interface {
	Dims() (r, c int)
	At(r, c int) float64
}

type Triplet struct {
	Row, Col int
	Value    float64
}

type rowWise []Triplet

func (r rowWise) Len() int      { return len(r) }
func (r rowWise) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r rowWise) Less(i, j int) bool {
	return r[i].Row < r[j].Row || (r[i].Row == r[j].Row && r[i].Col < r[j].Col)
}

// MulMatVec computes y = alpha * A * x + y.
func MulMatVec(alpha float64, a Matrix, x []float64, incx int, y []float64, incy int) {
	switch a := a.(type) {
	case *CSR:
		csrMulMatVec(alpha, a, x, incx, y, incy)
	case *DOK:
		dokMulMatVec(alpha, a, x, incx, y, incy)
	default:
		panic("unsupported matrix type")
	}
}
