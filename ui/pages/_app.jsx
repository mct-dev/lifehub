// import App from 'next/app'
import { useEffect, useState } from 'react'
import { handleKeyCombo } from '../commands/keyboard'

// eslint-disable-next-line react/prop-types
function MyApp ({ Component, pageProps }) {
  const [activeKeys, setActiveKeys] = useState([])

  useEffect(() => {
    document.addEventListener('keydown', handleKeyDown)
    document.addEventListener('keyup', handleKeyUp)

    return () => {
      document.removeEventListener('keydown', handleKeyDown)
      document.removeEventListener('keyup', handleKeyUp)
    }
  })

  const handleKeyDown = (ev) => {
    if (ev.key !== activeKeys.slice(activeKeys.length - 1)) {
      setActiveKeys([...activeKeys, ev.key])
    }
  }

  const handleKeyUp = (ev) => {
    if (activeKeys.length > 1) {
      handleKeyCombo(activeKeys)
    }

    setActiveKeys([])
  }

  return <Component {...pageProps} />
}

// Only uncomment this method if you have blocking data requirements for
// every single page in your application. This disables the ability to
// perform automatic static optimization, causing every page in your app to
// be server-side rendered.
//
// MyApp.getInitialProps = async (appContext) => {
//   // calls page's `getInitialProps` and fills `appProps.pageProps`
//   const appProps = await App.getInitialProps(appContext);
//
//   return { ...appProps }
// }

export default MyApp
