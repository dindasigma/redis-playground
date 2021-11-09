import { Box, HStack, LinkOverlay, LinkBox, Text } from '@chakra-ui/react';
import * as React from 'react';
import { ChevronRightIcon } from '@chakra-ui/icons';
import { IconType } from 'react-icons';
import Icon from '@chakra-ui/icon';
import { useRouter } from 'next/router';
import NextLink from 'next/link';

interface NavItemProps {
  href?: string;
  label: string;
  icon: IconType;
}

export const NavItem = (props: NavItemProps) => {
  const { icon, label, href } = props;
  const router = useRouter();

  return (
    <LinkBox>
      <HStack
        aria-current={router.pathname === props.href ? 'page' : undefined}
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
        <Box fontSize="lg">
          <Icon as={icon} />
        </Box>
        <NextLink href={href || ''} passHref>
          <LinkOverlay>
            <Text>{label}</Text>
          </LinkOverlay>
        </NextLink>
      </HStack>
    </LinkBox>
  );
};
