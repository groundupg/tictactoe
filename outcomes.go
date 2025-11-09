package main

func Win(b Board, p Player) bool {
	if horizontal_win(b, p) {
		return true
	}
	if vertical_win(b, p) {
		return true
	}
	if diagonal_win(b, p) {
		return true
	}
	return false
}

func Draw(b Board) bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func horizontal_win(b Board, p Player) bool {
	for i := 0; i < len(b); i++ {
		if p == b[i][0] {
			if p == b[i][1] {
				if p == b[i][2] {
					return true
				}
			}
		}
	}
	return false
}

func diagonal_win(b Board, p Player) bool {
	if p == b[1][1] {
		if p == b[0][0] {
			if p == b[2][2] {
				return true
			}
		}
		if p == b[0][2] {
			if p == b[2][0] {
				return true
			}
		}
	}
	return false
}

func vertical_win(b Board, p Player) bool {
	for i := 0; i < len(b); i++ {
		if p == b[0][i] {
			if p == b[1][i] {
				if p == b[2][i] {
					return true
				}
			}
		}
	}
	return false
}
