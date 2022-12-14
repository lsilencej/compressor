package gzip

import (
	"compress/gzip"
	"errors"
	"github.com/lsilencej/compressor"
	"io"
)

type gzipAlgorithm struct {
	level compressor.Level
}

type gzipCompressor struct {
	writer *gzip.Writer
}

type gzipUnCompressor struct {
	reader *gzip.Reader
}

func (c *gzipCompressor) Write(data []byte) (int, error) {
	return c.writer.Write(data)
}

func (c *gzipCompressor) Close() error {
	return c.writer.Close()
}

func (uc *gzipUnCompressor) Read(data []byte) (int, error) {
	return uc.reader.Read(data)
}

func (uc *gzipUnCompressor) Close() error {
	return uc.reader.Close()
}

func (g *gzipAlgorithm) NewAlgorithm() compressor.Algorithm {
	return &gzipAlgorithm{}
}

func (g *gzipAlgorithm) NewCompressor(w io.Writer) (compressor.Compressor, error) {
	c := &gzipCompressor{}
	if g.level == 0 {
		c.writer = gzip.NewWriter(w)
	} else {
		var err error
		if c.writer, err = gzip.NewWriterLevel(w, int(g.level)); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (g *gzipAlgorithm) NewUnCompressor(r io.Reader) (compressor.UnCompressor, error) {
	uc := &gzipUnCompressor{}
	var err error
	if uc.reader, err = gzip.NewReader(r); err != nil {
		return nil, err
	}
	return uc, err
}

func (g *gzipAlgorithm) Compress(data []byte) ([]byte, error) {
	return compressor.Compress(g, data)
}

func (g *gzipAlgorithm) UnCompress(data []byte) ([]byte, error) {
	return compressor.UnCompress(g, data)
}

func (g *gzipAlgorithm) SetLevel(level compressor.Level) error {
	g.level = level
	return nil
}

func (g *gzipAlgorithm) SetOrder(order compressor.Order) error {
	return errors.New("algorithm gzip don't need order")
}

func (g *gzipAlgorithm) SetLitWidth(litWidth compressor.LitWidth) error {
	return errors.New("algorithm gzip don't need litWidth")
}

func init() {
	compressor.Register("gzip", &gzipAlgorithm{})
}
