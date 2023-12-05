package methods

import (
	"reflect"
	"testing"
)

func TestArrayToMapByFieldInt(t *testing.T) {
	i1, i2, i3 := 1, 2, 3
	type (
		testDeeper struct {
			A *int
		}
		testInternal struct {
			B  bool
			O  string
			P  int
			d  int
			GF []*testDeeper
			BF map[int][]*testDeeper
		}

		test struct {
			A int
			S string
			Z testInternal
		}
	)

	for name, tt := range map[string]struct {
		arr    []test
		fields []string
		want   map[int]test
	}{
		"#1": {
			arr:    nil,
			fields: []string{"sadas"},
			want:   nil,
		},
		"#2": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
			},
			fields: []string{""},
			want:   nil,
		},
		"#3": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
				{
					A: 111,
					S: "test",
				},
			},
			fields: []string{"A"},
			want: map[int]test{
				1: {
					A: 1,
					S: "asda",
				},
				111: {
					A: 111,
					S: "test",
				},
			},
		},
		"#4": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "P"},
			want: map[int]test{
				8: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				19: {
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
		},
		"#5": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
			},
			fields: []string{"O"},
			want:   nil,
		},
		"#6": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"A", "P"},
			want:   nil,
		},
		"#7": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "W"},
			want:   nil,
		},
		"#8": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						d: 10,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "d"},
			want:   nil,
		},
		"#9": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						d: 10,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "O"},
			want:   nil,
		},
		"#10": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						GF: []*testDeeper{
							{
								A: &i1,
							},
							{
								A: &i2,
							},
							nil,
						},
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "GF", "A"},
			want: map[int]test{
				1: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						GF: []*testDeeper{
							{
								A: &i1,
							},
							{
								A: &i2,
							},
							nil,
						},
					},
				},
				2: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						GF: []*testDeeper{
							{
								A: &i1,
							},
							{
								A: &i2,
							},
							nil,
						},
					},
				},
			},
		},
		"#11": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[int][]*testDeeper{
							1: {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							2: {
								{
									A: &i2,
								},
							},
							3: nil,
						},
					},
				},
			},
			fields: []string{"Z", "BF", "A"},
			want: map[int]test{
				1: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[int][]*testDeeper{
							1: {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							2: {
								{
									A: &i2,
								},
							},
							3: nil,
						},
					},
				},
				2: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[int][]*testDeeper{
							1: {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							2: {
								{
									A: &i2,
								},
							},
							3: nil,
						},
					},
				},
				3: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[int][]*testDeeper{
							1: {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							2: {
								{
									A: &i2,
								},
							},
							3: nil,
						},
					},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayStructToMapByField[[]test, test, int](tt.arr, tt.fields)
			if len(got) == 0 && len(tt.want) == 0 {

			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStructToMapByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToMapByFieldString(t *testing.T) {
	i1, i2, i3 := "1", "2", "3"
	type (
		testDeeper struct {
			A *string
		}
		testInternal struct {
			B  bool
			O  string
			P  int
			d  int
			GF []*testDeeper
			BF map[string][]*testDeeper
		}

		test struct {
			A int
			S string
			Z testInternal
		}
	)

	for name, tt := range map[string]struct {
		arr    []test
		fields []string
		want   map[string]test
	}{
		"#1": {
			arr:    nil,
			fields: []string{"sadas"},
			want:   nil,
		},
		"#2": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
			},
			fields: []string{""},
			want:   nil,
		},
		"#3": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
				{
					A: 111,
					S: "test",
				},
			},
			fields: []string{"S"},
			want: map[string]test{
				"asda": {
					A: 1,
					S: "asda",
				},
				"test": {
					A: 111,
					S: "test",
				},
			},
		},
		"#4": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "O"},
			want: map[string]test{
				"test": {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				"ty": {
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
		},
		"#5": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
			},
			fields: []string{"O"},
			want:   nil,
		},
		"#6": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"A", "P"},
			want:   nil,
		},
		"#7": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "W"},
			want:   nil,
		},
		"#8": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						d: 10,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "d"},
			want:   nil,
		},
		"#9": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						d: 10,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "P"},
			want:   nil,
		},
		"#10": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						GF: []*testDeeper{
							{
								A: &i1,
							},
							{
								A: &i2,
							},
							nil,
						},
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "GF", "A"},
			want: map[string]test{
				"1": {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						GF: []*testDeeper{
							{
								A: &i1,
							},
							{
								A: &i2,
							},
							nil,
						},
					},
				},
				"2": {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						GF: []*testDeeper{
							{
								A: &i1,
							},
							{
								A: &i2,
							},
							nil,
						},
					},
				},
			},
		},
		"#11": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[string][]*testDeeper{
							"1": {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							"2": {
								{
									A: &i2,
								},
							},
							"3": nil,
						},
					},
				},
			},
			fields: []string{"Z", "BF", "A"},
			want: map[string]test{
				"1": {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[string][]*testDeeper{
							"1": {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							"2": {
								{
									A: &i2,
								},
							},
							"3": nil,
						},
					},
				},
				"2": {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[string][]*testDeeper{
							"1": {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							"2": {
								{
									A: &i2,
								},
							},
							"3": nil,
						},
					},
				},
				"3": {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						BF: map[string][]*testDeeper{
							"1": {
								{
									A: &i1,
								},
								{
									A: &i3,
								},
							},
							"2": {
								{
									A: &i2,
								},
							},
							"3": nil,
						},
					},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayStructToMapByField[[]test, test, string](tt.arr, tt.fields)
			if len(got) == 0 && len(tt.want) == 0 {

			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStructToMapByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToMapByFieldStruct(t *testing.T) {
	type (
		testInternal struct {
			B bool
			O string
			P int
			d int
		}

		test struct {
			A int
			S string
			Z testInternal
		}
	)

	for name, tt := range map[string]struct {
		arr    []test
		fields []string
		want   map[testInternal]test
	}{
		"#1": {
			arr:    nil,
			fields: []string{"sadas"},
			want:   nil,
		},
		"#2": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
			},
			fields: []string{""},
			want:   nil,
		},
		"#3": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						P: 1,
					},
				},
				{
					A: 111,
					S: "test",
					Z: testInternal{
						O: "test int",
					},
				},
			},
			fields: []string{"Z"},
			want: map[testInternal]test{
				testInternal{
					P: 1,
				}: {
					A: 1,
					S: "asda",
					Z: testInternal{
						P: 1,
					},
				},
				testInternal{
					O: "test int",
				}: {
					A: 111,
					S: "test",
					Z: testInternal{
						O: "test int",
					},
				},
			},
		},
		"#4": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z"},
			want: map[testInternal]test{
				testInternal{
					B: true,
					O: "test",
					P: 8,
				}: {
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				testInternal{
					B: false,
					O: "ty",
					P: 19,
				}: {
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
		},
		"#5": {
			arr: []test{
				{
					A: 1,
					S: "asda",
				},
			},
			fields: []string{"O"},
			want:   nil,
		},
		"#6": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "P"},
			want:   nil,
		},
		"#7": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "W"},
			want:   nil,
		},
		"#8": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						d: 10,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "d"},
			want:   nil,
		},
		"#9": {
			arr: []test{
				{
					A: 1,
					S: "asda",
					Z: testInternal{
						B: true,
						O: "test",
						P: 8,
						d: 10,
					},
				},
				{
					A: 15,
					S: "test",
					Z: testInternal{
						B: false,
						O: "ty",
						P: 19,
					},
				},
			},
			fields: []string{"Z", "P"},
			want:   nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayStructToMapByField[[]test, test, testInternal](tt.arr, tt.fields)
			if len(got) == 0 && len(tt.want) == 0 {

			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStructToMapByField() = %v, want %v", got, tt.want)
			}
		})
	}
}
