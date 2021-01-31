package jsonplaceholder

import "fmt"

// User structure
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

// GetUsers returns list of User
func GetUsers() (users []User, err error) {
	err = get("users", &users)
	if err != nil {
		return nil, err
	}
	return users, err
}

// GetUser returns User by ID
func GetUser(id int) (user User, err error) {
	err = get(fmt.Sprintf("users/%d", id), &user)
	if err != nil {
		return user, err
	}
	return user, err
}
