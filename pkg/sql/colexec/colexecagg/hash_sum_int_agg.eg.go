// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecagg

import (
	"strings"
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

func newSumIntHashAggAlloc(
	allocator *colmem.Allocator, t *types.T, allocSize int64,
) (aggregateFuncAlloc, error) {
	allocBase := aggAllocBase{allocator: allocator, allocSize: allocSize}
	switch t.Family() {
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &sumIntInt16HashAggAlloc{aggAllocBase: allocBase}, nil
		case 32:
			return &sumIntInt32HashAggAlloc{aggAllocBase: allocBase}, nil
		default:
			return &sumIntInt64HashAggAlloc{aggAllocBase: allocBase}, nil
		}
	default:
		return nil, errors.Errorf("unsupported sum %s agg type %s", strings.ToLower("Int"), t.Name())
	}
}

type sumIntInt16HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg int64
		// vec points to the output vector we are updating.
		vec []int64
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ AggregateFunc = &sumIntInt16HashAgg{}

func (a *sumIntInt16HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Int64()
}

func (a *sumIntInt16HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int16(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntInt16HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumIntInt16HashAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt16HashAgg
}

var _ aggregateFuncAlloc = &sumIntInt16HashAggAlloc{}

const sizeOfSumIntInt16HashAgg = int64(unsafe.Sizeof(sumIntInt16HashAgg{}))
const sumIntInt16HashAggSliceOverhead = int64(unsafe.Sizeof([]sumIntInt16HashAgg{}))

func (a *sumIntInt16HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sumIntInt16HashAggSliceOverhead + sizeOfSumIntInt16HashAgg*a.allocSize)
		a.aggFuncs = make([]sumIntInt16HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntInt32HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg int64
		// vec points to the output vector we are updating.
		vec []int64
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ AggregateFunc = &sumIntInt32HashAgg{}

func (a *sumIntInt32HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Int64()
}

func (a *sumIntInt32HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int32(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntInt32HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumIntInt32HashAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt32HashAgg
}

var _ aggregateFuncAlloc = &sumIntInt32HashAggAlloc{}

const sizeOfSumIntInt32HashAgg = int64(unsafe.Sizeof(sumIntInt32HashAgg{}))
const sumIntInt32HashAggSliceOverhead = int64(unsafe.Sizeof([]sumIntInt32HashAgg{}))

func (a *sumIntInt32HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sumIntInt32HashAggSliceOverhead + sizeOfSumIntInt32HashAgg*a.allocSize)
		a.aggFuncs = make([]sumIntInt32HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntInt64HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg int64
		// vec points to the output vector we are updating.
		vec []int64
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ AggregateFunc = &sumIntInt64HashAgg{}

func (a *sumIntInt64HashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.scratch.vec = vec.Int64()
}

func (a *sumIntInt64HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int64(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{
						result := int64(a.scratch.curAgg) + int64(col[i])
						if (result < int64(a.scratch.curAgg)) != (int64(col[i]) < 0) {
							colexecerror.ExpectedError(tree.ErrIntOutOfRange)
						}
						a.scratch.curAgg = result
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntInt64HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumIntInt64HashAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntInt64HashAgg
}

var _ aggregateFuncAlloc = &sumIntInt64HashAggAlloc{}

const sizeOfSumIntInt64HashAgg = int64(unsafe.Sizeof(sumIntInt64HashAgg{}))
const sumIntInt64HashAggSliceOverhead = int64(unsafe.Sizeof([]sumIntInt64HashAgg{}))

func (a *sumIntInt64HashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sumIntInt64HashAggSliceOverhead + sizeOfSumIntInt64HashAgg*a.allocSize)
		a.aggFuncs = make([]sumIntInt64HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
