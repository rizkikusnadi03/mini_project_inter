package domain

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Phone     string     `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	Password  string     `gorm:"not null" json:"-"`
	Role      string     `gorm:"type:varchar(20);default:'user'" json:"role"` // 'admin' or 'user'
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Store     *Store     `json:"store,omitempty"`
	Addresses []Address  `json:"addresses,omitempty"`
}

type Store struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"uniqueIndex;not null" json:"user_id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"`
	Description    string    `gorm:"type:text" json:"description"`
	ProfilePicture string    `gorm:"type:varchar(255)" json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Products       []Product `json:"products,omitempty"`
}

type Address struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null;index" json:"user_id"`
	Title          string    `gorm:"type:varchar(100);not null" json:"title"`
	AddressDetails string    `gorm:"type:text;not null" json:"address_details"`
	ProvID         string    `gorm:"type:varchar(50)" json:"prov_id"`
	CityID         string    `gorm:"type:varchar(50)" json:"city_id"`
	IsPrimary      bool      `gorm:"default:false" json:"is_primary"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StoreID     uint      `gorm:"not null;index" json:"store_id"`
	CategoryID  uint      `gorm:"not null;index" json:"category_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null" json:"stock"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	UserID        uint         `gorm:"not null;index" json:"user_id"`
	TotalPrice    float64      `gorm:"not null" json:"total_price"`
	Quantity      int          `gorm:"not null" json:"quantity"`
	PaymentMethod string       `gorm:"type:varchar(50);not null" json:"payment_method"`
	Status        string       `gorm:"type:varchar(50);default:'completed'" json:"status"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	ProductLogs   []ProductLog `json:"product_logs,omitempty"`
}

type ProductLog struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TransactionID uint      `gorm:"not null;index" json:"transaction_id"`
	ProductID     uint      `gorm:"not null" json:"product_id"`
	ProductName   string    `gorm:"type:varchar(255);not null" json:"product_name"`
	ProductPrice  float64   `gorm:"not null" json:"product_price"`
	CreatedAt     time.Time `json:"created_at"`
}
