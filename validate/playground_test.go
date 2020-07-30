package validate

import (
	"reflect"
	"testing"

	"github.com/go-playground/locales"
)

func TestNewValid(t *testing.T) {
	type args struct {
		translator locales.Translator
	}
	tests := []struct {
		name string
		args args
		want *Valid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValid(tt.args.translator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValid_NameVar(t *testing.T) {
	type args struct {
		name  string
		field interface{}
		tag   string
	}
	tests := []struct {
		name    string
		v       *Valid
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.NameVar(tt.args.name, tt.args.field, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("Valid.NameVar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValid_FirstError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		v       *Valid
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.FirstError(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Valid.FirstError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValid_Errors(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		v    *Valid
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Errors(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Valid.Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}
