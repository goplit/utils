package files

import (
	"github.com/goplit/utils/utilerror"
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
			"Test reading env file field as CSV",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("NO_PROXY").AsCSVList()
				if val.Error() != nil {
					t.Errorf("cannot read env file NO_PROXY key due %v", val.Error())
					return
				}
				if !reflect.DeepEqual(val.Value(), []string{".svc.cluster.local", "localhost", "127.0.0.1"}) {
					t.Errorf("cannot verify env file NO_PROXY key data due %v != %v", val.Value(), []string{".svc.cluster.local", "localhost", "127.0.0.1"})
					return
				}
				t.Logf("successfully read and verified NO_PROXY CSV key in env test file")
			},
		},
		{
			"Test reading env file field as URL",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("HTTPS_PROXY").AsUrl()
				if val.Error() != nil {
					t.Errorf("cannot read env file HTTPS_PROXY key due %v", val.Error())
					return
				}
				if !reflect.DeepEqual(val.Value().String(), "https://some.proxy.com:8033") {
					t.Errorf("cannot verify env file HTTPS_PROXY key data due %v != %v", val.Value().String(), "https://some.proxy.com:8033")
					return
				}
				t.Logf("successfully read and verified HTTPS_PROXY key in env test file")
			},
		},
		{
			"Test reading env file field as BOOL",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("BOOL_VALUE").AsBoolean()
				if val.Error() != nil {
					t.Errorf("cannot read env file BOOL_VALUE key due %v", val.Error())
					return
				}
				if !reflect.DeepEqual(val.Value(), true) {
					t.Errorf("cannot verify env file BOOL_VALUE key data due %v != %v", val.Value(), true)
					return
				}
				t.Logf("successfully read and verified BOOL_VALUE key in env test file")
			},
		},
		{
			"Test reading env file field as INT64",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("INT_VALUE").AsInt()
				if val.Error() != nil {
					t.Errorf("cannot read env file INT_VALUE key due %v", val.Error())
					return
				}
				if val.Value() != 5678909003 {
					t.Errorf("cannot verify env file INT_VALUE key data due %v != %v", val.Value(), 5678909003)
					return
				}
				t.Logf("successfully read and verified INT_VALUE key in env test file")
			},
		},
		{
			"Test reading env file field as JSON object",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("JSON_VALUE").AsJSON()
				if val.Error() != nil {
					t.Errorf("cannot read env file JSON_VALUE key due %v", val.Error())
					return
				}
				if !reflect.DeepEqual(val.Value(), map[string]interface{}{
					"just_string": "strval",
					"key": map[string]interface{}{
						"nested_array": []interface{}{"a", "b", "c"},
					}}) {
					t.Errorf("cannot verify env file JSON_VALUE key data due %v != %v", val.Value(), "{\"key\": {\"nested_array\": [\"a\", \"b\", \"c\"]}, \"just_string\": \"strval\"}")
					return
				}
				t.Logf("successfully read and verified JSON_VALUE key in env test file")
			},
		},
		{
			"Test failing env file field as JSON object",
			args{path: "test-tools/example.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("NOT_A_JSON").AsJSON()
				if val.Error() == nil {
					t.Errorf("should produce error on env file NOT_A_JSON key")
					return
				}
				t.Logf("successfully read and verified NOT_A_JSON key in env test file")
			},
		},
		{
			"Test reading non existent file",
			args{path: "test-tools/example1.env"},
			func(t *testing.T, product TypifyProduct) {
				val := product.EquationLineTransforms().GetKey("NOT_A_JSON").AsJSON()
				if val.Error() == nil {
					t.Errorf("should produce error for non existent file to read")
					return
				}
				if val.AsError(utilerror.NotFoundError) == nil {
					t.Errorf("as error should show not found error for faulty read")
				}
				t.Logf("%v", val.Error())
				t.Logf("successfully verified errors for non existent file chain")
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
