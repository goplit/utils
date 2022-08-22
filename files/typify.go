package files

import (
	"fmt"
	"strings"
)

type TypifyProduct struct {
	fileProduct FileProduct
	errChain    []error
	err         error
}

func (tp TypifyProduct) IsError() bool {
	if len(tp.errChain) > 0 {
		return true
	}
	return false
}

func (tp TypifyProduct) AsError() error {
	if len(tp.errChain) > 0 {
		errStr := ""
		for index, err := range tp.errChain {
			errStr += fmt.Errorf("[%d] %w", index, err).Error()
		}
		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (tp TypifyProduct) EquationLineTransforms() TransformStringMapProduct {
	out := TransformStringMapProduct{
		product:  make(map[string]TransformStringProduct),
		errChain: tp.errChain,
	}
	if len(tp.fileProduct.bytes) == 0 {
		out.errChain = tp.errChain
		return out
	}
	lines := strings.Split(string(tp.fileProduct.bytes), "\n")
	for _, line := range lines {
		equation := strings.Split(line, "=")
		if len(equation) >= 2 {
			key := equation[0]
			val := strings.Join(equation[1:], "=")
			out.product[key] = TransformStringProduct{
				product: val,
			}
		}
	}
	return out
}

func (tp TypifyProduct) EquationLineMap() (map[string]string, error) {
	return nil, nil
}

func (tp TypifyProduct) YamlMap() (map[string]interface{}, error) {
	return nil, nil
}
