package main

// PieceType represents one of the 7 standard Tetris pieces
type PieceType int

const (
	PieceI PieceType = iota // Cyan - straight piece
	PieceO                  // Yellow - square piece
	PieceT                  // Purple - T-shaped piece
	PieceS                  // Green - S-shaped piece
	PieceZ                  // Red - Z-shaped piece
	PieceJ                  // Blue - J-shaped piece
	PieceL                  // Orange - L-shaped piece
)

// RotationState represents one of 4 rotation states (0, R, 2, L)
type RotationState int

const (
	Rotation0 RotationState = iota // Spawn state
	RotationR                      // 90° clockwise
	Rotation2                      // 180°
	RotationL                      // 90° counter-clockwise
)

// Piece represents a tetromino with position and rotation
type Piece struct {
	Type     PieceType
	Rotation RotationState
	Row      int // Top-left corner of bounding box
	Col      int // Top-left corner of bounding box
}

// NewPiece creates a new piece at the specified position
func NewPiece(pieceType PieceType, row, col int) *Piece {
	return &Piece{
		Type:     pieceType,
		Rotation: Rotation0,
		Row:      row,
		Col:      col,
	}
}

// Color returns the color for this piece type
func (p *Piece) Color() CellColor {
	switch p.Type {
	case PieceI:
		return ColorCyan
	case PieceO:
		return ColorYellow
	case PieceT:
		return ColorPurple
	case PieceS:
		return ColorGreen
	case PieceZ:
		return ColorRed
	case PieceJ:
		return ColorBlue
	case PieceL:
		return ColorOrange
	default:
		return ColorEmpty
	}
}
