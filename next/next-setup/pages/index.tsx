import Container from 'components/container';
import Header from 'components/layout/header';
import { NextSeo } from 'next-seo';

export default function App() {
  return (
    <Container>
      <NextSeo title="Index" description="App Dashboard" />
      <Header title="Home" />
    </Container>
  );
}
