package validate
import "github.com/go-playground/validator/v10"

var instanceValidate = validator.New()

func Struct(obj interface{}) error{
	return instanceValidate.Struct(obj)
}