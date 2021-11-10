import { useState } from 'react';
import { Box, useBreakpointValue, ChakraProviderProps } from '@chakra-ui/react';
import Navbar from 'components/layout/navbar';
import Sidebar from 'components/layout/sidebar';
import Page from 'components/layout/page';
import { DefaultSeo } from 'next-seo';

const smVariant = { navigation: 'drawer', navigationButton: true };
const mdVariant = { navigation: 'sidebar', navigationButton: false };

export default function Container({ children }: ChakraProviderProps) {
  const [isSidebarOpen, setSidebarOpen] = useState(false);
  const variants = useBreakpointValue({ base: smVariant, md: mdVariant });

  const toggleSidebar = () => setSidebarOpen(!isSidebarOpen);
  return (
    <>
      <DefaultSeo
        titleTemplate="%s | Envelope"
        defaultTitle="Envelope"
        description="Chakra UI Template"
        additionalLinkTags={[
          {
            rel: 'icon',
            href: '/favicon.ico',
          },
        ]}
      />
      <Sidebar
        variant={variants?.navigation}
        isOpen={isSidebarOpen}
        onClose={toggleSidebar}
      />
      <Box ml={!variants?.navigationButton && 200}>
        <Navbar
          showSidebarButton={variants?.navigationButton}
          onShowSidebar={toggleSidebar}
        />
        <Page>{children}</Page>
      </Box>
    </>
  );
}
