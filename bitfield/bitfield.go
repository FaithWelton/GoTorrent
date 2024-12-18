package bitfield

type Bitfield []byte // Represents the pieces a peer has

// Check if a BitField has a particular index set
func (bf Bitfield) HasPiece(index int) bool {
	byteIndex := index / 8

	offset := index % 8

	if byteIndex < 0 || byteIndex >= len(bf) {
		return false
	}

	return bf[byteIndex]>>(7-offset)&1 != 0
}

// Set a bit in the Bitfield
func (bf Bitfield) SetPiece(index int) {
	byteIndex := index / 8

	offset := index % 8

	// Discard invalid bounded index
	if byteIndex < 0 || byteIndex >= len(bf) {
		return
	}

	bf[byteIndex] |= 1 << (7 - offset)
}
