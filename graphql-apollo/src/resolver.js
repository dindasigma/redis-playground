const { ApolloServer, gql } = require("apollo-server");

// args in resolver
// 1. argument
// 2. context: define on ApolloServer

// also can create a resolver from field, and will overriding that field

const typeDefs = gql`
	type Query {
		hello(name: String): String!
		user: User
	}

	type User {
		id: ID!
		username: String
		firstLetterOfUsername: String
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
		register(userInfo: UserInfo): RegisterResponse!
		login(userInfo: UserInfo): String!
	}
`;

const resolvers = {
	User: {
		firstLetterOfUsername: parent => {
			return parent.username ? parent.username[0] : null;
		}
	},
	Query: {
		hello: (parent, { name }) => {
			return `hello world ${name}`
		},
		user: () => ({
			id: 1,
			username: "bob"
		})
	},
	Mutation: {
		login: (parent, { userInfo: { username } }, context, info) => {
			console.log(context)
			return username;
		},
		register: () => ({
			errors: null,
			user: {
				id: 1,
				username: "bob"
			}

		})
	}
}

const server = new ApolloServer({
	typeDefs,
	resolvers,
	context: ({ req, res }) => ({ req, res })
});

server.listen().then(({ url }) => console.log(`server started at ${url}`));
