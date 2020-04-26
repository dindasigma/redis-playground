const { ApolloServer, gql, PubSub } = require("apollo-server");

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

    type Subscription {
        newUser: User!
    }
`;

const NEW_USER = "NEW_USER";

const resolvers = {
	Subscription: {
		newUser: {
			subscribe: (_, __, { pubsub }) => pubsub.asyncIterator(NEW_USER)
		}
	},
	User: {
		firstLetterOfUsername: parent => {
			return parent.username[0]
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
		register: (_, { userInfo: { username } }, { pubsub }) => {
			const user = {
				id: 1,
				username
			}

			pubsub.publish(NEW_USER, {
				newUser: user
			})

			return {
				errors: null,
				user

			};
		}
	}
}

const pubsub = new PubSub();

const server = new ApolloServer({
	typeDefs,
	resolvers,
	context: ({ req, res }) => ({ req, res, pubsub })
});

server.listen().then(({ url }) => console.log(`server started at ${url}`));
