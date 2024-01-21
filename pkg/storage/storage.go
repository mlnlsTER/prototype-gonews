package storage

// Публикация, получаемая из RSS.
type Post struct {
    ID      int    // номер записи
    Title   string // заголовок публикации
    Content string // содержание публикации
    PubTime int64  // время публикации
    Link    string // ссылка на источник
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts(int) ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	DeletePost(Post) error  // удаление публикации по ID
}
