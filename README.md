# Crest

Simple CRUD rest. This project uses generic to convert functions to http handlers.
This is POC and only create is currently `investigated`.

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
func CreateUser(r *http.Request, s *CreateUser) (*User, error) {
    // either return error
    // return nil, crest.Error(http.status.BadRequest, nil)
    // or return user
    return &User{
        ID: uuid.New().String,
        Username: s.Username,
        Email: s.Email,
    }, nil
}

func init() {
    // now let's register this handler to e.g. gorilla mux router
    r := mux.NewRouter()
    r.Path("/api/user/", create.Handler(CreateUser)).Method(http.MethodPut)
}
```

We have also ability to provide handler that does not return struct type, but generic `Responder` interface.
So let's try an example and let's suppose we have defined structures from previous example

```go
func CreateUser(*http.Request, s *CreateUser) (Responder, error) {
    // create user and return back as custom response
    return Response(http.StatusCreated, User{ID:uuid.New().String}), nil
}

```

# Author
Peter Vrba <phonkee@phonkee.eu>
