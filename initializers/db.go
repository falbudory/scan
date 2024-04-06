package initializers

import (
	"github.com/zetamatta/go-outputdebug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"serverWeb/models"
	"serverWeb/utils"
	"time"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("USER_DB") + ":" + os.Getenv("PASSWORD") + `@tcp(127.0.0.1:3306)/` + os.Getenv("DATABASE") + `?charset=utf8mb4&parseTime=True&loc=Local`

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		outputdebug.String("[BCRW]: " + err.Error())

	}
}

func MigrateDB() {
	if err := DB.AutoMigrate(&models.Role{}); err != nil {
		outputdebug.String("[BCRW]: " + err.Error())
	}
	if err := DB.AutoMigrate(&models.Permission{}); err != nil {
		outputdebug.String("[BCRW]: " + err.Error())
	}
	if err := DB.AutoMigrate(&models.RolePermission{}); err != nil {
		outputdebug.String("[BCRW]: " + err.Error())
	}
	if err := DB.AutoMigrate(&models.TypeUser{}); err != nil {
		outputdebug.String("[BCRW]: " + err.Error())
	}
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		outputdebug.String("[BCRW]: " + err.Error())
	}
}
func GenData() {
	roles := []models.Role{
		{Name: "Administrator"},
		{Name: "Sale Agent"},
		{Name: "Business"},
		{Name: "Instructor"},
		{Name: "Student"},
	}

	typeUsers := []models.TypeUser{
		{Name: "Administrator"},
		{Name: "Sale Agent"},
		{Name: "Business"},
		{Name: "Instructor"},
		{Name: "Student"},
	}

	permissions := []models.Permission{
		{Name: "Home", Permission: "home"},
		{Name: "Dashboard", Permission: "dashboard"},
		{Name: "Courses", Permission: "courses"},
		{Name: "Classes", Permission: "classes"},
		{Name: "Management Instructors", Permission: "mng_instructors"},
		{Name: "Management Students", Permission: "mng_students"},
		{Name: "Management Sales Agents and Business", Permission: "mng_sale_business"},
		{Name: "Management Classes", Permission: "mng_classes"},
		{Name: "Management Courses", Permission: "mng_courses"},
		{Name: "Management Users", Permission: "account_users"},
		{Name: "Management Roles", Permission: "account_roles"},
		{Name: "Management Information", Permission: "management_info"},
		{Name: "Management Account", Permission: "management_account"},
	}

	rolePermissions := []models.RolePermission{
		{RoleID: 1, PermissionID: 1},
		{RoleID: 1, PermissionID: 2},
		{RoleID: 1, PermissionID: 5},
		{RoleID: 1, PermissionID: 6},
		{RoleID: 1, PermissionID: 7},
		{RoleID: 1, PermissionID: 8},
		{RoleID: 1, PermissionID: 9},
		{RoleID: 1, PermissionID: 10},
		{RoleID: 1, PermissionID: 11},
		{RoleID: 1, PermissionID: 12},
		{RoleID: 1, PermissionID: 13},

		{RoleID: 2, PermissionID: 1},
		{RoleID: 2, PermissionID: 2},
		{RoleID: 2, PermissionID: 5},
		{RoleID: 2, PermissionID: 6},
		{RoleID: 2, PermissionID: 12},

		{RoleID: 3, PermissionID: 1},
		{RoleID: 3, PermissionID: 2},
		{RoleID: 3, PermissionID: 5},
		{RoleID: 3, PermissionID: 6},
		{RoleID: 3, PermissionID: 12},

		{RoleID: 4, PermissionID: 1},
		{RoleID: 4, PermissionID: 3},
		{RoleID: 4, PermissionID: 4},

		{RoleID: 5, PermissionID: 1},
		{RoleID: 5, PermissionID: 3},
		{RoleID: 5, PermissionID: 4},
	}

	user := models.User{
		TypeUserID:  1,
		CodeUser:    "ADM0001",
		FirstName:   "Administrators",
		RoleID:      1,
		Email:       "Admin@gmail.com",
		PhoneNumber: "0123456789",
		Address:     "California",
		Username:    "Administrators",
		Password:    utils.HashingPassword("Admin@#$2024"),
		State:       true,
		Verify:      true,
		Token:       "",
		Deleted:     false,
		UpdatedBy:   1,
		CreatedBy:   1,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}

	for _, data := range typeUsers {
		if err := DB.Create(&data).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
		}
	}

	for _, data := range roles {
		if err := DB.Create(&data).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
		}
	}

	for _, data := range permissions {
		if err := DB.Create(&data).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
		}
	}

	for _, data := range rolePermissions {
		if err := DB.Create(&data).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
		}
	}

	if err := DB.Create(&user).Error; err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
	}

}
