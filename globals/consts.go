package globals

import "time"

const (
	JwtTokenExpiry = int64(time.Hour) * 12
	Dsn            = "host=localhost user=authadmin password=dev123 dbname=authapi port=5432 sslmode=disable"

	AlphabetLower      = "abcdefghijklmnopqrstuvwxyz"
	AlphabetUpper      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetAll        = AlphabetLower + AlphabetUpper
	Numerics           = "0123456789"
	Alphanumnerics     = AlphabetAll + Numerics
	Symbols            = "`~!@#$%^&*()-=_+[]{}\\|;:'\"/?.>,<"
	StandardCharacters = Alphanumnerics + Symbols
)
