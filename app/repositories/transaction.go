package repositories

func (uow *UnitOfWork) Start() *UnitOfWork {
	tx := uow.Tx.Begin()
	return &UnitOfWork{
		PSQLRepository: NewPSQLRepository(tx),
		Tx:             tx,
	}
}

func (uow *UnitOfWork) Complete() error {
	return uow.Tx.Commit().Error
}

func (uow *UnitOfWork) Dispose() error {
	return uow.Tx.Rollback().Error
}
