# Crest

Simple CRUD rest. This project uses generic to convert functions to http handlers

# Example

Let's try to create create handler from functions

```go

type CreateUser struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
}

type User struct {
    ID string `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
}

// CreateUser knows how to create user
func CreateUser(r *http.Request, s CreateUser) (*User, error) {
    // some logic how to create user
}

func init() {
    // now let's register this handler to e.g. gorilla mux router
    r := mux.NewRouter()
    r.Path("/api/user/", create.Handler(CreateUser)).Method(http.MethodPut)
}
```

# Author
Peter Vrba <phonkee@phonkee.eu>
