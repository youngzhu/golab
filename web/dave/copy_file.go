package dave

func CopyFile(to, from string) error {

	return nil
}

// 上面的方法，两个参数相同，可能会颠倒导致错用
// 改成下面的方法，就不会再被误用
// CopyFile 方法还可以改为小写将其隐藏

type Source string

func (s Source) CopyTo(to string) error {
	return CopyFile(to, string(s))
}
