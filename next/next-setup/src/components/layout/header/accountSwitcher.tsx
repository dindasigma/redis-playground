import {
  Menu,
  MenuItemOption,
  MenuList,
  MenuOptionGroup,
  useColorModeValue,
} from '@chakra-ui/react';
import * as React from 'react';
import { AccountSwitcherButton } from './accountSwitcherButton';

export const AccountSwitcher = () => {
  return (
    <Menu>
      <AccountSwitcherButton />
      <MenuList
        shadow="lg"
        py="4"
        color={useColorModeValue('gray.600', 'gray.200')}
        px="3"
        fontSize="sm"
      >
        <MenuOptionGroup defaultValue="chakra-ui">
          <MenuItemOption value="chakra-ui" fontWeight="semibold" rounded="md">
            Project 1
          </MenuItemOption>
          <MenuItemOption value="careerlyft" fontWeight="semibold" rounded="md">
            Project 2
          </MenuItemOption>
        </MenuOptionGroup>
      </MenuList>
    </Menu>
  );
};
