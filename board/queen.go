package board

import "math"

func queenMoves(board board, color string) []move {
	var friendly uint64
	var bb uint64
	var occ uint64 = occupied(board)
	var moves []move = make([]move, 0)

	if (color == "w") {
		bb = board.whiteQueens
		friendly = whitePieces(board)
	} else {
		bb = board.blackQueens
		friendly = blackPieces(board)
	}


	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := uint(math.Log2(float64(square)))

		moveBb := diagBB(occ, squareNum) | straightBB(occ, squareNum)
		legalMovesBb := moveBb & (^friendly)

		newMoves := bbToMoves(legalMovesBb, func(_ uint64) uint64 {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func queenAttackBB(board board, color string) uint64 {
	var bb uint64
	var occ uint64 = occupied(board)
	var attackBB uint64 = 0

	if (color == "w") {
		bb = board.whiteQueens
	} else {
		bb = board.blackQueens
	}


	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := uint(math.Log2(float64(square)))

		attackBB |= diagBB(occ, squareNum) | straightBB(occ, squareNum)

	}

	return attackBB
}