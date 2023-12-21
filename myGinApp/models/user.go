package models

// User represents the user model
type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// GetAllUsers returns a list of users (mocked data for now)
func GetAllUsers() ([]User, error) {
    // In a real app, you would query your database here
    users := []User{
        {ID: 1, Name: "Alice", Email: "alice@example.com"},
        {ID: 2, Name: "Bob", Email: "bob@example.com"},
    }
    return users, nil
}
