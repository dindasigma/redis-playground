import {
  Box,
  Button,
  Heading,
  HStack,
  Input,
  InputGroup,
  InputRightElement,
  Stack,
  Text,
  useColorModeValue as mode,
} from '@chakra-ui/react';
import * as React from 'react';
import { FiSearch } from 'react-icons/fi';

interface HeaderProps {
  title: string;
  withSearchBox: Boolean;
  withAddButton: Boolean;
}

export default function Header({
  title,
  withSearchBox,
  withAddButton,
}: HeaderProps) {
  return (
    <Stack
      spacing="5"
      direction={{ base: 'column', md: 'row' }}
      justify="space-between"
      align={{ base: 'flex-start', md: 'center' }}
      pb="10"
    >
      <Stack>
        <Heading size="lg">{title}</Heading>
      </Stack>

      <HStack
        justify="flex-end"
        flex="1"
        w={{ base: 'full', md: 'auto' }}
        spacing={{ base: '2', md: '4' }}
      >
        {withSearchBox && (
          <InputGroup maxW={{ md: '80' }} w="full">
            <InputRightElement color="gray.400">
              <FiSearch />
            </InputRightElement>
            <Input
              bg={mode('white', 'gray.800')}
              placeholder={`Search for ${title}`}
              fontSize="sm"
            />
          </InputGroup>
        )}
        {withAddButton && (
          <Button
            bg="teal.600"
            color="gray.50"
            flexShrink={0}
            fontWeight="bold"
            fontSize="sm"
          >
            Add New {title}
          </Button>
        )}
      </HStack>
    </Stack>
  );
}
