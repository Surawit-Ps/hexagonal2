package service

import ("fmt"
"hexagonal2/core/ports"
"hexagonal2/core/entity")

type userService struct{
	repo ports.UserRepository
	
}

func NewUserService (repo ports.UserRepository) *userService {
	return &userService{repo: repo}
}

func (r userService) GetAllUser()([]entity.UserRes,error){
	user,err := r.repo.GetUsers()
	if err != nil{
		fmt.Print(err)
		return nil,err
	}
	var usRes []entity.UserRes
	for _,u := range user{
		usResp := entity.UserRes{
			Name: u.Name,
			LastName: u.LastName,
			Email: u.Email,
			Tel: u.Tel,
		}
		usRes = append(usRes,usResp)
	}
	return usRes,nil
}

func (r userService) GetUser(id string)(*entity.UserRes,error){
	user,err := r.repo.GetUser(id)
	if err != nil{
		fmt.Print(err)
		return nil,err
	}
	usResp := entity.UserRes{
			Name: user.Name,
			LastName: user.LastName,
			Email: user.Email,
			Tel: user.Tel,
		}
	return &usResp,nil
	
}

func (r userService) AddUser(p entity.User)error{
	err := r.repo.AddUser(p)
	if err != nil{
		fmt.Print(err)
		return err
	}
	return nil
}