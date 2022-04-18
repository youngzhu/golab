package channel

// 向nil chan传值，永远阻塞
// 从nil chan取值，永远阻塞
// 向关闭的chan传值，panic
// 从关闭的chan取值，会立即得到相应类型的空值
