package users_postgres_repository

import "github.com/jettyjunk/goland-todoapp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(usersModels []UserModel) []domain.User {
	usersDomains := make([]domain.User, len(usersModels))

	for index, user := range usersModels {
		usersDomains[index] = domain.NewUser(user.ID, user.Version, user.FullName, user.PhoneNumber)
	}

	return usersDomains
}
