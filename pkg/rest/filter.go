package rest

type Filters []Filter

type Filter struct {
	Field  string      `json:"field" validate:""`
	Symbol string      `json:"symbol"`
	Val    interface{} `json:"val"`
	Left   interface{} `json:"left"`
	Right  interface{} `json:"right"`
}

func (f *Filter) ValidateField() {
	return
}

func (f *Filter) ValidateSymbol() {

}
