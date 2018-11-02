package entity

type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

func GetName(u User) string {
	return u.Name
}
func GetPhone(u User) string {
	return u.Phone
}
func GetEmail(u User) string {
	return u.Email
}
func GetPassword(u User) string {
	return u.Password
}
