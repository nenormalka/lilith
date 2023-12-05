package methods

import (
	"reflect"
	"testing"
)

func TestArrayStructToArrayValuesByFieldInt(t *testing.T) {
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
		want   []int
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
			want:   []int{1, 111},
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
			want:   []int{8, 19},
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
			want:   []int{1, 2},
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
			want:   []int{1, 3, 2},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayStructToArrayValuesByField[[]test, test, int](tt.arr, tt.fields)
			if len(got) == 0 && len(tt.want) == 0 {

			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStructToMapByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStructToArrayValuesByFieldString(t *testing.T) {
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
		want   []string
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
			want:   []string{"asda", "test"},
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
			want:   []string{"test", "ty"},
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
			want:   []string{"1", "2"},
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
			want:   []string{"1", "3", "2"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayStructToArrayValuesByField[[]test, test, string](tt.arr, tt.fields)
			if len(got) == 0 && len(tt.want) == 0 {

			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStructToMapByField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStructToArrayValuesByFieldStruct(t *testing.T) {
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
		want   []testInternal
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
			want: []testInternal{
				{
					P: 1,
				},
				{
					O: "test int",
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
			want: []testInternal{
				{
					B: true,
					O: "test",
					P: 8,
				},
				{
					B: false,
					O: "ty",
					P: 19,
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
			got := ArrayStructToArrayValuesByField[[]test, test, testInternal](tt.arr, tt.fields)
			if len(got) == 0 && len(tt.want) == 0 {

			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStructToMapByField() = %v, want %v", got, tt.want)
			}
		})
	}
}
