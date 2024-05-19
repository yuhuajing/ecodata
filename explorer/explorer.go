package explorer

import (
	"encoding/csv"
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	html "github.com/gofiber/template/html/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"main/common/tabletypes"
	"main/core/database"
	"main/core/sendtx"
	txdata2 "main/core/txdata"
	"os"
	"path/filepath"
	"strings"
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
	app.Get("/searchchaindata", searchchaindata)

	app.Post("/upload", upload)

	log.Fatal(app.Listen(":3005"))
}

func getchaindatauser(c *fiber.Ctx) error {
	//fmt.Println(123)
	name := c.Query("username")
	//fmt.Println(name)
	_, resdata := database.QueryUserChainData(name)
	return c.Render("chaindatauser", fiber.Map{
		"Data": resdata,
	})
}

func checkdata(c *fiber.Ctx) error {
	name := c.Query("username")
	_, resdata := database.QueryUserChainData(name)
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

func approveuser(c *fiber.Ctx) error {
	_, resdata := database.QueryAllUserInfo()
	return c.Render("approveuser", fiber.Map{
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

func upload(c *fiber.Ctx) error {
	// Retrieve the file from the form data
	handler, err := c.FormFile("file")
	file, err := handler.Open()
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	// Create the uploads directory if it doesn't exist
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(filepath.Join("uploads", handler.Filename))
	//fmt.Println(handler.Filename)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the server
	if _, err := io.Copy(dst, file); err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}

	if strings.HasSuffix(handler.Filename, ".csv") {
		fmt.Println("3453456")
		path := filepath.Join("uploads", handler.Filename)
		err := readCSV(path)
		if err != nil {
			return c.Status(400).JSON(DataResponse{
				Error:   err.Error(),
				Success: false,
				Data:    "",
			})
		}
	}
	return c.SendStatus(200)
}

func readCSV(filepath string) error {
	// Open CSV file
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return err
	}
	defer f.Close()
	fmt.Println("READCSV")
	utf8Reader := transform.NewReader(f, simplifiedchinese.GBK.NewDecoder())
	r := csv.NewReader(utf8Reader)
	// Read File into *lines* variable
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("无法读取CSV文件:", err)
		return err
	}

	for index, line := range lines {
		if index == 0 {
			continue
		}
		id := utils.UUID() //  strings.TrimSpace(line[0])
		waterdata := strings.TrimSpace(line[0])
		if waterdata == "" {
			break
		}
		ecode := strings.TrimSpace(line[1])
		codata := strings.TrimSpace(line[2])
		operator := strings.TrimSpace(line[3])
		err, _ := database.InsertProdInfo(id, ecode, codata, operator, waterdata)
		if err != nil {
			return err
		}
		//addressesStr = append(addressesStr, address)
		//quantities = append(quantities, quantity)
	}
	return nil
}

func searchchaindata(c *fiber.Ctx) error {
	fmt.Println(123)
	search := c.Query("search")
	fmt.Println(search)
	err, resdata := database.QueryChainDataByFileter(search)
	if err != nil {
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	return c.Render("searchData", fiber.Map{"Data": resdata})
	//return c.Status(200).JSON(DataResponse{
	//	Error:   "",
	//	Success: true,
	//	Data:    hash,
	//})
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
