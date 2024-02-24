package field

import (
	"google.golang.org/protobuf/types/known/structpb"
)

type textType struct {
}

func (t textType) Name() string {
	return "文本"
}

func (t textType) Validate(_ *structpb.Value) bool {
	return true
}

func (t textType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (t textType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
