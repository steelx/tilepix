package tilepix

import "testing"

func TestProperty_String(t *testing.T) {
	type fields struct {
		Name  string
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Basic string",
			fields: fields{
				Name:  "name p",
				Value: "value p",
			},
			want: "Property{name p: value p}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Property{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
