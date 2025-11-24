import React, {useState} from 'react'
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import Auth from './pages/Auth'
import Home from './pages/Home'
import ClientsList from './pages/ClientsList'
import ClientForm from './pages/ClientForm'

export default function App(){
  const [token, setToken] = useState(localStorage.getItem('token') || '')

  function onLogin(t){
    setToken(t)
    localStorage.setItem('token', t)
  }

  function onLogout(){
    setToken('')
    localStorage.removeItem('token')
  }

  return (
    <BrowserRouter>
      <div className="app">
        <header>
          <h1>FE Fiber React (JS)</h1>
          <nav>
            <Link to="/auth"><button>Auth</button></Link>
            <Link to="/home"><button>Home</button></Link>
            <Link to="/clients"><button>Clients</button></Link>
            {token ? <button onClick={onLogout}>Logout</button> : null}
          </nav>
        </header>

        <main>
          <Routes>
            <Route path="/" element={<Home token={token} />} />
            <Route path="/home" element={<Home token={token} />} />
            <Route path="/auth" element={<Auth onLogin={onLogin} />} />

            <Route path="/clients" element={<ClientsList token={token} />} />
            <Route path="/clients/new" element={<ClientForm token={token} />} />
            <Route path="/clients/:id" element={<ClientForm token={token} />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  )
}
