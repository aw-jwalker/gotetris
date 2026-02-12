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

// Offset represents a (row, col) offset within a piece's bounding box
type Offset struct {
	Row int
	Col int
}

// pieceShapes defines the 4 cells for each piece at each rotation state
// Coordinates are relative to the top-left of the bounding box
// I piece uses 4×4 box, others use 3×3 box
var pieceShapes = map[PieceType][4][4]Offset{
	PieceI: {
		// State 0 (spawn): horizontal in middle of 4×4 box
		{
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
			{Row: 1, Col: 2},
			{Row: 1, Col: 3},
		},
		// State R: vertical on right side
		{
			{Row: 0, Col: 2},
			{Row: 1, Col: 2},
			{Row: 2, Col: 2},
			{Row: 3, Col: 2},
		},
		// State 2: horizontal in middle (row 2)
		{
			{Row: 2, Col: 0},
			{Row: 2, Col: 1},
			{Row: 2, Col: 2},
			{Row: 2, Col: 3},
		},
		// State L: vertical on left side
		{
			{Row: 0, Col: 1},
			{Row: 1, Col: 1},
			{Row: 2, Col: 1},
			{Row: 3, Col: 1},
		},
	},

	PieceO: {
		// State 0 (all states same): 2×2 square in top-left of 3×3 box
		{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		// States R, 2, L: identical (O piece doesn't rotate visually)
		{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
	},

	PieceT: {
		// State 0: T pointing up (flat bottom)
		{
			{Row: 0, Col: 1}, // Top center
			{Row: 1, Col: 0}, // Bottom left
			{Row: 1, Col: 1}, // Bottom center
			{Row: 1, Col: 2}, // Bottom right
		},
		// State R: T pointing right
		{
			{Row: 0, Col: 1}, // Top
			{Row: 1, Col: 1}, // Center
			{Row: 1, Col: 2}, // Right
			{Row: 2, Col: 1}, // Bottom
		},
		// State 2: T pointing down (flat top)
		{
			{Row: 1, Col: 0}, // Top left
			{Row: 1, Col: 1}, // Top center
			{Row: 1, Col: 2}, // Top right
			{Row: 2, Col: 1}, // Bottom center
		},
		// State L: T pointing left
		{
			{Row: 0, Col: 1}, // Top
			{Row: 1, Col: 0}, // Left
			{Row: 1, Col: 1}, // Center
			{Row: 2, Col: 1}, // Bottom
		},
	},

	PieceS: {
		// State 0: horizontal S (right high)
		{
			{Row: 0, Col: 1}, // Top right
			{Row: 0, Col: 2}, // Top right
			{Row: 1, Col: 0}, // Bottom left
			{Row: 1, Col: 1}, // Bottom left
		},
		// State R: vertical S
		{
			{Row: 0, Col: 1}, // Top
			{Row: 1, Col: 1}, // Center
			{Row: 1, Col: 2}, // Center right
			{Row: 2, Col: 2}, // Bottom right
		},
		// State 2: horizontal S (same as 0)
		{
			{Row: 0, Col: 1},
			{Row: 0, Col: 2},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		// State L: vertical S (same as R)
		{
			{Row: 0, Col: 1},
			{Row: 1, Col: 1},
			{Row: 1, Col: 2},
			{Row: 2, Col: 2},
		},
	},

	PieceZ: {
		// State 0: horizontal Z (left high)
		{
			{Row: 0, Col: 0}, // Top left
			{Row: 0, Col: 1}, // Top left
			{Row: 1, Col: 1}, // Bottom right
			{Row: 1, Col: 2}, // Bottom right
		},
		// State R: vertical Z
		{
			{Row: 0, Col: 2}, // Top right
			{Row: 1, Col: 1}, // Center
			{Row: 1, Col: 2}, // Center right
			{Row: 2, Col: 1}, // Bottom left
		},
		// State 2: horizontal Z (same as 0)
		{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 1},
			{Row: 1, Col: 2},
		},
		// State L: vertical Z (same as R)
		{
			{Row: 0, Col: 2},
			{Row: 1, Col: 1},
			{Row: 1, Col: 2},
			{Row: 2, Col: 1},
		},
	},

	PieceJ: {
		// State 0: J with bottom-left corner
		{
			{Row: 0, Col: 0}, // Top left
			{Row: 1, Col: 0}, // Bottom left
			{Row: 1, Col: 1}, // Bottom center
			{Row: 1, Col: 2}, // Bottom right
		},
		// State R: J with top-left corner
		{
			{Row: 0, Col: 1}, // Top
			{Row: 0, Col: 2}, // Top right
			{Row: 1, Col: 1}, // Center
			{Row: 2, Col: 1}, // Bottom
		},
		// State 2: J with top-right corner
		{
			{Row: 1, Col: 0}, // Top left
			{Row: 1, Col: 1}, // Top center
			{Row: 1, Col: 2}, // Top right
			{Row: 2, Col: 2}, // Bottom right
		},
		// State L: J with bottom-right corner
		{
			{Row: 0, Col: 1}, // Top
			{Row: 1, Col: 1}, // Center
			{Row: 2, Col: 0}, // Bottom left
			{Row: 2, Col: 1}, // Bottom
		},
	},

	PieceL: {
		// State 0: L with bottom-right corner
		{
			{Row: 0, Col: 2}, // Top right
			{Row: 1, Col: 0}, // Bottom left
			{Row: 1, Col: 1}, // Bottom center
			{Row: 1, Col: 2}, // Bottom right
		},
		// State R: L with top-right corner
		{
			{Row: 0, Col: 1}, // Top
			{Row: 1, Col: 1}, // Center
			{Row: 2, Col: 1}, // Bottom
			{Row: 2, Col: 2}, // Bottom right
		},
		// State 2: L with top-left corner
		{
			{Row: 1, Col: 0}, // Top left
			{Row: 1, Col: 1}, // Top center
			{Row: 1, Col: 2}, // Top right
			{Row: 2, Col: 0}, // Bottom left
		},
		// State L: L with bottom-left corner
		{
			{Row: 0, Col: 0}, // Top left
			{Row: 0, Col: 1}, // Top
			{Row: 1, Col: 1}, // Center
			{Row: 2, Col: 1}, // Bottom
		},
	},
}

// Cells returns the absolute board coordinates of the 4 cells
// that make up this piece in its current rotation state
func (p *Piece) Cells() [4]Offset {
	offsets := pieceShapes[p.Type][p.Rotation]
	var result [4]Offset

	for i, offset := range offsets {
		result[i] = Offset{
			Row: p.Row + offset.Row,
			Col: p.Col + offset.Col,
		}
	}

	return result
}

// WallKickOffset represents a wall kick test position
type WallKickOffset struct {
	Row int
	Col int
}

// wallKickData defines the 5 test positions for each rotation transition
// Format: [fromRotation][toRotation][testIndex]
// Tests are tried in order; first successful position is used

// Wall kicks for J, L, S, T, Z pieces (standard kicks)
var wallKicksJLSTZ = map[RotationState]map[RotationState][5]WallKickOffset{
	Rotation0: {
		RotationR: {
			{Row: 0, Col: 0},   // Test 1: no offset
			{Row: 0, Col: -1},  // Test 2: left
			{Row: -1, Col: -1}, // Test 3: left and up
			{Row: 2, Col: 0},   // Test 4: down 2
			{Row: 2, Col: -1},  // Test 5: left and down 2
		},
		RotationL: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},  // Test 2: right
			{Row: -1, Col: 1}, // Test 3: right and up
			{Row: 2, Col: 0},  // Test 4: down 2
			{Row: 2, Col: 1},  // Test 5: right and down 2
		},
	},
	RotationR: {
		Rotation0: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 1},
			{Row: -2, Col: 0},
			{Row: -2, Col: 1},
		},
		Rotation2: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 1},
			{Row: -2, Col: 0},
			{Row: -2, Col: 1},
		},
	},
	Rotation2: {
		RotationR: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: -1, Col: 1},
			{Row: 2, Col: 0},
			{Row: 2, Col: 1},
		},
		RotationL: {
			{Row: 0, Col: 0},
			{Row: 0, Col: -1},
			{Row: -1, Col: -1},
			{Row: 2, Col: 0},
			{Row: 2, Col: -1},
		},
	},
	RotationL: {
		Rotation0: {
			{Row: 0, Col: 0},
			{Row: 0, Col: -1},
			{Row: 1, Col: -1},
			{Row: -2, Col: 0},
			{Row: -2, Col: -1},
		},
		Rotation2: {
			{Row: 0, Col: 0},
			{Row: 0, Col: -1},
			{Row: 1, Col: -1},
			{Row: -2, Col: 0},
			{Row: -2, Col: -1},
		},
	},
}

// Wall kicks for I piece (larger kicks due to 4×4 bounding box)
var wallKicksI = map[RotationState]map[RotationState][5]WallKickOffset{
	Rotation0: {
		RotationR: {
			{Row: 0, Col: 0},
			{Row: 0, Col: -2},
			{Row: 0, Col: 1},
			{Row: 1, Col: -2},
			{Row: -2, Col: 1},
		},
		RotationL: {
			{Row: 0, Col: 0},
			{Row: 0, Col: -1},
			{Row: 0, Col: 2},
			{Row: -2, Col: -1},
			{Row: 1, Col: 2},
		},
	},
	RotationR: {
		Rotation0: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 2},
			{Row: 0, Col: -1},
			{Row: -1, Col: 2},
			{Row: 2, Col: -1},
		},
		Rotation2: {
			{Row: 0, Col: 0},
			{Row: 0, Col: -1},
			{Row: 0, Col: 2},
			{Row: -2, Col: -1},
			{Row: 1, Col: 2},
		},
	},
	Rotation2: {
		RotationR: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 0, Col: -2},
			{Row: 2, Col: 1},
			{Row: -1, Col: -2},
		},
		RotationL: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 2},
			{Row: 0, Col: -1},
			{Row: -1, Col: 2},
			{Row: 2, Col: -1},
		},
	},
	RotationL: {
		Rotation0: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 0, Col: -2},
			{Row: 2, Col: 1},
			{Row: -1, Col: -2},
		},
		Rotation2: {
			{Row: 0, Col: 0},
			{Row: 0, Col: 2},
			{Row: 0, Col: -1},
			{Row: -1, Col: 2},
			{Row: 2, Col: -1},
		},
	},
}

// GetWallKicks returns the 5 wall kick test positions for rotating
// from the current rotation state to the target rotation state
func (p *Piece) GetWallKicks(targetRotation RotationState) [5]WallKickOffset {
	// O piece doesn't kick
	if p.Type == PieceO {
		return [5]WallKickOffset{{Row: 0, Col: 0}}
	}

	// I piece has special wall kicks
	if p.Type == PieceI {
		return wallKicksI[p.Rotation][targetRotation]
	}

	// J, L, S, T, Z use standard wall kicks
	return wallKicksJLSTZ[p.Rotation][targetRotation]
}
