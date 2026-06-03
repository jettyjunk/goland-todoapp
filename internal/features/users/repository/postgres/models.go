package users_postgres_repository

import "github.com/jettyjunk/goland-todoapp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	usersDomains := make([]domain.User, len(users))

	for index, user := range users {
		usersDomains[index] = domain.NewUser(user.ID, user.Version, user.FullName, user.PhoneNumber)
	}

	return usersDomains
}
