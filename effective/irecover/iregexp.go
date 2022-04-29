package irecover

import "regexp"

type Regexp regexp.Regexp

// Error is the type of  parse error
// it satisfies the error interface
type Error string

func (e Error) Error() string {
	return string(e)
}

// error is a method of *Regexp that reports parsing errors
// by panicking with an Error
func (r *Regexp) error(err string) {
	panic(Error(err))
}

func (r *Regexp) doParse(str string) *Regexp {
	return nil
}

// Compile
// 有用的设计模式：不将panic暴露给客户端
func Compile(str string) (regexp *Regexp, err error) {
	regexp = new(Regexp)
	// doParse will panic if there is a parse error
	defer func() {
		if e := recover(); e != nil {
			// defer 可以改变命名的返回变量的值
			regexp = nil    // clear return value
			err = e.(Error) // will re-panic if not a parse error
		}
	}()
	return regexp.doParse(str), nil
}
