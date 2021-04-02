package rsmt2d

import "fmt"

type bitMatrix struct {
	mask       []uint64
	squareSize int
}

func NewBitMatrix(squareSize, bits int) bitMatrix {
	if squareSize*squareSize != bits {
		panic(fmt.Sprintf("invalid arguments %v*%v != %v", squareSize, squareSize, bits))
	}
	return bitMatrix{mask: make([]uint64, (bits+63)/64), squareSize: squareSize}
}

// i = rowIndex*squareSize+colIdx
func (bm bitMatrix) SetFlat(i int) {
	bm.mask[i/64] |= uint64(1) << uint(i%64)
}

func (bm bitMatrix) Get(row, col int) bool {
	idx := row*bm.squareSize + col
	return bm.mask[idx/64]&(uint64(1)<<uint(idx%64)) > 0
}

func (bm *bitMatrix) Set(row, col int) {
	idx := row*bm.squareSize + col
	bm.mask[idx/64] |= uint64(1) << uint(idx%64)
}

func (bm bitMatrix) ColumnIsOne(c int) bool {
	for r := 0; r < bm.squareSize; r++ {
		if !bm.Get(r, c) {
			return false
		}
	}
	return true
}

func (bm bitMatrix) RowIsOne(r int) bool {
	for c := 0; c < bm.squareSize; c++ {
		if !bm.Get(r, c) {
			return false
		}
	}
	return true
}

func (bm bitMatrix) NumOnesInRow(r int) int {
	var counter int
	for i := 0; i < bm.squareSize; i++ {
		if bm.Get(r, i) {
			counter++
		}
	}

	return counter
}

func (bm bitMatrix) NumOnesInCol(c int) int {
	var counter int
	for i := 0; i < bm.squareSize; i++ {
		if bm.Get(i, c) {
			counter++
		}
	}

	return counter
}

func (bm bitMatrix) RowRangeIsOne(r, start, end int) bool {
	for c := start; c < end && c < bm.squareSize; c++ {
		if !bm.Get(r, c) {
			return false
		}
	}
	return true
}

func (bm bitMatrix) ColRangeIsOne(c, start, end int) bool {
	for r := start; r < end && r < bm.squareSize; r++ {
		if !bm.Get(r, c) {
			return false
		}
	}
	return true
}
