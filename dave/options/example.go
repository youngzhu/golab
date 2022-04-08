package options

import "github.com/pkg/term"

/* Term (original) */
func foo() {
	t, err := term.Open("/dev/ttyUSB0")
	// handle error
	err = t.SetSpeed(115200)
	// handle error
	err = t.SetRaw()
	// handle error
	// ...

	handle(err)
}

func handle(err error) {

}

/* Term (improved) */
func OpenTerm(dev string, options ...func(t *term.Term) error) (*term.Term, error) {
	t, err := term.Open(dev)
	// handle err
	handle(err)

	for _, opt := range options {
		err = opt(t)
		// handle err
	}

	return t, nil
}
func bar() {
	// just open the terminal
	t, _ := term.Open("/dev/ttyUSB0")
	t.Flush()

	// open at 115200 baud in raw mode
	t2, _ := term.Open("/dev/ttyUSB0",
		Speed(115200),
		RewMode)
	t2.Flush()
}

// RewMode places the terminal into raw mode.
func RewMode(t *term.Term) error {
	return t.SetRaw()
}

// Speed sets the baud rate option for the terminal
// Speed自己需要一个参数，所以它返回了一个func
func Speed(baud int) func(tt *term.Term) error {
	return func(t *term.Term) error {
		return t.SetSpeed(baud)
	}
}
