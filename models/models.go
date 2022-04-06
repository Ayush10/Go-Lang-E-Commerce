package models

type User struct {
	ID
	FirstName
	LastName
	Password
	Email
	Phone
	Token
	RefreshToken
	CreatedAt
	UpdatedAt
	UserID
	UserCart
	AddressDetails
	OrderStatus
}
