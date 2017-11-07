package scanner

import (
	"bufio"
	"io"
	"os"
)

type Scanner struct {
	r     *bufio.Reader
	token []byte
	err   error
}

func New(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

func OpenFile(name string) *Scanner {
	f, err := os.Open(name)
	return &Scanner{
		r:   bufio.NewReader(f),
		err: err,
	}
}

func (s *Scanner) Scan() bool {
	// check s.openFile()
	if s.err != nil {
		return false
	}

	s.token, s.err = s.r.ReadSlice('\n')
	if s.err != nil {
		if s.err == io.EOF {
			s.err = nil
		}
		return false
	}
	s.token = s.token[:len(s.token)-1]
	return true
}

func (s *Scanner) Bytes() []byte {
	return s.token
}

func (s *Scanner) Text() string {
	return string(s.token)
}

func (s *Scanner) Err() error {
	return s.err
}
