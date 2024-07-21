package User

type USERS struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

func UsersData() []USERS {

	users := []USERS{
		{Name: "Islan", Email: "islan@teste.com", CreateAt: "26/05/1997", UpdateAt: "10/08/2024"},
		{Name: "Vinicius", Email: "vini@teste.com", CreateAt: "21/05/1997", UpdateAt: "10/08/2024"},
	}

	finalUsers := []USERS{}

	for index, user := range users {
		var currentUser USERS

		currentUser.ID = int(index) + 1
		currentUser.Name = user.Name
		currentUser.Email = user.Email
		currentUser.CreateAt = user.CreateAt
		currentUser.UpdateAt = user.UpdateAt

		finalUsers = append(finalUsers, currentUser)
	}

	return finalUsers
}
