package VigenereCipher

import (
	"testing"
)

func TestVigenereCipher(t *testing.T) {
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
			got, err := vigenereCipher(tt.args.plainText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("vigenereCipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("vigenereCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVigenereDecipher(t *testing.T) {
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
			got, err := vigenereDecipher(tt.args.cipherText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("vigenereDecipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("vigenereDecipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
