import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { PrimeReactProvider } from 'primereact/api'
import { Button } from 'primereact/button'; 
import "primereact/resources/themes/lara-light-cyan/theme.css";
import TemplateDemo from './components/menubar'
import { getWinners, getdata } from './api/winners'
import AuthButton from './components/authbutton'
function App() {



  return (
    <>
    <PrimeReactProvider>
      
        <TemplateDemo></TemplateDemo>
        


        <AuthButton></AuthButton>
        <button onClick={() => getdata()}>egg</button>
  
 
      </PrimeReactProvider>
    </>
  )
}

export default App
