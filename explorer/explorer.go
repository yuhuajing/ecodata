package explorer

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	html "github.com/gofiber/template/html/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"log"
	"main/common/tabletypes"
	"main/core/database"
	"main/core/sendtx"
	txdata2 "main/core/txdata"
	"time"
)

func Explorer() {

	app := fiber.New(fiber.Config{
		Views: html.New("./explorer/views", ".html"),
	})
	app.Static("/", "./explorer/public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})
	//app.Get("/data", dealquery)
	app.Post("/login", deallogin)
	app.Post("/registerUser", registerUser)
	app.Post("/adminfunc", manageuser)
	app.Post("/userfunc", userfunc)
	app.Get("/approveuser", approveuser)
	app.Post("/giveapprove", giveapprove)
	app.Post("/removeapprove", removeapprove)
	app.Post("/updateAdminUser", updateAdminUser)
	app.Post("/addecodata", addecodata)
	app.Get("/getchaindata", getchaindata)
	app.Get("/getchaindatauser", getchaindatauser)
	app.Post("/delchaindata", delchaindata)
	app.Get("/gettxbyhash", gettxbyhash)
	app.Get("/checkdata", checkdata)
	app.Get("/checktx", checktx)

	//app.Use(jwtware.New(jwtware.Config{
	//	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	//}))
	app.Post("/adminlogin", dealadminlogin)

	app.Post("/updateUser", updateUser)

	app.Post("/manageuser", manageuser)
	app.Post("/managetousu", managetousu)
	app.Get("/adminusers", adminusers)
	app.Get("/alltousu", alltousu)

	app.Post("/managedrugs", managedrugs)
	app.Post("/updateProdData", updateProdData)

	app.Post("/tousu", addtousu)
	app.Post("/dealtousu", dealtousu)

	log.Fatal(app.Listen(":3005"))
}

type Tousu struct {
	Email string `json:"email"`
	Tousu string `json:"tousu"`
}

func getchaindatauser(c *fiber.Ctx) error {
	_, resdata := database.QueryChainData()
	return c.Render("chaindatauser", fiber.Map{
		"Data": resdata,
	})
}

func checkdata(c *fiber.Ctx) error {
	_, resdata := database.QueryChainData()
	return c.Render("chaindatacheck", fiber.Map{
		"Data": resdata,
	})
}
func getchaindata(c *fiber.Ctx) error {
	_, resdata := database.QueryChainData()
	return c.Render("chaindata", fiber.Map{
		"Data": resdata,
	})
}

func addtousu(c *fiber.Ctx) error {
	payload := &Tousu{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err := database.InsertTousu(payload.Email, payload.Tousu)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

//func dealquery(c *fiber.Ctx) error {
//	resdata := database.QueryStartBlock()
//	return c.Render("index", fiber.Map{
//		"Data": resdata,
//	})
//}

func dealtousu(c *fiber.Ctx) error {
	id := c.Query("id")
	database.DealTousu(id)
	return nil
}

func alltousu(c *fiber.Ctx) error {
	resdata := database.QueryAlltousu()
	return c.Render("alltousu", fiber.Map{
		"Data": resdata,
	})
}

func adminusers(c *fiber.Ctx) error {
	resdata := database.QueryAdminUserInfo()
	return c.Render("adminusers", fiber.Map{
		"Data": resdata,
	})
}

func approveuser(c *fiber.Ctx) error {
	_, resdata := database.QueryAllUserInfo()
	return c.Render("approveuser", fiber.Map{
		"Data": resdata,
	})
}

func managedrugs(c *fiber.Ctx) error {
	resdata := database.QueryAllTrace()
	return c.Render("manegedrugs", fiber.Map{
		"Data": resdata,
	})
}

func managetousu(c *fiber.Ctx) error {
	resdata := database.QueryUndealttousu()
	return c.Render("managetousu", fiber.Map{
		"Data": resdata,
	})
}

type ecodatas struct {
	Waterdata string `json:"waterdata"`
	Codata    string `json:"codata"`
	Ecodata   string `json:"ecodata"`
	Operator  string `json:"operator"`
}

func checktx(c *fiber.Ctx) error {
	id := c.Query("id")
	err, waterdata, codata, ecodata, operator := sendtx.GetProdData(id)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(ecodatas{
		Waterdata: waterdata,
		Codata:    codata,
		Ecodata:   ecodata,
		Operator:  operator,
	})
}

type txdata struct {
	Chainid   uint64    `json:"chainid"`
	Blockhash string    `json:"blockhash"`
	Blocknum  uint64    `json:"blocknum"`
	Txtime    time.Time `json:"txtime"`
	Contract  string    `json:"contract"`
	Gas       uint64    `json:"gas"`
	Confirmed uint64    `json:"confirmed"`
}

func gettxbyhash(c *fiber.Ctx) error {
	hash := c.Query("hash")
	err, chainid, blockhash, blocknum, txtime, contract, gas, blocks := txdata2.Gettxbyhash(hash)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(txdata{
		Chainid:   chainid,
		Blockhash: blockhash,
		Blocknum:  blocknum,
		Txtime:    txtime,
		Contract:  contract,
		Gas:       gas,
		Confirmed: blocks,
	})
}

func userfunc(c *fiber.Ctx) error {
	payload := &tabletypes.Info{}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, resdata := database.QueryUserInfoByName(payload.Username)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return c.Status(200).JSON(tabletypes.UserInfo{
		Username: resdata.Username,
		Id:       resdata.Id,
		Email:    resdata.Email,
		Phone:    resdata.Phone,
		Identity: resdata.Identity,
		Approved: resdata.Approved,
	})
}
func manageuser(c *fiber.Ctx) error {
	payload := &tabletypes.Info{}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, resdata := database.QueryUserInfoByName(payload.Username)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return c.Status(200).JSON(tabletypes.UserInfo{
		Username: resdata.Username,
		Id:       resdata.Id,
		Email:    resdata.Email,
		Phone:    resdata.Phone,
		Identity: resdata.Identity,
		Approved: resdata.Approved,
	})
}

type updatetraceData struct {
	Id       string `json:"id"`
	Baseinfo string `json:"baseinfo"`
}

func updateProdData(c *fiber.Ctx) error {
	payload := &updatetraceData{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//err := database.UpdateProdInfo(payload.Id, payload.Baseinfo)
	//if err != nil {
	//	return c.Status(400).JSON(DataResponse{
	//		Error:   err.Error(),
	//		Success: false,
	//		Data:    "",
	//	})
	//}
	return c.SendStatus(200)
}

func giveapprove(c *fiber.Ctx) error {
	id := c.Query("id")
	err := database.ApproveUserInfo(id)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

func delchaindata(c *fiber.Ctx) error {
	id := c.Query("id")
	err := database.DeleteChaindata(id)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}
func removeapprove(c *fiber.Ctx) error {
	id := c.Query("id")
	err := database.DeleteUserInfo(id)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

type UserRes struct {
	Token string `json:"token"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
}

func deallogin(c *fiber.Ctx) error {
	payload := &tabletypes.UserInfo{}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//fmt.Println(payload)
	err, userinfo := database.QueryUserInfos(payload.Username, payload.Password)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	myClaims := &jwt.MapClaims{
		"name": "John Doe",
		"uid":  utils.UUIDv4(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	c.Locals("user", myClaims)
	fmt.Println(t)
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Token:   t,
		Data:    payload.Username,
		Type:    userinfo.Identity,
	})
}

func dealadminlogin(c *fiber.Ctx) error {
	payload := &tabletypes.UserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, userinfo := database.QueryUserInfo(payload.Username, payload.Password, tabletypes.Admin)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
			Type:    userinfo.Identity,
		})
	}
	return c.SendStatus(200)
}

type DataResponse struct {
	Error   string              `json:"error"`
	Success bool                `json:"success" default:"false"`
	Data    string              `json:"data"`
	Token   string              `json:"token"`
	Type    tabletypes.Identity `json:"type"`
}

func registerUser(c *fiber.Ctx) error {
	//fmt.Println(string(c.Body()))
	payload := &tabletypes.UserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err := database.InsertUserInfo(payload.Username, payload.Password, payload.Email, payload.Phone, tabletypes.User)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "Duplicated",
		})
	}
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Token:   "",
		Data:    "Register successfully!!",
		Type:    "",
	})
}
func updateAdminUser(c *fiber.Ctx) error {
	payload := &tabletypes.NewAdminUserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	err := database.UpdateAdminUserInfo(payload.Username, payload.Phone, payload.Email)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}
func updateUser(c *fiber.Ctx) error {
	payload := &tabletypes.NewUserInfo{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	err := database.UpdateUserInfo(payload.Username, payload.Newusername, payload.Password, payload.Phone, payload.Email)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.SendStatus(200)
}

func addecodata(c *fiber.Ctx) error {
	payload := &tabletypes.EcoResData{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	err, hash := database.InsertProdInfo(payload.Id, payload.Ecodata, payload.Codata, payload.Operator, payload.Waterdata)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	//return c.Render("success", fiber.Map{})
	return c.Status(200).JSON(DataResponse{
		Error:   "",
		Success: true,
		Data:    hash,
	})
}
