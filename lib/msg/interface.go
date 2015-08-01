package msg

type Interface interface {
	Decode(data []byte) error
	Encode() ([]byte, error)
}
