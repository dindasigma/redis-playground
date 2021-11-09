import { Box } from '@chakra-ui/layout';
import { ReactJSXElementChildrenAttribute } from '@emotion/react/types/jsx-namespace';
import React from 'react';

const Page = (props: { children: ReactJSXElementChildrenAttribute }) => {
  return (
    <Box w="full" px="4" py="4">
      {props.children}
    </Box>
  );
};

export default Page;
