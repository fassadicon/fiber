import React, { useState } from 'react'
import api from '../services/client'

export default function Auth(){
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [message, setMessage] = useState<string | null>(null)

  async function doLogin(e: React.FormEvent){
    e.preventDefault()
    try{
      const res = await api.post('/auth/login', { email, password })
      // expected token in response (depending on your backend)
      const token = res.data?.token || res.data?.data?.token
      if(token){
        localStorage.setItem('token', token)
        api.defaults.headers.common['Authorization'] = `Bearer ${token}`
        setMessage('Login successful, token stored')
      }else{
        setMessage(JSON.stringify(res.data))
      }
    }catch(err:any){
      setMessage(err?.response?.data || String(err))
    }
  }

  async function doRegister(e: React.FormEvent){
    e.preventDefault()
    try{
      const res = await api.post('/auth/register', { email, password })
      setMessage('Register response: ' + JSON.stringify(res.data))
    }catch(err:any){
      setMessage(err?.response?.data || String(err))
    }
  }

  return (
    <div>
      <h2>Auth</h2>
      <form onSubmit={doLogin} className="card">
        <label>Email</label>
        <input value={email} onChange={e => setEmail(e.target.value)} />
        <label>Password</label>
        <input value={password} onChange={e => setPassword(e.target.value)} type="password" />
        <div className="row">
          <button type="submit">Login</button>
          <button onClick={doRegister} style={{marginLeft:10}}>Register</button>
        </div>
      </form>

      <pre>{message}</pre>
    </div>
  )
}
