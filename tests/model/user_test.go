package model

import (
	"database/sql"
	"fmt"
	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pioniry/controllers"
	"strings"
	"testing"
	"time"
)

var jsonData = `{"Nip": "19203394", "Firstname": "Gamal", "Lastname": "Akbar", "Email": "gamaltn@gmail.com", "Image": "a.jpg", "Password": "23456kjiuadiaod90276370q3098", "Token": "2345609","IsActive": 0}`

func init() {
	txdb.Register("mysqltx", "mysql", "root:@tcp(localhost:3307)/pioniry_test")
}

func dbPrepare() (*sql.DB, error) {
	cName := fmt.Sprintf("connection_%d", time.Now().UnixNano())
	op, err := sql.Open("mysqltx", cName)
	if err != nil {
		return nil, err
	}
	return op, nil
}
func TestCreateUser(t *testing.T) {
	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Percobaan123"), 8)

	con, err := dbPrepare()
	if err != nil {
		fmt.Println(err)
	}

	defer func(con *sql.DB) {
		err := con.Close()
		if err != nil {

		}
	}(con)

	//jsonData, _ := json.Marshal(userTest)
	//strJson := string(jsonData)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "localhost:3000/api/v1/user", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	uscon := controllers.UserController{}

	if assert.NoError(t, uscon.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
	}
	//assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetUser(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "localhost:3000/api/v1/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	uscon := controllers.UserController{}
	if assert.NoError(t, uscon.GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
	}
}
