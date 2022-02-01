package handler

// type API struct {
// 	userService service.IUserService
// }

// func NewAPI(
// 	usService service.IUserService,
// ) *API {
// 	return &API{
// 		userService: usService,
// 	}
// }

// func (a API) Start(host, port string) {
// 	r := mux.NewRouter()

// 	fmt.Println(host, port)

// 	server := http.Server{
// 		Addr:    endPoint(host, port),
// 		Handler: r,
// 	}

// 	apiRoute := r.PathPrefix("/api/v1").Subrouter()

// 	NewUserHandler(apiRoute, a.userService)

// 	fmt.Printf("Listening %s to port %s", host, port)
// 	err := server.ListenAndServe()

// 	if err != nil {
// 		logrus.Error("error Listen serve ", err)
// 		return
// 	}
// }

// func endPoint(host, port string) string {
// 	return fmt.Sprintf("%s:%s", host, port)
// }
