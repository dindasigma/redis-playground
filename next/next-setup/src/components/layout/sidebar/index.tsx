import {
  Box,
  Drawer,
  DrawerOverlay,
  DrawerCloseButton,
  DrawerHeader,
  DrawerBody,
  DrawerContent,
  Stack,
} from '@chakra-ui/react';
import { BiHome, BiUserCircle } from 'react-icons/bi';
import { Logo } from './logo';
import { NavItem } from './menuItem';

interface Props {
  onClose: Function;
  isOpen: boolean;
  variant: 'drawer' | 'sidebar';
}

export const routes = [
  { label: 'Home', href: '/', icon: BiHome },
  { label: 'User', href: '/user', icon: BiUserCircle },
];

const SidebarContent = ({ onClick }: { onClick: Function }) => (
  <Stack spacing="1">
    {routes.map((props, rid) => (
      <NavItem {...props} key={`nav-item-${rid}`} />
    ))}
  </Stack>
);

const Sidebar = ({ isOpen, variant, onClose }: Props) => {
  return variant === 'sidebar' ? (
    <>
      <Box
        position="fixed"
        left={0}
        p={5}
        w="200px"
        top={0}
        h="100%"
        bg="teal.600"
      >
        <Logo color="gray.50" h="6" mb="6" mt="1" />
        <SidebarContent onClick={onClose} />
      </Box>
    </>
  ) : (
    <Drawer isOpen={isOpen} placement="left" onClose={onClose}>
      <DrawerOverlay>
        <DrawerContent bg="teal.600">
          <DrawerCloseButton />
          <DrawerHeader>
            <Logo color="gray.50" h="6" mt="5" />
          </DrawerHeader>
          <DrawerBody>
            <SidebarContent onClick={onClose} />
          </DrawerBody>
        </DrawerContent>
      </DrawerOverlay>
    </Drawer>
  );
};

export default Sidebar;
