package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func Compose() error {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://xblpqldtkjrpel:76e877a9b94725f547c1f3e5955c6c14461d1bd231ee88872e19473f74d1401f@ec2-44-194-92-192.compute-1.amazonaws.com:5432/d3q3rao0sk6vgl")

	if err != nil {
		return errors.New("не удалось создать инстанс миграций")
	}

	version, _, _ := m.Version()

	if version != 1 {
		if err := m.Up(); err != nil {
			return err
		}
	}
	return nil
}
