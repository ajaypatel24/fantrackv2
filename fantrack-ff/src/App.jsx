
import { PrimeReactProvider } from 'primereact/api'
import TemplateDemo from './components/menubar'
import AuthButton from './components/authbutton'
import Tailwind from 'primereact/passthrough/tailwind';
function App() {



  return (
    <>
    <PrimeReactProvider value={{unstyled: true, pt: Tailwind}}>

        <TemplateDemo></TemplateDemo>
        


        <AuthButton></AuthButton>

  
 
      </PrimeReactProvider>
    </>
  )
}

export default App
