import React, { useState } from 'react'
import Home from './pages/Home'
import Auth from './pages/Auth'

export default function App() {
  const [route, setRoute] = useState<'home'|'auth'>('home')

  return (
    <div className="app">
      <header>
        <h1>Fiber Frontend</h1>
        <nav>
          <button onClick={() => setRoute('home')}>Home</button>
          <button onClick={() => setRoute('auth')}>Auth</button>
        </nav>
      </header>
      <main>
        {route === 'home' ? <Home /> : <Auth />}
      </main>
    </div>
  )
}
