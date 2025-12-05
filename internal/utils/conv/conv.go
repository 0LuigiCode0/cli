package conv

type num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func ArrNum[to ~[]t2, from ~[]t1, t1, t2 num](in from) (out to) {
	l := len(in)
	out = make(to, l)
	for i, v := range in {
		out[i] = t2(v)
	}
	return
}
