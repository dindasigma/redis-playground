const { ApolloServer, gql, PubSub } = require("apollo-server");

const typeDefs = gql`
  type Query {
    hello(name: String): String
  }

  type Mutation {
    message(complaint: String!): String!
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
    hello: (_, { name }) => {
      return `hey ${name}`;
    }
  },
  Mutation: {
    message: (_, { complaint }, { pubsub }) => {
        pubsub.publish(NEW_MESSAGE, {
            newMessage: complaint
        });

        return complaint
    }
  }
};

const pubsub = new PubSub();

const server = new ApolloServer({
  typeDefs,
  resolvers,
  context: ({ req, res }) => ({ req, res, pubsub })
});

server.listen().then(({ url }) => console.log(`server started at ${url}`));