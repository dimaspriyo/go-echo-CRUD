package repository

type IPostgresqlRepository interface {
	get()
	getAll()
	create()
	update()
	delete()
}
