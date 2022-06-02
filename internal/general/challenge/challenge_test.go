package challenge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateLeadingZeros(t *testing.T) {
	tests := []struct {
		name      string
		hash      string
		wantZeros int
	}{
		{
			name:      "0 leading zeros",
			hash:      "A0BC123",
			wantZeros: 0,
		},
		{
			name:      "hash with 00000",
			hash:      "00000A0BC123",
			wantZeros: 20,
		},
		{
			name:      "hash with 000100",
			hash:      "000100A0BC123",
			wantZeros: 15,
		},
		{
			name:      "hash with 000200",
			hash:      "000200A0BC123",
			wantZeros: 14,
		},
		{
			name:      "hash with 000300",
			hash:      "000300A0BC123",
			wantZeros: 14,
		},
		{
			name:      "hash with 000400",
			hash:      "000400A0BC123",
			wantZeros: 13,
		},
		{
			name:      "hash with 000500",
			hash:      "000500A0BC123",
			wantZeros: 13,
		},
		{
			name:      "hash with 000600",
			hash:      "000600A0BC123",
			wantZeros: 13,
		},
		{
			name:      "hash with 000700",
			hash:      "000700A0BC123",
			wantZeros: 13,
		},
		{
			name:      "hash with 000800",
			hash:      "000800A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000900",
			hash:      "000900A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000A00",
			hash:      "000A00A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000B00",
			hash:      "000B00A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000C00",
			hash:      "000C00A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000D00",
			hash:      "000D00A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000E00",
			hash:      "000E00A0BC123",
			wantZeros: 12,
		},
		{
			name:      "hash with 000F00",
			hash:      "000F00A0BC123",
			wantZeros: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotZeros := calculateLeadingZeros(tt.hash)

			assert.Equal(t, tt.wantZeros, gotZeros)
		})
	}
}

func TestChallenge_CheckAnswer(t *testing.T) {
	tests := []struct {
		name      string
		challenge string
		answer    string
		want      bool
	}{
		{
			name:      "answer is too short",
			challenge: "BpLnfgDsc2WD8F2qNfHK",
			answer:    "wJgC",
			want:      false,
		},
		{
			name:      "answer is too long",
			challenge: "BpLnfgDsc2WD8F2qNfHK",
			answer:    "2GHx49lT6LeGae",
			want:      false,
		},
		{
			name:      "sha2(challenge + answer) has less than 20 zeros",
			challenge: "BpLnfgDsc2WD8F2qNfHK",
			answer:    "AgpltYPRk6",
			want:      false,
		},
		{
			name:      "answer with unsupported symbols",
			challenge: "BpLnfgDsc2WD8F2qNfHK",
			answer:    "$�`��u<��L",
			want:      false,
		},
		{
			name:      "sha2(challenge + answer) has 20 zeros",
			challenge: "BpLnfgDsc2WD8F2qNfHK",
			answer:    "3LuC26cvLM",
			want:      true,
		},
		{
			name:      "sha2(challenge + answer) has more than 20 zeros",
			challenge: "BpLnfgDsc2WD8F2qNfHK",
			answer:    "M7OLXXt6V2",
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Challenge{
				Challenge: tt.challenge,
			}
			assert.Equalf(t, tt.want, c.CheckAnswer(tt.answer), "CheckAnswer(%v)", tt.answer)
		})
	}
}
