package factories

import (
	services "github.com/matheusF23/pecunia-go/internal/data/services/user"
	"github.com/matheusF23/pecunia-go/internal/infra/repositories/database"
	"github.com/matheusF23/pecunia-go/internal/presentation/contracts"
	controllers "github.com/matheusF23/pecunia-go/internal/presentation/controllers/user"
)

func MakeCreateUserController() contracts.Controller {
	conn, err := database.OpenConnection()
	if err != nil {
		panic(err)
	}
	userRepository := database.NewUserRepository(conn)
	createUser := services.NewCreateUserService(userRepository)
	return controllers.NewCreateUserController(createUser)
}
