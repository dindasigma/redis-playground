import Container from 'components/container';
import Header from 'components/layout/header';
import TableApp from 'components/table';
import { NextSeo } from 'next-seo';
import { Badge } from '@chakra-ui/react';
import { UserColumn } from './userColumn';

const data = [
  {
    user: {
      fullname: 'John Doe',
      email: 'johndoe@gmail.com',
    },
    role: 'Assasin',
    status: 'Active',
  },
  {
    user: {
      fullname: 'Doe Doe',
      email: 'doedoe@gmail.com',
    },
    role: 'Mage',
    status: 'Pending',
  },
  {
    user: {
      fullname: 'John John',
      email: 'johnjohn@gmail.com',
    },
    role: 'Support',
    status: 'Inactive',
  },
];

const badgeEnum: Record<string, string> = {
  Active: 'green',
  Pending: 'orange',
  Inactive: 'red',
};

const columns = [
  {
    Header: 'Member',
    accessor: 'user',
    Cell: ({ cell: { value } }) => {
      return <UserColumn data={value} />;
    },
  },
  {
    Header: 'Role',
    accessor: 'role',
  },
  {
    Header: 'Status',
    accessor: 'status',
    Cell: ({ cell: { value } }) => {
      return (
        <Badge fontSize="xs" colorScheme={badgeEnum[value]}>
          {value}
        </Badge>
      );
    },
  },
];

export default function User() {
  return (
    <Container>
      <NextSeo title="User" description="App User" />
      <Header title="User" withAddButton withSearchBox />
      <TableApp columns={columns} data={data} />
    </Container>
  );
}
