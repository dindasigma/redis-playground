import {
  Menu,
  MenuItem,
  MenuButton,
  MenuList,
  IconButton,
} from '@chakra-ui/react';
import { BiDotsVerticalRounded } from 'react-icons/bi';

export default function tableActions() {
  return (
    <Menu>
      <MenuButton
        as={IconButton}
        aria-label="Options"
        icon={<BiDotsVerticalRounded />}
        variant="outline"
      />
      <MenuList>
        <MenuItem>View</MenuItem>
        <MenuItem>Delete</MenuItem>
        <MenuItem>Close</MenuItem>
      </MenuList>
    </Menu>
  );
}
