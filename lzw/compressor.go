package lzw

import (
	"compress/lzw"
	"errors"
	"github.com/lsilencej/compressor"
	"io"
)

type lzwAlgorithm struct {
	order    compressor.Order
	litWidth compressor.LitWidth
}

type lzwCompressor struct {
	writer io.WriteCloser
}

type lzwUnCompressor struct {
	reader io.ReadCloser
}

func (l lzwCompressor) Write(data []byte) (int, error) {
	return l.writer.Write(data)
}

func (l lzwCompressor) Close() error {
	return l.writer.Close()
}

func (l lzwUnCompressor) Read(data []byte) (int, error) {
	return l.reader.Read(data)
}

func (l lzwUnCompressor) Close() error {
	return l.reader.Close()
}

func (l *lzwAlgorithm) NewAlgorithm() compressor.Algorithm {
	return &lzwAlgorithm{}
}

func (l *lzwAlgorithm) NewCompressor(w io.Writer) (compressor.Compressor, error) {
	return &lzwCompressor{
		writer: lzw.NewWriter(w, lzw.Order(l.order), int(l.litWidth)),
	}, nil
}

func (l *lzwAlgorithm) NewUnCompressor(r io.Reader) (compressor.UnCompressor, error) {
	return &lzwUnCompressor{
		reader: lzw.NewReader(r, lzw.Order(l.order), int(l.litWidth)),
	}, nil
}

func (l *lzwAlgorithm) Compress(data []byte) ([]byte, error) {
	return compressor.Compress(l, data)
}

func (l *lzwAlgorithm) UnCompress(data []byte) ([]byte, error) {
	return compressor.UnCompress(l, data)
}

func (l *lzwAlgorithm) SetLevel(level compressor.Level) error {
	return errors.New("algorithm lzw don't need level")
}

func (l *lzwAlgorithm) SetOrder(order compressor.Order) error {
	l.order = order
	return nil
}

func (l *lzwAlgorithm) SetLitWidth(litWidth compressor.LitWidth) error {
	l.litWidth = litWidth
	return nil
}

func init() {
	compressor.Register("lzw", &lzwAlgorithm{})
}
