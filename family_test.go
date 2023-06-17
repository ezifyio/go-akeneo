package goakeneo

import "testing"

func TestFamilyOp_CreateFamily(t *testing.T) {
	tests := []struct {
		name    string
		family  Family
		wantErr bool
	}{

		{
			name:    "CreateFamilyInvalid",
			family:  Family{},
			wantErr: true,
		},
		{
			name: "CreateFamily",
			family: Family{
				Code:             "test",
				AttributeAsLabel: "test_family",
			},
			wantErr: false,
		},
	}
	c := MockClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.Family.CreateFamily(tt.family); (err != nil) != tt.wantErr {
				t.Errorf("CreateFamily() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
