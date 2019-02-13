package persian

import "testing"

func TestToPersianDigits(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"empty", "", ""},
		{"zero", "0", "۰"},
		{"digits", "123", "۱۲۳"},
		{"mixed-digits", "12۳45", "۱۲۳۴۵"},
		{"mixed-letters", "12m34", "۱۲m۳۴"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPersianDigits(tt.str); got != tt.want {
				t.Errorf("ToPersianDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToASCIIDigits(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"empty", "", ""},
		{"zero", "۰", "0"},
		{"digits", "۱۲۳", "123"},
		{"mixed-digits", "۱۲3۴۵", "12345"},
		{"mixed-letters", "۱۲m۳۴", "12m34"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToASCIIDigits(tt.str); got != tt.want {
				t.Errorf("ToASCIIDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
