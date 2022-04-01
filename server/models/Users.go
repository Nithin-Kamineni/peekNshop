package models

type SignInReply struct {
	Msg string
}

type LogInReply struct {
	Msg         string
	UserDetails User3
	AllowUsers  bool
}

type Address struct {
	Address string
}

type Coardinates struct {
	Lat string
	Lon string
}

type HomePageCity struct {
	City string `json:"city"`
}

type User3 struct {
	ID             string `gorm:"primary_key" json:"id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `gorm:"uniqueIndex:idx_first_second" json:"email"`
	Password       string `json:"password"`
	Acesskey       string
	RefreshKey     string
	FavorateStores []string
	Address1       string
	Address2       string
	Address3       string
}

type FavorateStoresObj struct {
	ID            string
	Acesskey      string
	FavorateStore string
}

type Orders struct {
	OrderID     string
	ProductID   string
	storeID     string
	UserID      string
	orderedOn   string
	deliveredOn string
	pickedUpOn  string
}

type ChangeUserAddress struct {
	ID         string `gorm:"primary_key" json:"id"`
	Acesskey   string
	RefreshKey string
	Address    string
}

type ContactMsgObj struct {
	Name  string `gorm:"primary_key" json:"id"`
	Email string
	Msg   string
}

type RetrevalDetails struct {
	Email string `json:"email"`
}
