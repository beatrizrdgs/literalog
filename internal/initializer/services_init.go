package initializer

func InitServices() error {
	var err error

	bookSvc, err = initBookSvc()
	if err != nil {
		return err
	}

	userSvc, err = initUserSvc()
	if err != nil {
		return err
	}

	logbookSvc, err = initLogbookSvc()
	if err != nil {
		return err
	}

	return nil
}
