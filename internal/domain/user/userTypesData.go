package user

type USERS struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreateAt  string `json:"createAt"`
	UpdatedAt string `json:"updateAt"`
}

type USER_STATUS_UPDATE_INPUT struct {
	UserID          int    `json:"userid"`
	Status          string `json:"status"`
	CurrentUserType string `json:"currentusertype"`
}

// dados mockados para teste antes de ser com o banco, apenas para testar a rota de GET
func UsersData() []USERS {

	users := []USERS{
		{Name: "Islan", Email: "islan@teste.com", CreateAt: "26/05/1997", UpdatedAt: "10/08/2024"},
		{Name: "Vinicius", Email: "vini@teste.com", CreateAt: "21/05/1997", UpdatedAt: "10/08/2024"},
	}

	finalUsers := []USERS{}

	for index, user := range users {
		var currentUser USERS

		currentUser.ID = int(index) + 1
		currentUser.Name = user.Name
		currentUser.Email = user.Email
		currentUser.CreateAt = user.CreateAt
		currentUser.UpdatedAt = user.UpdatedAt

		finalUsers = append(finalUsers, currentUser)
	}

	return finalUsers
}
