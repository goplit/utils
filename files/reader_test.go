package files

import (
	"testing"
)

func TestReadFileToProduct(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		args   args
		cmdSeq []string
	}{
		{
			"Test reading env file",
			args{path: "test-tools/example.env"},
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := ReadFileToProduct(tt.args.path); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ReadFileToProduct() = %v, want %v", got, tt.want)
			//}
			val := ReadFileToProduct(tt.args.path).As().EquationLineTransforms().GetKey("NO_PROXY").AsCSVList().Value()
			t.Logf("output: %v", val)
		})
	}
}
