package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	// should create a new stack
	s := NewStack()

	if s == nil {
		t.Errorf("The stack should be created when the constructor is called.")
	}
}

// *** Below are some table-driven tests.  They have some advantages over writing individual
// test cases, but some disadvantages as well:
//
// Pros:
//   - They allow you to concisely create new test cases that involve changing data quickly
//   - You can focus only on the arguments you pass and results you expect
//
// Cons:
//   - It is difficult to integrate exception cases to them (eg. cases where you may expect a panic)
//     where the assertions may need to be different.
//   - It is difficult to use this method for stateful functions whose behavior may change over multiple
//     executions (I can see an argument against writing functions like this) since each case is created
//     and executed in isolation

func Test_stack_Push(t *testing.T) {
	lotsOfItems := make([]interface{}, 100)

	for i := 0; i < 100; i++ {
		lotsOfItems[i] = i
	}

	type fields struct {
		elements []interface{}
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "should push a value when there is something in the stack",
			fields: fields{
				elements: []interface{}{
					"some string",
				},
			},
			args: args{
				val: "some other string",
			},
		}, {
			name: "should push a value to an empty stack",
			fields: fields{
				elements: []interface{}{},
			},
			args: args{
				val: "some new string",
			},
		}, {
			name: "should push even with many values in the stack",
			fields: fields{
				elements: lotsOfItems,
			},
			args: args{
				val: "some new item",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				elements: tt.fields.elements,
			}
			s.Push(tt.args.val)
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "should pop a single element",
			fields: fields{
				elements: []interface{}{
					"some element",
				},
			},
			want: "some element",
		}, {
			name: "should pop the LAST element",
			fields: fields{
				elements: []interface{}{
					"some element",
					"some other element",
					"yet another element",
				},
			},
			want: "yet another element",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				elements: tt.fields.elements,
			}
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
