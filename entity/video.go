package entity

type Person struct {
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    Age int8 `json:"age"`
    Email string `json:"email" validate:"required, email"`
}

type Video struct {
    Title string `json:"title" binding: "min=3 max=20" validate:"is-cool"`
    Description string `json:"description" binding:"max=200"`
    URL string `json:"url" binding: "required, url"`
    Author Person `json:"author" binding:"required"`
}

