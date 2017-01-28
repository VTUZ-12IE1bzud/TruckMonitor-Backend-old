package model

/** Модель данных. Пользователь. */
type User struct {
	Id			int
	RoleId		int
	Email		string
	Password	string
	Surname		string
	Name		string
	Patronymic	string
	DateOfBirth string
	Photo		string
	Phone		string
}
