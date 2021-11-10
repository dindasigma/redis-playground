import { IconButton, Flex, Spacer } from '@chakra-ui/react';
import { HamburgerIcon } from '@chakra-ui/icons';
import { UserInfo } from './userInfo';
import { AccountSwitcher } from './accountSwitcher';

interface Props {
  onShowSidebar: Function;
  showSidebarButton?: boolean;
}

const Navbar = ({ showSidebarButton = true, onShowSidebar }: Props) => {
  return (
    <Flex bg="gray.50" p={4} color="gray.700" borderBottomWidth="1px">
      {showSidebarButton && (
        <IconButton
          icon={<HamburgerIcon />}
          color="gray.400"
          onClick={onShowSidebar}
          mr="5"
        />
      )}
      <AccountSwitcher />
      <Spacer />
      <UserInfo />
    </Flex>
  );
};

export default Navbar;
