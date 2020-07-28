package types

const (
	NEXT_IMAGE = iota
	PREV_IMAGE
	JUMP_TO_IMAGE

	NEXT_DIR
	PREV_DIR
	JUMP_TO_DIR
)

var IsImage = map[string]bool{
	".jpg": true,
	".png": true,
}
