package controllers

import (
	"database/sql"
	"time"

	"demo2501/backend/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT secret key
var jwtSecret = []byte("your_secret_key")

// Xử lý đăng ký
func Register(ctx iris.Context) {
	var user User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Kết nối DB
	db := config.ConnectDB()
	defer db.Close()

	// Kiểm tra username đã tồn tại chưa
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&exists)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi kiểm tra username"})
		return
	}
	if exists {
		ctx.StatusCode(iris.StatusConflict)
		ctx.JSON(iris.Map{"error": "Username đã tồn tại"})
		return
	}

	// Thêm người dùng mới
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi tạo người dùng"})
		return
	}

	ctx.JSON(iris.Map{"message": "Đăng ký thành công"})
}

// Xử lý đăng nhập
func Login(ctx iris.Context) {
	var user User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Kết nối DB
	db := config.ConnectDB()
	defer db.Close()

	// Kiểm tra username và password
	var id int
	err = db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", user.Username, user.Password).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "Sai username hoặc password"})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "Lỗi xác thực"})
		}
		return
	}

	// Tạo JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi tạo token"})
		return
	}

	ctx.JSON(iris.Map{"token": tokenString})
}
