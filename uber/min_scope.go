package uber

import "io/ioutil"

// 尽量缩小变量的作用范围

func badScope() error {
	err := ioutil.WriteFile("", []byte("bad"), 0644)
	if err != nil {
		return err
	}
	return nil
}

func goodScope() error {
	// err 只作用在 if 范围内
	if err := ioutil.WriteFile("", []byte(`Good`), 0644); err != nil {
		return err
	}
	return nil
}
