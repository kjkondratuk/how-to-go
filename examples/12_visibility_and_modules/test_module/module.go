package test_module

type TestStruct struct {
	// Notice that in go, capital letters indicate visibility to keep syntax concise

	privateValue string
	PublicValue  string
}

// SetPrivateValue : is an accessor method to update the private value stored in this struct
func (t *TestStruct) SetPrivateValue(s string) {
	t.privateValue = s
}
