package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// used in Init() as sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User some comment
type User struct {
	gorm.Model
	Username string
	Password string
}

// Blog some comment
type Blog struct {
	gorm.Model
	BlogTitle   string
	BlogContent string
}

// DBHandler comment
type DBHandler struct {
	db *gorm.DB
}

// Init comment
func (db_handler *DBHandler) Init() {
	db, err := gorm.Open("sqlite3", "test.db")
	db_handler.db = db
	if err != nil {
		log.Fatalf("Error when connect database, the error is '%v'", err)
	}
}

// Migrate comment
func (db_handler *DBHandler) Migrate() {

	db_handler.db.AutoMigrate(&User{})
	db_handler.db.AutoMigrate(&Blog{})

	db_handler.db.Model(&User{}).AddUniqueIndex("idx_user_name", "Username")
}

// GetDBInstance comment
func (db_handler *DBHandler) GetDBInstance() *gorm.DB {
	return db_handler.db
}

// GetUser comment
func (db_handler *DBHandler) GetUser(username string, reply *User) error {
	db_handler.db.First(reply, "Username = ?", username)
	return nil
}

// GetAllUsers comment
func (db_handler *DBHandler) GetAllUsers(nullObject string, reply *[]User) error {
	db_handler.db.Find(reply)
	return nil
}

// GetAllBlogs comment
func (db_handler *DBHandler) GetAllBlogs(nullObject string, reply *[]Blog) error {
	db_handler.db.Find(reply)
	return nil
}

// GetBlog Comment
func (db_handler *DBHandler) GetBlog(identifier uint, reply *Blog) error {
	db_handler.db.First(reply, "ID = ?", identifier)
	return nil
}

// CreateUser comment
func (db_handler *DBHandler) CreateUser(user User, reply *User) error {
	db_handler.db.Create(&user)
	return db_handler.GetUser(user.Username, reply)
}

// CreateBlog comment
func (db_handler *DBHandler) CreateBlog(blog Blog, reply *Blog) error {
	db_handler.db.Create(&blog)
	return db_handler.GetBlog(blog.ID, reply)
}

// UpdateBlogContent comment
func (db_handler *DBHandler) UpdateBlogContent(blog Blog, reply *Blog) error {
	// fmt.Println("-->", blog)

	db_handler.db.Model(&blog).Update("BlogContent", blog.BlogContent)
	return db_handler.GetBlog(blog.ID, reply)
}

// UpdateBlogTitle comment
func (db_handler *DBHandler) UpdateBlogTitle(blog Blog, reply *Blog) error {
	db_handler.db.Model(&blog).Update("BlogTitle", blog.BlogTitle)
	return db_handler.GetBlog(blog.ID, reply)
}

// BatchInsertBlogs comment
func (db_handler *DBHandler) BatchInsertBlogs(blog []Blog, reply *[]Blog) error {
	fmt.Println("BatchInsertBlogs remote ", blog)
	for index := 0; index < len(blog); index++ {
		db_handler.db.Create(&blog[index])
	}
	return nil
}

// BatchInsertUsers comment
func (db_handler *DBHandler) BatchInsertUsers(user []User, reply *[]User) error {
	for index := 0; index < len(user); index++ {
		db_handler.db.Create(&user[index])
	}
	return nil
}
