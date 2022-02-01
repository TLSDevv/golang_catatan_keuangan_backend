package cmd

var (
	conf *Config
)

func init() {
	loadConfig()
	conf = initConfig()
	dbconf := loadConfigDB()
	db = initDatabase(dbconf)
}

func Execute() {
	defer db.Close()
	//register all server needs, db,repo, etc

	// userRepo := repository.NewUserRepository()
	// userService := service.NewUserService(userRepo, db)

	// api := handler.NewAPI(userService)
	// api.Start(conf.Host, conf.Port)
}
