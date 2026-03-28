package ports
import "hexagonal2/core/entity"

type DogsRepository interface{
	GetDogs()([]entity.Dogs,error)
	GetADogs(string)(*entity.Dogs,error)
	AddDog(entity.Dogs,string)error
}
