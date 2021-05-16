package constant

const (
	INTERNAL_SERVER_ERROR = 1

	EMPTY_REQUEST = 1001
)

var ERROR_MAPPING = map[int]string{
	EMPTY_REQUEST: "empty request",
	INTERNAL_SERVER_ERROR: "internal error",
}
