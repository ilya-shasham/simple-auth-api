package handlers

import (
	"auth-api/checks"
	"auth-api/dbutils"
	"auth-api/globals"
	"auth-api/models"
	"errors"
)

var usernameChecks = []checks.CheckFunc[byte]{
	checks.ContainsOnlyFrom([]byte(globals.Alphanumnerics + "_")),
	checks.LengthAtMost[byte](255),
	func(b []byte) error {
		_, exists := dbutils.RowExists[models.Member]("username = ?", func(m *models.Member) bool { return m.Username != "" }, string(b))

		if exists {
			return errors.New("username taken")
		}

		return nil
	},
}

var passwordChecks = []checks.CheckFunc[byte]{
	checks.ContainsFromAtLeast([]byte(globals.AlphabetLower), 3),
	checks.ContainsFromAtLeast([]byte(globals.AlphabetUpper), 3),
	checks.ContainsFromAtLeast([]byte(globals.Numerics), 3),
	checks.ContainsFromAtLeast([]byte(globals.Symbols+" "), 1),
	checks.LengthAtMost[byte](70),
}
