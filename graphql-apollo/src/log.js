const { ApolloServer, gql, PubSub } = require("apollo-server");

const typeDefs = gql`
  type Query { # todo pagination
    log: LogSession!
  }

  type LogSession {
    id: ID!
    createdAt: String!
    args: String! # message
    resolver: String! # field name
    ipAddress: String
    origin: String!
  }

  type Mutation {
    message(complaint: String!): String!
    createLogSession(complaint: String!): LogSession!
  }
  type Subscription {
    newMessage: String!
  }
`;

const NEW_MESSAGE = "NEW_MESSAGE";

const resolvers = {
  Subscription: {
    newMessage: {
      subscribe: (_, __, { pubsub }) => pubsub.asyncIterator(NEW_MESSAGE)
    }
  },
  Query: {
    log: () => {
      return {
        id: 1,
        createdAt: "test",
        args: "test",
        resolver: "test",
        origin: "test",
        ipAddress: "test"
      }
    }
  },
  Mutation: {
    message: async (_, { complaint }, { context, pubsub }) => {
        await MiddlewareSession(complaint, context)
        pubsub.publish(NEW_MESSAGE, {
            newMessage: complaint
        });

        return complaint
    },
    createLogSession: (_, { complaint }, context, info) => {
        console.log(context.req.header.IPAddress)
        const data = {
            id: 1,
            createdAt: new Date(),
            args: complaint,
            resolver: info.fieldName,
            origin: context.req.header('origin'),
            ipAddress: context.req.header('x-forwarded-for')
        }
        return data
    }

  }
};

const MiddlewareSession = (complaint, context) => {
    //context.db.Mutation.createLogSession({})
}

const pubsub = new PubSub();

const server = new ApolloServer({
  typeDefs,
  resolvers,
  context: ({ req, res }) => ({ req, res, pubsub })
});

//server.applyMiddleware({ app });

server.listen().then(({ url }) => console.log(`server started at ${url}`));