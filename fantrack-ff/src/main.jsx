import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import App from './App.jsx'
import './index.css'
import WinningMatchup from './components/winningmatchup.jsx'




const router = createBrowserRouter([
  {
    path: "/tester",
    element: <WinningMatchup />
  },
  {
    path: "/",
    element:  null
  }
])

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <div>
      <App />
    <RouterProvider router={router} />
    </div>
  </StrictMode>,
)
