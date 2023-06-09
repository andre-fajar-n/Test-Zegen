package models

func (BookAuthor) TableName() string {
	return "book_author"
}

func (BookAuthorForeignKeyAuthor) TableName() string {
	return "authors"
}

func (BookAuthorForeignKeyBook) TableName() string {
	return "books"
}
