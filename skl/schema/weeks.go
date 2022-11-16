package schema

type weeks uint32

// Check if the schema run at specify week[1-31]
func (w *weeks) Check(n int) bool {
	if n > 32 || n < 1 {
		return false
	}
	return *w&(1<<n) == 1<<n
}
