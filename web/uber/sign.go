package uber

import "time"

// 避免可变的全局变量

// Bad
var _timeNow = time.Now

func sign(msg string) string {
	now := _timeNow()
	return signWithTime(msg, now)
}

// Good
type signer struct {
	now func() time.Time
}

func newSigner() *signer {
	return &signer{now: time.Now}
}
func (s *signer) Sign(msg string) string {
	now := s.now()
	return signWithTime(msg, now)
}

func signWithTime(msg string, now time.Time) string {
	// do something
	return ""
}
