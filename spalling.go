package spalling // import "github.com/jbowes/spalling"

import (
	"fmt"

	"golang.org/x/xerrors"
)

func Wrap(err error, msg string, skip uint) *Wrapper {
	return &Wrapper{*Seal(err, msg, skip+1)}
}

func Seal(err error, msg string, skip uint) *Sealer {
	return &Sealer{
		msg:   msg,
		err:   err,
		frame: xerrors.Caller(int(2 + skip)),
	}
}

type Sealer struct {
	msg   string
	err   error
	frame xerrors.Frame
}

func (e *Sealer) Error() string              { return fmt.Sprint(e) }
func (e *Sealer) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *Sealer) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

type Wrapper struct {
	Sealer
}

func (e *Wrapper) Unwrap() error { return e.err }
