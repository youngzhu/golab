package paralle

import "runtime"

type Vector []float64

func (v Vector) Op(f float64) float64 {
	return 0
}

// DoSome apply the operation to v[i...n-1]
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // signal that this piece is done
}

var numCPU = runtime.NumCPU()

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU) // buffering optional but sensible
	// 这是个神奇的分配法
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	// drain the channel
	for i := 0; i < numCPU; i++ {
		<-c // wait for one task to complete
	}
	// all done
}
