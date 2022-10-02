package VignereCipher

import (
	"testing"
)

func TestVignereCipher(t *testing.T) {
	type args struct {
		plainText string
		key       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Plain Text or key contains non-uppercase letters",
			args: args{
				plainText: "Iloveyou",
				key:       "Yero",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				plainText: "ILOVEYOU",
				key:       "YERO",
			},
			want:    "GPFJCCFI",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vignereCipher(tt.args.plainText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("vignereCipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("vignereCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVignereDecipher(t *testing.T) {
	type args struct {
		cipherText string
		key        string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Cipher Text or key contains non-uppercase letters",
			args: args{
				cipherText: "LXFOPVEFRNHR",
				key:        "yero",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				cipherText: "GPFJCCFI",
				key:        "YERO",
			},
			want:    "ILOVEYOU",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vignereDecipher(tt.args.cipherText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("vignereDecipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("vignereDecipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
