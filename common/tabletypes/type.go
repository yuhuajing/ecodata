package tabletypes

type Identity string

var (
	User  Identity = "user"
	Admin Identity = "admin"
)

type UserInfo struct {
	Id       string   `json:"id" gorm:"primary_key"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Identity Identity `json:"identity"`
	Approved int64    `json:"approved"`
}

type Info struct {
	Username string `json:"username"`
}

type Tousu struct {
	Id     string `json:"id" gorm:"primary_key"`
	Email  string `json:"email"`
	Tousu  string `json:"tousu"`
	Dealed int64  `json:"dealed"`
}

type TraceInfo struct {
	TraceID       string `json:"traceID"`
	ProdID        string `json:"prodID" `
	ProdBaseInfo  string `json:"prodBaseInfo"`
	ProdTxHash    string `json:"prodTxHash"`
	StoreID       string `json:"storeID" `
	StoreBaseInfo string `json:"storeBaseInfo"`
	StoreTxHash   string `json:"storeTxHash"`
	LogisID       string `json:"logisID" `
	LogisBaseInfo string `json:"logisBaseInfo"`
	LogisTxHash   string `json:"logisTxHash"`
	SaleID        string `json:"saleID" `
	SaleBaseInfo  string `json:"saleBaseInfo"`
	SaleTxHash    string `json:"saleTxHash"`
}

type ProduceInfo struct {
	Id           string `json:"id" gorm:"primary_key"`
	Prodcompany  string `json:"prodcompany"`
	Name         string `json:"name"`
	Prodbatchnum string `json:"prodbatchnum"`
	Prodtraceid  string `json:"prodtraceid"`
	Baseinfo     string `json:"baseinfo"`
	TxHash       string `json:"txHash"`
}
type EcoData struct {
	Waterdata string `json:"waterdata"`
	Codata    string `json:"codata"`
	Ecodata   string `json:"ecodata"`
	Operator  string `json:"operator"`
}

type EcoResData struct {
	Id        string `json:"id"`
	Waterdata string `json:"waterdata"`
	Codata    string `json:"codata"`
	Ecodata   string `json:"ecodata"`
	Operator  string `json:"operator"`
	Hash      string `json:"hash"`
}

type StoreInfo struct {
	Id            string `json:"id" gorm:"primary_key"`
	Storecompany  string `json:"storecompany"`
	Storebatchnum string `json:"storebatchnum"`
	Storetraceid  string `json:"storetraceid"`
	Baseinfo      string `json:"baseinfo"`
	TxHash        string `json:"txHash"`
}

type LogisInfo struct {
	Id            string `json:"id" gorm:"primary_key"`
	Logiscompany  string `json:"logiscompany"`
	Logisbatchnum string `json:"logisbatchnum"`
	Logistraceid  string `json:"logistraceid"`
	Baseinfo      string `json:"baseinfo"`
	TxHash        string `json:"txHash"`
}

type SaleInfo struct {
	Id           string `json:"id" gorm:"primary_key"`
	Salecompany  string `json:"salecompany"`
	Salebatchnum string `json:"salebatchnum"`
	Saletraceid  string `json:"saletraceid"`
	Baseinfo     string `json:"baseinfo"`
	TxHash       string `json:"txHash"`
}

type NewUserInfo struct {
	Username    string `json:"username"`
	Newusername string `json:"newusername"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
type NewAdminUserInfo struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
