package cerrors

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWrapError(t *testing.T) {
	type args struct {
		e   *CError
		err error
	}
	tests := []struct {
		name string
		args args
		want *CError
	}{
		{
			name: "internal: server error: some error",
			args: args{
				e:   ErrInternal,
				err: errors.New("some error"),
			},
			want: &CError{
				Code: 500,
				Message: map[string][]any{
					"internal": {map[string]string{"server error": "some error"}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrapError(tt.args.e, tt.args.err); !cmp.Equal(err, tt.want) {
				t.Errorf(cmp.Diff(err, tt.want))
			}
		})
	}
}
