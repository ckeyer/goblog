package msg

type Interface interface {
	Decode(data []byte) error
	Encode() ([]byte, error)
}

func (i *Interface) String() string {
	bs := i.Encode()
	return string(bs)
}
func (i *Interface) Bytes() []byte {
	return i.Encode()
}
