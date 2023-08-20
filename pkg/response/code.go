package response

type code int

const (
	SUCCESS code = iota + 2000
	CREATED
	ACCEPTED
	ALREADY
)

const (
	REDIRECT code = iota + 3000
	MOVED
	FOUND
)

const (
	BAD_REQUEST code = iota + 4000
	UNAUTHORIZED
	FORBIDDEN
	INVALID_PARAMS
	NOTFOUND
)

const (
	UNAVAILABLE code = iota + 5000
)
