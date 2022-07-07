package iselect

// select 的主要功能：防止协程永久阻塞。通常跟channel一起使用

// default 防止永久性阻塞，并不会一直执行default(s2)
// 空的select会引发永久阻塞
// select 的case是随机执行(s3)
