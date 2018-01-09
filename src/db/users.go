package db

import (
	"fmt"
	"os"
)

func GetArma2Users() (*[]User, error) {
	return internalGetUsers("user_arma2_id", "user_arma2_nick")
}

func GetArma3Users() (*[]User, error) {
	return internalGetUsers("user_arma3_id", "user_arma3_nick")
}

func internalGetUsers(idField string, nickField string) (*[]User, error) {
	var query = `
		SELECT
			users.user_id,
			users.user_name,
			users.` + idField + `,
			users.` + nickField + `
		FROM
			fusion7_users AS users
		WHERE
			users.user_status = 0
			AND
			CHAR_LENGTH(users.` + idField + `) > 0
			AND
			CHAR_LENGTH(users.` + nickField + `) > 0
		ORDER BY
			users.user_name
	`
	rows, err := Database.Query(query)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying database: %q\n", err)
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Arma.ID, &user.Arma.Nick)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading user from database: %q\n", err)
		} else {
			users = append(users, user)
		}
	}

	return &users, nil
}
