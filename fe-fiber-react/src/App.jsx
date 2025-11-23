import React, {useState} from 'react'
import Auth from './pages/Auth'
import Home from './pages/Home'

export default function App(){
  const [page, setPage] = useState('auth')
  const [token, setToken] = useState(localStorage.getItem('token') || '')

  function onLogin(t){
    setToken(t)
    localStorage.setItem('token', t)
    setPage('home')
  }

  function onLogout(){
    setToken('')
    localStorage.removeItem('token')
    setPage('auth')
  }

  return (
    <div className="app">
      <header>
        <h1>FE Fiber React (JS)</h1>
        <nav>
          <button onClick={() => setPage('auth')}>Auth</button>
          <button onClick={() => setPage('home')}>Home</button>
          {token ? <button onClick={onLogout}>Logout</button> : null}
        </nav>
      </header>

      <main>
        {page === 'auth' && <Auth onLogin={onLogin} />}
        {page === 'home' && <Home token={token} />}
      </main>
    </div>
  )
}
