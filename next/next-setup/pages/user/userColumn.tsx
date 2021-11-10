import { Box, Avatar, Stack } from '@chakra-ui/react';
import * as React from 'react';

interface UserProps {
  data: {
    fullname: string;
    email: string;
  };
}

export const UserColumn = (props: UserProps) => {
  const { fullname, email } = props.data;
  return (
    <Stack direction="row" spacing="4" align="center">
      <Avatar name={`${fullname}`} />
      <Box>
        <Box fontSize="sm" fontWeight="medium">
          {fullname}
        </Box>
        <Box fontSize="sm" color="gray.500">
          {email}
        </Box>
      </Box>
    </Stack>
  );
};
