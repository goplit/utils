package files

import "io/ioutil"

func ReadFileToProduct(path string) FileProduct {
	product := FileProduct{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		product.err = err
		return product
	}
	product.bytes = file
	return product
}
