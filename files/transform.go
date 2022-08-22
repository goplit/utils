package files

import (
	"errors"
	"fmt"
	"github.com/goplit/utils/utilerror"
	"net/url"
	"strconv"
	"strings"
)

type TransformStringMapProduct struct {
	product  map[string]TransformStringProduct
	errChain []error
}

func (tsm TransformStringMapProduct) GetKey(k string) TransformStringProduct {
	if item, found := tsm.product[k]; found {
		return item
	}
	err := fmt.Errorf(utilerror.NotFoundError)
	var errArr []error
	if len(tsm.errChain) > 0 {
		errArr = make([]error, len(tsm.errChain))
		copy(errArr, tsm.errChain)
		errArr = append(errArr, err)
	} else {
		errArr = []error{err}
	}
	return TransformStringProduct{errChain: errArr, err: err}
}

type TransformStringProduct struct {
	product  string
	errChain []error
	err      error
}

func (trp TransformStringProduct) IsError() bool {
	if len(trp.errChain) > 0 || trp.err != nil {
		return true
	}
	return false
}

func (trp TransformStringProduct) AsError() error {
	if len(trp.errChain) > 0 {
		errStr := ""
		for index, err := range trp.errChain {
			errStr += fmt.Errorf("[%d] %w", index, err).Error()
		}
		if trp.err != nil {
			errStr += fmt.Errorf("[%d] %w", len(trp.errChain), trp.err).Error()
		}
		return fmt.Errorf("%s", errStr)
	}
	return nil
}

func (trp TransformStringProduct) AsCSVList() StringArrayValue {
	splits := strings.Split(trp.product, ",")
	if len(splits) == 0 {
		return StringArrayValue{
			product:  splits,
			err:      errors.New(utilerror.EmptyError),
			errChain: trp.errChain,
		}
	}
	return StringArrayValue{
		product:  splits,
		err:      nil,
		errChain: trp.errChain,
	}
}

func (trp TransformStringProduct) AsUrl() UrlValue {
	u, err := url.Parse(trp.product)
	if err != nil {
		trp.err = err
		return UrlValue{
			product:  nil,
			err:      err,
			errChain: trp.errChain,
		}
	}
	return UrlValue{
		product:  u,
		err:      nil,
		errChain: trp.errChain,
	}
}

func (trp TransformStringProduct) AsJSON() JsonObjectValue {
	return JsonObjectValue{
		product:  nil,
		err:      fmt.Errorf("parsing JSON not implemented yet"),
		errChain: trp.errChain,
	}
}

func (trp TransformStringProduct) AsBoolean() BoolValue {
	val, err := strconv.ParseBool(trp.product)
	if err != nil {
		return BoolValue{
			product:  false,
			err:      err,
			errChain: trp.errChain,
		}
	}
	return BoolValue{
		product:  val,
		err:      nil,
		errChain: trp.errChain,
	}
}

func (trp TransformStringProduct) AsInt() IntValue {
	val, err := strconv.ParseInt(trp.product, 10, 64)
	if err != nil {
		return IntValue{
			product:  0,
			err:      err,
			errChain: trp.errChain,
		}
	}
	return IntValue{
		product:  val,
		err:      nil,
		errChain: trp.errChain,
	}
}
