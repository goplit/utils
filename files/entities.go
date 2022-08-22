package files

type FileProduct struct {
	err   error
	bytes []byte
}

func (p FileProduct) IsError() bool {
	if p.err != nil {
		return true
	}
	return false
}

func (p FileProduct) Bytes() []byte {
	return p.bytes
}

func (p FileProduct) As() TypifyProduct {
	var errChain []error
	if p.err != nil {
		errChain = append(errChain, p.err)
	}
	return TypifyProduct{
		fileProduct: p,
		errChain:    errChain,
	}
}

func (p FileProduct) AsError() error {
	return p.err
}
