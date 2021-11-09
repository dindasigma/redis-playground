import Container from 'components/container';
import { NextSeo } from 'next-seo';

export default function App() {
  return (
    <Container>
      <NextSeo title="Index" description="App Dashboard" />
      Index
    </Container>
  );
}
