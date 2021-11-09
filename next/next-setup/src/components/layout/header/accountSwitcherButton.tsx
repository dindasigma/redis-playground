import { Box, Flex, FlexProps, HStack, useMenuButton } from '@chakra-ui/react';
import * as React from 'react';
import { HiSelector } from 'react-icons/hi';

export const AccountSwitcherButton = (props: FlexProps) => {
  const buttonProps = useMenuButton(props);
  return (
    <Flex
      as="button"
      {...buttonProps}
      display="flex"
      alignItems="center"
      fontSize="sm"
      userSelect="none"
      cursor="pointer"
      outline="0"
    >
      <HStack flex="1" spacing="3">
        <Box textAlign="start">
          <Box isTruncated fontWeight="semibold">
            Project 1
          </Box>
        </Box>
      </HStack>
      <Box fontSize="lg" color="gray.400" ml="2">
        <HiSelector />
      </Box>
    </Flex>
  );
};
