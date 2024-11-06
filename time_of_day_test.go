package TimeOfDay

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		inputTime time.Time
		want      TimeOfDay
	}{
		{
			name:      "Morning",
			inputTime: time.Date(2022, time.January, 1, 8, 0, 0, 0, time.UTC),
			want:      TimeOfDay{Day, AM, Morning},
		},
		{
			name:      "Afternoon",
			inputTime: time.Date(2022, time.January, 1, 14, 0, 0, 0, time.UTC),
			want:      TimeOfDay{Day, PM, Afternoon},
		},
		{
			name:      "Evening",
			inputTime: time.Date(2022, time.January, 1, 20, 0, 0, 0, time.UTC),
			want:      TimeOfDay{Day, PM, Evening},
		},
		{
			name:      "Night",
			inputTime: time.Date(2022, time.January, 1, 2, 0, 0, 0, time.UTC),
			want:      TimeOfDay{Night, AM, Night},
		},
	}
	for _, testTime := range tests {
		t.Run(testTime.name, func(t *testing.T) {
			got := Parse(testTime.inputTime)
			if got != testTime.want {
				t.Errorf("Parse() = %v, want %v", got, testTime.want)
			}
		})
	}
}
