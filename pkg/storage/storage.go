package storage

// Publication retrieved from RSS.
type Post struct {
    ID      int    // record number
    Title   string // publication title
    Content string // publication content
    PubTime int64  // publication time
    Link    string // publication link
}

// Interface specifies the contract for working with the database.
type Interface interface {
	Posts(int) ([]Post, error) // Get the last n publications from the database.
	AddPosts([]Post) error     // Add publications to the database.
}
