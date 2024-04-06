package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/zetamatta/go-outputdebug"
	"serverWeb/initializers"
	"serverWeb/models"
	"serverWeb/structs"
	"serverWeb/utils"
	"time"
)

var SessAuth = session.New(session.Config{
	CookieSessionOnly: true,
})

func GetLogin(c *fiber.Ctx) error {
	var user models.User
	initializers.DB.First(&user)
	return c.Render("pages/login/index", fiber.Map{
		"User": user,
	}, "layouts/main")
}

func PostLogin(c *fiber.Ctx) error {
	var user models.User

	DB := initializers.DB
	form := new(models.User)

	if err := c.BodyParser(form); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())

		return c.Render("pages/login/index", fiber.Map{
			//"Notify":  true,
			//"Message": "Enter complete information",
		})
	}

	if err := DB.Where(
		"BINARY username = ?", form.Username).Where(
		"deleted", false).Where(
		"state", true).First(&user).Error; err != nil {

		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())

		return c.Render("pages/login/index", fiber.Map{
			//"Notify":  true,
			//"Message": "Username or Password is incorrect",
		})
	}

	if utils.CheckPasswordHash(form.Password, user.Password) {

		user.Session = "session_" + form.Username
		if err := DB.Save(&user).Error; err != nil {

			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
			return c.Render("pages/login/index", fiber.Map{
				//"Notify":  true,
				//"Message": "Username or Password is incorrect",
			})
		}

		sess, _ := SessAuth.Get(c)
		sess.Set("username", user.Username)
		sess.Set("login_success", "authenticated")
		sess.Set("sessionId", "session_"+form.Username)
		if err := sess.Save(); err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
		}

		return c.Redirect("/home")
	}

	return c.Render("pages/login/index", fiber.Map{
		//"Notify":  true,
		//"Message": "Username or Password is incorrect",
	})
}

func PostSignupStudent(c *fiber.Ctx) error {
	var user structs.User
	var account models.User
	DB := initializers.DB
	if err := c.BodyParser(&user); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Format User Fail")
	}

	if err := DB.Where("username", user.Username).Or("email", user.Email).First(&models.User{}).Error; err != nil {
		account.TypeUserID = 5
		account.FirstName = user.FirstName
		account.LastName = user.LastName
		account.RoleID = 5
		account.Email = user.Email
		account.PhoneNumber = user.PhoneNumber
		account.Address = user.Address
		account.Username = user.Username
		account.Password = utils.HashingPassword(user.Password)
		account.ReferralCode = user.ReferralCode
		account.Session = ""
		account.NameBusiness = ""
		account.FullNameRepresentative = ""
		account.State = true
		account.Verify = false
		account.Deleted = false
		account.CreatedAt = time.Now()
		account.DeletedAt = time.Now()
		account.UpdatedAt = time.Now()
		account.DeletedBy = 0
		account.UpdatedBy = 0

		if err := DB.Create(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		account.CodeUser = "STUD" + fmt.Sprintf("%04d", account.UserID)
		account.CreatedBy = account.UserID

		if err := DB.Updates(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		return c.JSON("Success")
	}

	return c.JSON("Username or Email already exists")
}

func PostSignupInstructor(c *fiber.Ctx) error {
	var user structs.User
	var account models.User
	DB := initializers.DB
	if err := c.BodyParser(&user); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Format User Fail")
	}

	if err := DB.Where("username", user.Username).Or("email", user.Email).First(&models.User{}).Error; err != nil {
		account.TypeUserID = 4
		account.FirstName = user.FirstName
		account.LastName = user.LastName
		account.RoleID = 4
		account.Email = user.Email
		account.PhoneNumber = user.PhoneNumber
		account.Address = user.Address
		account.Username = user.Username
		account.Password = utils.HashingPassword(user.Password)
		account.ReferralCode = user.ReferralCode
		account.Session = ""
		account.NameBusiness = ""
		account.FullNameRepresentative = ""
		account.State = false
		account.Verify = false
		account.Deleted = false
		account.CreatedAt = time.Now()
		account.DeletedAt = time.Now()
		account.UpdatedAt = time.Now()
		account.DeletedBy = 0
		account.UpdatedBy = 0

		if err := DB.Create(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		account.CodeUser = "INSTR" + fmt.Sprintf("%04d", account.UserID)
		account.CreatedBy = account.UserID

		if err := DB.Updates(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		return c.JSON("Success")
	}

	return c.JSON("Username or Email already exists")
}

func PostSignupSale(c *fiber.Ctx) error {
	var user structs.User
	var account models.User
	DB := initializers.DB
	if err := c.BodyParser(&user); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Format User Fail")
	}

	if err := DB.Where("username", user.Username).Or("email", user.Email).First(&models.User{}).Error; err != nil {
		account.TypeUserID = 2
		account.FirstName = user.FirstName
		account.LastName = user.LastName
		account.RoleID = 2
		account.Email = user.Email
		account.PhoneNumber = user.PhoneNumber
		account.Address = user.Address
		account.Username = user.Username
		account.Password = utils.HashingPassword(user.Password)
		account.ReferralCode = ""
		account.Session = ""
		account.NameBusiness = ""
		account.FullNameRepresentative = ""
		account.State = false
		account.Verify = false
		account.Deleted = false
		account.CreatedAt = time.Now()
		account.DeletedAt = time.Now()
		account.UpdatedAt = time.Now()
		account.DeletedBy = 0
		account.UpdatedBy = 0

		if err := DB.Create(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		account.CodeUser = "SAL" + fmt.Sprintf("%04d", account.UserID)
		account.CreatedBy = account.UserID

		if err := DB.Updates(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		return c.JSON("Success")
	}

	return c.JSON("Username or Email already exists")
}

func PostSignupBusiness(c *fiber.Ctx) error {
	var user structs.User
	var account models.User
	DB := initializers.DB
	if err := c.BodyParser(&user); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Format User Fail")
	}

	if err := DB.Where("username", user.Username).Or("email", user.Email).First(&models.User{}).Error; err != nil {
		account.TypeUserID = 3
		account.FirstName = ""
		account.LastName = ""
		account.RoleID = 3
		account.Email = user.Email
		account.PhoneNumber = user.PhoneNumber
		account.Address = user.Address
		account.Username = user.Username
		account.Password = utils.HashingPassword(user.Password)
		account.ReferralCode = ""
		account.Session = ""
		account.NameBusiness = user.NameBusiness
		account.FullNameRepresentative = user.FullNameRepresentative
		account.State = false
		account.Verify = false
		account.Deleted = false
		account.CreatedAt = time.Now()
		account.DeletedAt = time.Now()
		account.UpdatedAt = time.Now()
		account.DeletedBy = 0
		account.UpdatedBy = 0

		if err := DB.Create(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		account.CodeUser = "BUSIN" + fmt.Sprintf("%04d", account.UserID)
		account.CreatedBy = account.UserID

		if err := DB.Updates(&account).Error; err != nil {
			outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + "Can not create account")
			return c.JSON("Can not create account")
		}

		return c.JSON("Success")
	}

	return c.JSON("Username or Email already exists")
}

func GetLogout(c *fiber.Ctx) error {
	sess, _ := SessAuth.Get(c)
	if err := sess.Reset(); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
	}

	if err := sess.Save(); err != nil {
		outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [BCRW]: " + err.Error())
	}

	return c.Redirect("/login")
}
