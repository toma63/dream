package dream
// object which represents a flat design

import (
	"toma/geom2d"
)

type DBlock struct {
	Version float64
	Design string
	DBU int
	Insts map[string]DInst
	Nets map[string]DNet
}

type DInst struct {
	InstName string
	ModelName string
	Point geom2d.Point
	Orient ORIENT
}

type ORIENT int 

const (
	N ORIENT = iota
	W
	S
	E
	FN
	FW
	FS
	FE
)

type DNet struct {
	Name string
	InstPins []DInstPin
}

type DInstPin struct {
	InstName string
	Pin string
}
