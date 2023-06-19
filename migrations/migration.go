package migrations

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

/* func MigrateUp(conn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		log.Println("Error running migration: ", err)
		return err
	}
	return nil
}

func MigrateDownAll(conn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Down(); err != nil {
		log.Println("Error running migration: ", err)
		return err
	}
	return nil
}

func Steps(n int, conn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Steps(n); err != nil {
		log.Println("Error running migration: ", err)
		return err
	}
	return nil
}

func MigrateForce(conn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Force(-1); err != nil {
		log.Println("Error running migration: ", err)
		return err
	}
	return nil
}
*/
