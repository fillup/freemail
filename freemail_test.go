package freemail

import "testing"

func TestIsFreeDomain(t *testing.T) {
	tests := []struct {
		name       string
		domain     string
		additional []string
		want       bool
	}{
		{
			name:       "empty",
			domain:     "",
			additional: nil,
			want:       false,
		},
		{
			name:       "not in list",
			domain:     "123456.com",
			additional: nil,
			want:       false,
		},
		{
			name:       "in default list",
			domain:     "gmail.com",
			additional: nil,
			want:       true,
		},
		{
			name:       "in additional list",
			domain:     "123456.com",
			additional: []string{"123456.com"},
			want:       true,
		},
		{
			name:       "in list, given mixed case and padding",
			domain:     " GmAiL.COM ",
			additional: nil,
			want:       true,
		},
		{
			name:       "in additional list with mixed case and padding",
			domain:     " myDomain.com    ",
			additional: []string{" myDOMAIN.com"},
			want:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFreeDomain(tt.domain, tt.additional...); got != tt.want {
				t.Errorf("IsFreeDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFreeEmail(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		additional []string
		want       bool
		wantErr    bool
	}{
		{
			name:       "empty",
			email:      "",
			additional: nil,
			want:       false,
			wantErr:    true,
		},
		{
			name:       "invalid email",
			email:      "123456.com",
			additional: nil,
			want:       false,
			wantErr:    true,
		},
		{
			name:       "in default list",
			email:      "joe.schmoe@gmail.com",
			additional: nil,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "in additional list",
			email:      "first.last@123456.com",
			additional: []string{"123456.com"},
			want:       true,
			wantErr:    false,
		},
		{
			name:       "in list, given mixed case and padding",
			email:      " Joey.The.Schmoey@GmAiL.COM ",
			additional: nil,
			want:       true,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFreeEmail(tt.email, tt.additional...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFreeEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFreeEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
