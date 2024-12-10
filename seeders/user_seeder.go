package seeders

import (
    "database/sql"
    "fmt"
    "math/rand"
    "github.com/bxcodec/faker/v3"
)

func SeedUsers(db *sql.DB) error {
    query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`

    for i := 0; i < 10; i++ {
        name := faker.Name()
        email := faker.Email()
        age := rand.Intn(50) + 20 // Random age between 20 and 70

        // Insert user into the users table
        _, err := db.Exec(query, name, email, age)
        if err != nil {
            return fmt.Errorf("Error seeding user %d: %v", i+1, err)
        }

        fmt.Printf("User %d inserted: Name=%s, Email=%s, Age=%d\n", i+1, name, email, age)
    }

    fmt.Println("Seeding complete!")
    return nil
}
