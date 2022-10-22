package zlib

import (
	"compress/zlib"
	"errors"
	"github.com/lsilencej/compressor"
	"io"
)

type zlibAlgorithm struct {
	level compressor.Level
}

type zlibCompressor struct {
	writer io.WriteCloser
}

type zlibUnCompressor struct {
	reader io.ReadCloser
}

func (c *zlibCompressor) Write(data []byte) (int, error) {
	return c.writer.Write(data)
}

func (c *zlibCompressor) Close() error {
	return c.writer.Close()
}

func (uc *zlibUnCompressor) Read(data []byte) (int, error) {
	return uc.reader.Read(data)
}

func (uc *zlibUnCompressor) Close() error {
	return uc.reader.Close()
}

func (z *zlibAlgorithm) NewAlgorithm() compressor.Algorithm {
	return &zlibAlgorithm{}
}

func (z *zlibAlgorithm) NewCompressor(w io.Writer) (compressor.Compressor, error) {
	c := &zlibCompressor{}
	if z.level == 0 {
		c.writer = zlib.NewWriter(w)
	} else {
		var err error
		if c.writer, err = zlib.NewWriterLevel(w, int(z.level)); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (z *zlibAlgorithm) NewUnCompressor(r io.Reader) (compressor.UnCompressor, error) {
	uc := &zlibUnCompressor{}
	var err error
	if uc.reader, err = zlib.NewReader(r); err != nil {
		return nil, err
	}
	return uc, nil
}

func (z *zlibAlgorithm) Compress(data []byte) ([]byte, error) {
	return compressor.Compress(z, data)
}

func (z *zlibAlgorithm) UnCompress(data []byte) ([]byte, error) {
	return compressor.UnCompress(z, data)
}

func (z *zlibAlgorithm) SetLevel(level compressor.Level) error {
	z.level = level
	return nil
}

func (z *zlibAlgorithm) SetOrder(order compressor.Order) error {
	return errors.New("algorithm zlib don't need order")
}

func (z *zlibAlgorithm) SetLitWidth(litWidth compressor.LitWidth) error {
	return errors.New("algorithm zlib don't need litWidth")
}

func init() {
	compressor.Register("zlib", &zlibAlgorithm{})
}
