package queries

const (
	Users = `
	query Users($username: String!) {
	users(where: {username: {_eq: $username}}) {
		id
		password
  }
}
`
)
