package snappy

import (
	"errors"
	"github.com/golang/snappy"
	"github.com/lsilencej/compressor"
	"io"
)

type snappyAlgorithm struct {
}

type snappyCompressor struct {
	writer *snappy.Writer
}

type snappyUnCompressor struct {
	reader *snappy.Reader
}

func (s *snappyCompressor) Write(data []byte) (int, error) {
	return s.writer.Write(data)
}

func (s *snappyCompressor) Close() error {
	return s.writer.Close()
}

func (s *snappyUnCompressor) Read(data []byte) (int, error) {
	return s.reader.Read(data)
}

func (s *snappyUnCompressor) Close() error {
	return nil
}

func (s *snappyAlgorithm) NewAlgorithm() compressor.Algorithm {
	return &snappyAlgorithm{}
}

func (s *snappyAlgorithm) NewCompressor(w io.Writer) (compressor.Compressor, error) {
	return &snappyCompressor{
		writer: snappy.NewBufferedWriter(w),
	}, nil
}

func (s *snappyAlgorithm) NewUnCompressor(r io.Reader) (compressor.UnCompressor, error) {
	return &snappyUnCompressor{
		reader: snappy.NewReader(r),
	}, nil
}

func (s *snappyAlgorithm) Compress(data []byte) ([]byte, error) {
	return compressor.Compress(s, data)
}

func (s *snappyAlgorithm) UnCompress(data []byte) ([]byte, error) {
	return compressor.UnCompress(s, data)
}

func (s *snappyAlgorithm) SetLevel(level compressor.Level) error {
	return errors.New("algorithm snappy don't need level")
}

func (s *snappyAlgorithm) SetOrder(order compressor.Order) error {
	return errors.New("algorithm snappy don't need order")
}

func (s *snappyAlgorithm) SetLitWidth(litWidth compressor.LitWidth) error {
	return errors.New("algorithm snappy don't need litWidth")
}

func init() {
	compressor.Register("snappy", &snappyAlgorithm{})
}
