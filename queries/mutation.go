package queries

const (
	SignupMutation = `
        mutation SignUp ($username: String!, $password: String!) {
            insert_users(objects: {username: $username, password: $password}) {
                returning {
                    id
                }
            }
        }
    `
	LoginMutation = `
        mutation Login($username: String!, $password: String!) {
            login_users(where: {username: {_eq: $username}, password: {_eq: $password}}) {
                id
                username
            }
        }
    `
)
