const { ApolloServer, gql } = require("apollo-server");

// args

const typeDefs = gql`
	type Query {
		hello: String!
		user: User
	}

	type User {
		id: ID!
		username: String!
	}

	type Error {
			field: String!
			message: String!
	}

	type RegisterResponse {
		errors: [Error]
		user: User   
	}

	input UserInfo {
		username: String!
		password: String!
		age: Int
	}

	type Mutation {
		register(UserInfo: UserInfo): RegisterResponse!
		login(UserInfo: UserInfo): Boolean!
	}
`;

const resolvers = {
	Query: {
		hello: () => "hello world",
		user: () => ({
			id: 1,
			username: "bob"
		})
	},
	Mutation: {
		login: () => true,
		register: () => ({
			errors: null,
			user: {
				id: 1,
				username: "bob"
			}

		})
	}
}

const server = new ApolloServer({ typeDefs, resolvers });

server.listen().then(({ url }) => console.log(`server started at ${url}`));

/*
{
  hello
  user {
    id,
    username
  }
}

mutation {
  register(UserInfo :{username:"dinda", password:"sigma"}) {
    errors {
      field,
      message
    }
    user {
  		id,
    	username
    }
  }
  
  login(UserInfo :{username:"dinda", password:"sigma"})
}
*/
