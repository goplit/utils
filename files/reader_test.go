package files

import (
	"reflect"
	"testing"
)

func TestReadFileToProduct(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name         string
		args         args
		cmdFuncTyped func(t *testing.T, product TypifyProduct)
	}{
		{
			"Test reading env file",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("NO_PROXY").AsCSVList()
				if val.Error() != nil {
					t.Errorf("cannot read env file NO_PROXY key due %v", val.AllErrors())
					return
				}
				if !reflect.DeepEqual(val.Value(), []string{".svc.cluster.local", "localhost", "127.0.0.1"}) {
					t.Errorf("cannot verify env file NO_PROXY key data due %v != %v", val.Value(), []string{".svc.cluster.local", "localhost", "127.0.0.1"})
					return
				}
				t.Logf("successfully read and verified NO_PROXY CSV key in env test file")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := ReadFileToProduct(tt.args.path); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ReadFileToProduct() = %v, want %v", got, tt.want)
			//}
			val := ReadFileToProduct(tt.args.path).As()
			tt.cmdFuncTyped(t, val)
		})
	}
}
