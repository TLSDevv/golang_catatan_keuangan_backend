package service

// type IUserService interface {
// 	Create(ctx context.Context, user dto.UserRequest) error
// 	Update(ctx context.Context, id int, user dto.UserRequest) error
// 	Delete(ctx context.Context, id int) error
// 	FindById(ctx context.Context, id int) (dto.User, error)
// 	List(ctx context.Context) ([]dto.User, error)
// }

// type UserService struct {
// 	UserRepository repository.IUserRepository
// 	DB             *sql.DB
// }

// func NewUserService(userRepo repository.IUserRepository, db *sql.DB) IUserService {
// 	return UserService{
// 		UserRepository: userRepo,
// 		DB:             db,
// 	}
// }

// func (service UserService) Create(ctx context.Context, userDto dto.UserRequest) error {
// 	tx, err := service.DB.Begin()

// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}

// 		tx.Commit()
// 	}()

// 	user := model.User{
// 		Username: userDto.Username,
// 		Email:    userDto.Email,
// 		Password: userDto.Password,
// 		Fullname: userDto.Fullname,
// 	}

// 	err = service.UserRepository.Create(ctx, tx, user)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (service UserService) Update(ctx context.Context, id int, userDto dto.UserRequest) error {
// 	tx, err := service.DB.Begin()

// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}

// 		tx.Commit()
// 	}()

// 	user, err := service.UserRepository.FindById(ctx, tx, id)

// 	if err != nil {
// 		return err
// 	}

// 	if user.ID == 0 {
// 		return errors.New("User Not Found")
// 	}

// 	user.Update(userDto)

// 	err = service.UserRepository.Update(ctx, tx, user)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (service UserService) Delete(ctx context.Context, id int) error {
// 	tx, err := service.DB.Begin()

// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}

// 		tx.Commit()
// 	}()

// 	user, err := service.UserRepository.FindById(ctx, tx, id)

// 	if err != nil {
// 		return err
// 	}

// 	if user.ID == 0 {
// 		return errors.New("User Not Found")
// 	}

// 	err = service.UserRepository.Delete(ctx, tx, user)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (service UserService) FindById(ctx context.Context, id int) (dto.User, error) {
// 	tx, err := service.DB.Begin()

// 	if err != nil {
// 		return dto.User{}, nil
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}

// 		tx.Commit()
// 	}()

// 	user, err := service.UserRepository.FindById(ctx, tx, id)

// 	if err != nil {
// 		return dto.User{}, err
// 	}

// 	if user.ID == 0 {
// 		return dto.User{}, errors.New("User Not Found")
// 	}

// 	userDto := dto.User{
// 		ID:       user.ID,
// 		Username: user.Username,
// 		Email:    user.Email,
// 		Fullname: user.Fullname,
// 	}

// 	return userDto, nil
// }

// func (service UserService) List(ctx context.Context) ([]dto.User, error) {
// 	tx, err := service.DB.Begin()

// 	if err != nil {
// 		return []dto.User{}, nil
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}

// 		tx.Commit()
// 	}()

// 	users, err := service.UserRepository.List(ctx, tx)

// 	return helper.UsersToDTO(users), nil
// }
