import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import App from './App.jsx'
import './index.css'
import './App.css'
import "primereact/resources/themes/lara-light-indigo/theme.css";
import Leaders from './components/leaders.jsx'
import { getCategoryMap, getLeaderData, getWinningMatchup } from "./api/winners";


const router = createBrowserRouter([
  {
    path: "/tester",
    element: <Leaders data={await getWinningMatchup()} />
  },
  {
    path: "/leader",
    element: <Leaders data={await getLeaderData()} />
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
