package users

type Users struct {
	ID          uint   `gorm:"unique;not null;auto_increment" json:"id_user"`
	KodePetugas string `gorm:"primary_key" json:"kode_petugas"`
	NamaPetugas string `json:"nama_petugas"`
	Username    string `gorm:"unique;not null" json:"username"`
	Password    string `json:"password"`
	Level       string `json:"level"`
	Active      string `json:"active"`
	KodeDaerah  string `json:"kode_daerah"`
	Alamat      string `json:"alamat"`
	NoHP        string `json:"no_hp"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Allow       string `json:"allow"`
	NoKantor    string `json:"no_kantor"`
	IsMobile    string `json:"is_mobile"`
}

type View_users struct {
	ID          uint   `gorm:"unique;not null;auto_increment" json:"id_user"`
	KodePetugas string `gorm:"primary_key" json:"kode_petugas"`
	NamaPetugas string `json:"nama_petugas"`
	Username    string `gorm:"unique;not null" json:"username"`
	Password    string `json:"password"`
	Level       string `json:"level"`
	Active      string `json:"active"`
	KodeDaerah  string `json:"kode_daerah"`
	Alamat      string `json:"alamat"`
	NoHP        string `json:"no_hp"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Allow       string `json:"allow"`
	NoKantor    string `json:"no_kantor"`
	IsMobile    string `json:"is_mobile"`
	NamaDaerah  string `json:"nama_daerah"`
	KodeWilayah string `json:"kode_wilayah"`
	NamaWilayah string `json:"nama_wilayah"`
	Tahun       string `json:"tahun"`
}

// Menentukan nama tabel yang digunakan untuk model User
func (Users) TableName() string {
	return "sys_users"
}

func (View_users) TableName() string {
	return "view_user_v2"
}
