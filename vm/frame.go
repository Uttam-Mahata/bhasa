package vm

import (
	"bhasa/code"
	"bhasa/object"
)

// Frame represents a function call frame
type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

// NewFrame creates a new frame
func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}
}

// Instructions returns the instructions for this frame
func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}

