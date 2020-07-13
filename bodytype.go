package rq

// bodyType represents all values allowed of a body request
type bodyType int

const (
	_json bodyType = iota
	text
)
