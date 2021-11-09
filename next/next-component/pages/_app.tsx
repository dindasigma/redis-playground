import type { AppProps } from 'next/app'
import NavBar from "../components/NavBar";

function MyApp({ Component, pageProps }: AppProps) {
  console.log('[App] render');
  return (
    <>
      <header>
        <NavBar />
      </header>
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
