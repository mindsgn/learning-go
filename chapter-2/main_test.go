package main

import "testing"

func TestNumberAssignment(t *testing.T) {
	tests := []struct {
		name       string
		wantByte   byte
		wantSmallI int32
		wantBigI   uint64
	}{
		{"Number Assignment", 128, -2147483648, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotByte, gotSmallI, gotBigI := NumberAssignment()

			if gotByte != tt.wantByte {
				t.Errorf("NumberAssignment() got int %d, want %d", gotByte, tt.wantByte)
			}

			if gotSmallI != tt.wantSmallI {
				t.Errorf("NumberAssignment() got int %d, want %d", gotSmallI, tt.wantSmallI)
			}

			if gotBigI != tt.wantBigI {
				t.Errorf("NumberAssignment() got int %d, want %d", gotBigI, tt.wantBigI)
			}
		})
	}
}

func TestAssignConstant(t *testing.T) {
	tests := []struct {
		name  string
		wantF float32
		wantI int
	}{
		{"Constant Assignment", 10.0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, gotI := AssignConstant()

			if gotI != tt.wantI {
				t.Errorf("AssignConstant() got int %d, want %d", gotI, tt.wantI)
			}

			if gotF != tt.wantF {
				t.Errorf("AssignConstant() got float %f, want %f", gotF, tt.wantF)
			}
		})
	}
}

func TestFloatAssign(t *testing.T) {
	tests := []struct {
		name  string
		wantF float32
		wantI int
	}{
		{"Float Assignment", 20.0, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, gotI := AssignFloat()

			if gotI != tt.wantI {
				t.Errorf("AssignFloat() got int %d, want %d", gotI, tt.wantI)
			}

			if gotF != tt.wantF {
				t.Errorf("AssignFloat got float %f, want %f", gotF, tt.wantF)
			}
		})

	}
}
