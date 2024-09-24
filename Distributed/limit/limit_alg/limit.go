package limit_alg

type Limit interface {
	Allow() bool
}
