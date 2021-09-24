package error

import "net/http"

var ErrInvalidGeneral = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "G001",
	Message:    "Kami tidak dapat memproses permintaan Anda",
}

var ErrInvalidInput = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0001",
	Message:    "Kami tidak dapat memproses permintaan Anda",
}

var ErrUnauthorized = APIError{
	StatusCode: http.StatusUnauthorized,
	Code:       "0002",
	Message:    "Anda tidak memiliki akses",
}

var ErrEmailNotFoundForgotPassword = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0003",
	Message:    "Email yang Anda masukkan belum terdaftar",
}

var ErrEmailAlreadyRegistered = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0004",
	Message:    "Akun sudah pernah didaftarkan",
}

var ErrEmailNotFoundLogin = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0005",
	Message:    "Email Anda belum terdaftar",
}

var ErrEmailPasswordInvalidLogin = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0005",
	Message:    "Email atau password Anda salah",
}

var ErrRequireSubscription = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0006",
	Message:    "Untuk mendaftar workshop ini silahkan berlangganan terlebih dahulu",
}

var ErrAccessArticleLimitReach = APIError{
	StatusCode: http.StatusBadRequest,
	Code:       "0007",
	Message:    "Maaf nih Pak, total akses ke artikel premium sudah mencapai batas, silakan subscribe untuk dapatkan akses tidak terbatas ke semua premium artikel",
}
