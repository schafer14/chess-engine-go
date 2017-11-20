package maurice



func (p Position) HumanFriendlyMoves() []string {
	allMoves := p.legalMoves()

	strMoves := make([]string, 0)

	for _, m := range allMoves {
		strMoves = append(strMoves, m.toString())
	}

	return strMoves
}

func (p Position) legalMoves() []Move {
	allMoves := p.pseudoMoves()
	var legalMoves []Move

	for _, m := range allMoves {
		isLegal := true
		nb := p.Move(m)
		var kings Bitboard

		if nb.color == 0 {
			kings = nb.kings(1)
		} else {
			kings = nb.kings(0)
		}

		if nb.attacks(nb.color) & kings > 0 {
			isLegal = false
		}

		if isLegal {
			legalMoves = append(legalMoves, m)
		}
	}



	return legalMoves
}

func (p Position) pseudoMoves() []Move {
	allMoves := make([]Move, 0)
	allMoves = append(allMoves, p.pawnMoves()...)
	allMoves = append(allMoves, p.knightMoves()...)
	allMoves = append(allMoves, p.bishopMoves()...)
	allMoves = append(allMoves, p.rookMoves()...)
	allMoves = append(allMoves, p.queenMoves()...)
	allMoves = append(allMoves, p.kingMoves()...)

	allMoves = append(allMoves, p.castle()...)


	return allMoves
}

func (p Position) attacks(color int) Bitboard {
	attacks := Bitboard(0)

	attacks |= p.pawnAttacks(color)
	attacks |= p.knightAttacks(color)
	attacks |= p.bishopAttacks(color)
	attacks |= p.rookAttacks(color)
	attacks |= p.queenAttacks(color)
	attacks |= p.kingAttacks(color)

	return attacks
}

/*
	Given a bitboard and a function to make a move given the end position
	will create a list of moves based on each bit in the resulting bitboard
	being a valid destination
*/
func movesFromBitboard(bb Bitboard, fn func(Bitboard) Bitboard) []Move {
	moves := make([]Move, 0)

	for bb > 0 {
		square := bb & -bb
		bb &= bb-1

		moves = append(moves, Move{fn(square), square, ""})
	}

	return moves
}