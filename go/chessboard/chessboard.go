package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	rank := cb[file]
	total := 0
	for _, sq := range rank {
		if sq {
			total++
		}
	}
	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	total := 0
    if(rank < 1 || rank >  8) {
		return 0
	}
	for _, files := range cb {
		if files[rank-1] {
			total++
		}
	}
	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	total := 0
	for _, files := range cb {
		total += len(files)
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
	for file := range cb {
		total += CountInFile(cb, file)
	}
	return total
}
