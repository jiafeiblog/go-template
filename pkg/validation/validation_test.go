package validation

import (
	"fmt"
	"testing"
)

// TestVal
type TestVal struct {
	Title string  `validate:"required,gt=10,lte=50" json:"title"`
	Price float64 `validate:"required,min=5,max=10" form:"price"`
}

func TestMyValidate_Check(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				value: &TestVal{
					Title: "",
					Price: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "test",
			args: args{
				value: &TestVal{
					Title: "1",
					Price: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GlobalValidate.Check(tt.args.value)
			if err != nil {
				fmt.Println(err.Error())
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
