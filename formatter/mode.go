package formatter

type Mode uint

const (
	ALL Mode = iota
	UNDONE
	DONE
	NONE
)

func NewMode(done, undone bool) Mode {
	if done && undone {
		return NONE
	}
	if done && !undone {
		return DONE
	}
	if !done && undone {
		return UNDONE
	}
	if !done && !undone {
		return ALL
	}
	return ALL
}
