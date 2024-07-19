package User

type USERS struct {
	ID       int
	Name     string
	Email    string
	CreatAt  string
	UpdateAt string
}

func UsersData() []USERS {

	users := []USERS{
		{Name: "Islan", Email: "islan@teste.com", CreatAt: "26/05/1997", UpdateAt: "10/08/2024"},
		{Name: "Vinicius", Email: "vini@teste.com", CreatAt: "21/05/1997", UpdateAt: "10/08/2024"},
	}

	finalUsers := []USERS{}

	for index, user := range users {
		var currentUser USERS

		currentUser.ID = int(index) + 1
		currentUser.Name = user.Name
		currentUser.Email = user.Email
		currentUser.CreatAt = user.CreatAt
		currentUser.UpdateAt = user.UpdateAt

		finalUsers = append(finalUsers, currentUser)
	}

	return finalUsers
}
