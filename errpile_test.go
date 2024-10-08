package errpile

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		template       string
		showLineNumber bool
		trackFullChain bool
	}
	tests := []struct {
		name string
		args args
		want *ErrorPile
	}{
		{
			name: "TestCreateNewErrorPile",
			args: args{
				template:       "%s->%v",
				showLineNumber: true,
				trackFullChain: false,
			},
			want: &ErrorPile{
				template:       "%s->%v",
				showLineNumber: true,
				trackFullChain: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.template, tt.args.showLineNumber, tt.args.trackFullChain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorPile_Error(t *testing.T) {
	type fields struct {
		template       string
		showLineNumber bool
		trackFullChain bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestErrorPileEmpty",
			fields: fields{
				template:       "%s->%v",
				showLineNumber: true,
				trackFullChain: false,
			},
			args: args{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "TestErrorPileWithErrorNoLineFullChain",
			fields: fields{
				template:       "%s->%v",
				showLineNumber: false,
				trackFullChain: true,
			},
			args: args{
				err: errors.New("TestErrorPileError"),
			},
			wantErr: true,
		},
		{
			name: "TestErrorPileWithErrorLineOriginalCallerContext",
			fields: fields{
				template:       "%s->%v",
				showLineNumber: true,
				trackFullChain: false,
			},
			args: args{
				err: errors.New("TestErrorPileError"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := &ErrorPile{
				template:       tt.fields.template,
				showLineNumber: tt.fields.showLineNumber,
				trackFullChain: tt.fields.trackFullChain,
			}
			if err := ep.Error(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("ErrorPile.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestErrorNil",
			args: args{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "TestErrorWithError",
			args: args{
				err: errors.New("TestErrorWithError"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
