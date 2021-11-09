import { Box, HStack } from '@chakra-ui/react';
import * as React from 'react';
import { ChevronRightIcon } from '@chakra-ui/icons';

interface NavItemProps {
  href?: string;
  label: string;
  active?: boolean;
  icon: React.ReactElement;
  endElement?: React.ReactElement;
  children?: React.ReactNode;
}

export const NavItem = (props: NavItemProps) => {
  const { active, icon, children, label, endElement, href } = props;
  return (
    <HStack
      as="a"
      href={href}
      aria-current={active ? 'page' : undefined}
      spacing="2"
      px="3"
      py="2"
      rounded="md"
      transition="all 0.2s"
      cursor="pointer"
      color="gray.200"
      _hover={{ bg: 'whiteAlpha.200' }}
      _activeLink={{ bg: 'blackAlpha.300', color: 'white' }}
    >
      <Box fontSize="lg">{icon}</Box>
      <Box flex="1" fontWeight="inherit">
        {label}
      </Box>
      {endElement && !children && <Box>{endElement}</Box>}
      {children && <Box fontSize="xs" flexShrink={0} as={ChevronRightIcon} />}
    </HStack>
  );
};
