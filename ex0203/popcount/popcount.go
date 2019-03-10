package popcount

// pc[i]はiのポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountをループによって書き直した実装
func PopCountByLoop(x uint64) int {
	var sum byte
	var i uint = 0
	for ; i < 8; i++ {
		sum = pc[byte(x>>(i*8))]
	}
	return int(sum)
}
