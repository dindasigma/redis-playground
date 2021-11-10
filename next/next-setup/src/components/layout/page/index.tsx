import { Box } from '@chakra-ui/layout';
import { ReactJSXElementChildrenAttribute } from '@emotion/react/types/jsx-namespace';
import React from 'react';

const Page = (props: { children: ReactJSXElementChildrenAttribute }) => {
  return (
    <Box p="8">
      <Box maxW="7xl" mx="auto">
        {props.children}
      </Box>
    </Box>
  );
};

export default Page;
