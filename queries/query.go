package queries

const (
	User = `
	query User($username: String!) {
		users(where: {username: {_eq: $username}}) {
			id
			password
		}
	}

   `
)
