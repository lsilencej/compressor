package compressor

import (
	"fmt"
	"io"
)

var algorithms = make(map[string]Algorithm)

// Algorithm interface
type Algorithm interface {
	NewAlgorithm() Algorithm
	NewCompressor(w io.Writer) (Compressor, error)
	NewUnCompressor(r io.Reader) (UnCompressor, error)
	Compress(data []byte) ([]byte, error)
	UnCompress(data []byte) ([]byte, error)
	SetLevel(level Level) error
	SetOrder(order Order) error
	SetLitWidth(litWidth int) error
}

// Compressor interface
type Compressor interface {
	Write(data []byte) (int, error)
	Close() error
}

// UnCompressor interface
type UnCompressor interface {
	Read(data []byte) (int, error)
	Close() error
}

// Level compression level
type Level int

const (
	NoCompression      Level = 0
	BestSpeed          Level = 1
	BestCompression    Level = 9
	DefaultCompression Level = -1
	HuffmanOnly        Level = -2
)

// Order specifies the bit ordering in an LZW data stream.
type Order int

const (
	// LSB means Least Significant Bits first, as used in the GIF file format.
	LSB Order = iota
	// MSB means Most Significant Bits first, as used in the TIFF and PDF
	// file formats.
	MSB
)

// Register algorithms
func Register(name string, algorithm Algorithm) {
	algorithms[name] = algorithm
}

// Registered check the algorithm's state of register
func Registered(name string) error {
	_, ok := algorithms[name]
	if !ok {
		return fmt.Errorf("algorithm not registered: %s", name)
	}
	return nil
}
