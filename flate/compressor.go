package flate

import (
	"compress/flate"
	"errors"
	"github.com/lsilencej/compressor"
	"io"
)

type deflateAlgorithm struct {
	level compressor.Level
}

type deflateCompressor struct {
	write *flate.Writer
}

type deflateUnCompressor struct {
	read io.ReadCloser
}

func (c *deflateCompressor) Write(data []byte) (int, error) {
	return c.write.Write(data)
}

func (c *deflateCompressor) Close() error {
	return c.write.Close()
}

func (uc *deflateUnCompressor) Read(data []byte) (int, error) {
	return uc.read.Read(data)
}

func (uc *deflateUnCompressor) Close() error {
	return uc.read.Close()
}

func (d *deflateAlgorithm) NewAlgorithm() compressor.Algorithm {
	return &deflateAlgorithm{}
}

func (d *deflateAlgorithm) NewCompressor(w io.Writer) (compressor.Compressor, error) {
	c := &deflateCompressor{}
	var err error
	if c.write, err = flate.NewWriter(w, int(d.level)); err != nil {
		return nil, err
	}
	return c, err
}

func (d *deflateAlgorithm) NewUnCompressor(r io.Reader) (compressor.UnCompressor, error) {
	uc := &deflateUnCompressor{}
	var err error
	uc.read = flate.NewReader(r)
	return uc, err
}

func (d *deflateAlgorithm) Compress(data []byte) ([]byte, error) {
	return compressor.Compress(d, data)
}

func (d *deflateAlgorithm) UnCompress(data []byte) ([]byte, error) {
	return compressor.UnCompress(d, data)
}

func (d *deflateAlgorithm) SetLevel(level compressor.Level) error {
	d.level = level
	return nil
}

func (d *deflateAlgorithm) SetOrder(order compressor.Order) error {
	return errors.New("algorithm flate don't need order")
}

func (d *deflateAlgorithm) SetLitWidth(litWidth compressor.LitWidth) error {
	return errors.New("algorithm flate don't need litWidth")
}

func init() {
	compressor.Register("flate", &deflateAlgorithm{})
}
