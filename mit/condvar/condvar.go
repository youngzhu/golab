package condvar

// condition var 的通用模式

/*
mu.Lock()
...
cond.Broadcast()
mu.Unlock()

mu.Lock()
for condition {
	cond.Wait()
}
mu.Unlock()
*/
