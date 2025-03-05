package hash

type HashMethod string

const (
	MethodMd5    HashMethod = "md5"
	MethodBase64 HashMethod = "base64"
	MethodSHA256 HashMethod = "sha256"
	MethodSHA512 HashMethod = "sha512"
	MethodCrc64  HashMethod = "crc64"
)

type ErrorCode int

const (
	InternalServerError ErrorCode = 10201
)
