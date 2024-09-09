package json

import (
	"testing"
)

func TestIsJSONEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "same string and json",
			args: args{
				s1: "{}",
				s2: "{}",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "same json different string",
			args: args{
				s1: `{"score":100,"name":"frans"}`,
				s2: `{"name":"frans","score":100}`,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "different string and json",
			args: args{
				s1: `{"score":100,"name":"frans"}`,
				s2: `{"name":"frans","score":99}`,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "string but not json s1",
			args: args{
				s1: `"score":100,"name":"frans"`,
				s2: `{"name":"frans","score":99}`,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "string but not json s2",
			args: args{
				s1: `{"score":100,"name":"frans"}`,
				s2: `"name":"frans","score":99`,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsJSONEqual(tt.args.s1, tt.args.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsJSONEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsJSONEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_Decode(t *testing.T) {
	var res string
	u := &Decoder{}
	u.Decode([]byte{}, &res)
}

func TestDecoder_Unmarshal(t *testing.T) {
	var res string
	u := &Decoder{}
	u.Unmarshal([]byte{}, &res)

}

func TestDecoder_Marshal(t *testing.T) {
	var res string
	u := &Decoder{}
	u.Marshal(res)
}
