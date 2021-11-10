import {
  Avatar,
  Box,
  Flex,
  Divider,
  Menu,
  MenuItem,
  MenuList,
  Text,
  useMenuButton,
  UseMenuButtonProps,
  useColorModeValue as mode,
} from '@chakra-ui/react';
import * as React from 'react';

const UserAvatar = () => <Avatar size="sm" name="Foo Bar" />;

const ProfileMenuButton = (props: UseMenuButtonProps) => {
  const buttonProps = useMenuButton(props);
  return (
    <Flex
      {...buttonProps}
      as="button"
      rounded="full"
      outline="0"
      _focus={{ shadow: 'outline' }}
      display="flex"
      alignItems="center"
    >
      <Box srOnly>Open user menu</Box>
      <UserAvatar />
      <Box lineHeight="1" textAlign="left" pl="2">
        <Text fontWeight="semibold">Foo Bar</Text>
        <Text mt="1" fontSize="xs" color="gray.500">
          foobar@mail.com
        </Text>
      </Box>
    </Flex>
  );
};

export const UserInfo = () => (
  <Menu>
    <ProfileMenuButton />
    <MenuList
      rounded="md"
      shadow="lg"
      color={mode('gray.600', 'inherit')}
      fontSize="sm"
    >
      <MenuItem fontWeight="medium">Switch Account</MenuItem>
      <Divider />
      <MenuItem fontWeight="medium">Feedback & Support</MenuItem>
      <MenuItem fontWeight="medium">Account Settings</MenuItem>
      <Divider />
      <MenuItem fontWeight="medium" color={mode('red.500', 'red.300')}>
        Sign out
      </MenuItem>
    </MenuList>
  </Menu>
);
