package japanese

type Wareki interface {
	Resolve() bool
	GetGengo() string
	GetYear() int
	GetMonth() int
	GetDay() int
	ToString() string
}
