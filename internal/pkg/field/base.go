package field

import "google.golang.org/protobuf/types/known/structpb"

type Type interface {
	Name() string
	Validate(in *structpb.Value) bool
	ToString(in *structpb.Value) string
	ToValue(in string) *structpb.Value
}

type field map[string]Type

var ins = make(field)

func register(key string, tp Type) {
	ins[key] = tp
}

type Field interface {
	GetType(tp string) Type
	GetTypes() map[string]Type
}

func init() {
	register("bool", &boolType{})
	register("datetime", &datetimeType{})
	register("email", &emailType{})
	register("id_card", &idCardType{})
	register("number", &numberType{})
	register("phone", &phoneType{})
	register("text", &textType{})
}

func (f field) GetType(tp string) Type {
	return f[tp]
}

func (f field) GetTypes() map[string]Type {
	return f
}

func New() Field {
	return ins
}
