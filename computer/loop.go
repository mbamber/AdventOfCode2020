package computer

import "errors"

// ErrInfiniteLoop is an error returned when running a computer results in an infinite loop
var ErrInfiniteLoop = errors.New("infinite loop detected")

// LoopDetector can detect an infinite loop in a computer
type LoopDetector interface {
	Detect(c *Computer) error
	Update(c *Computer)
}

type loopIpr struct {
	usedIprs []int
}

// NewIprLoopDetector returns a new loop detector that uses the
// previously used instruction pointers to determine if an infinite
// loop has been reached. No other state is considered
func NewIprLoopDetector() LoopDetector {
	return &loopIpr{}
}

func (l *loopIpr) Detect(c *Computer) error {
	for _, ipr := range l.usedIprs {
		if ipr == c.Ipr {
			return ErrInfiniteLoop
		}
	}
	return nil
}

func (l *loopIpr) Update(c *Computer) {
	l.usedIprs = append(l.usedIprs, c.Ipr)
}
