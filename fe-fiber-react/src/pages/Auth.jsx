import React, {useState} from 'react'
import { login, register } from '../services/client'

export default function Auth({onLogin}){
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [phoneNumber, setPhoneNumber] = useState('')
  const [message, setMessage] = useState('')

  async function handleLogin(e){
    e.preventDefault()
    setMessage('...')
    try{
      const res = await login({email, password})
      if(res && res.token){
        setMessage('login successful')
        onLogin(res.token)
      } else {
        setMessage(JSON.stringify(res))
      }
    }catch(err){
      setMessage(err.response?.message || err.message || String(err))
    }
  }

  async function handleRegister(e){
    e.preventDefault()
    setMessage('...')
    try{
      const payload = {
        email,
        password,
        first_name: firstName,
        last_name: lastName,
        phone_number: phoneNumber,
      }
      const res = await register(payload)
      setMessage('registered: ' + JSON.stringify(res))
    }catch(err){
      setMessage(err.response?.message || err.message || String(err))
    }
  }

  return (
    <div>
      <h2>Auth (Login / Register)</h2>
      <form onSubmit={handleLogin}>
        <div>
          <label>Email</label>
          <input value={email} onChange={e=>setEmail(e.target.value)} />
        </div>
        <div>
          <label>Password</label>
          <input type="password" value={password} onChange={e=>setPassword(e.target.value)} />
        </div>

        <hr style={{margin:'12px 0'}} />

        <div>
          <label>First name</label>
          <input value={firstName} onChange={e=>setFirstName(e.target.value)} />
        </div>
        <div>
          <label>Last name</label>
          <input value={lastName} onChange={e=>setLastName(e.target.value)} />
        </div>
        <div>
          <label>Phone (optional)</label>
          <input value={phoneNumber} onChange={e=>setPhoneNumber(e.target.value)} />
        </div>

        <div style={{marginTop:8}}>
          <button type="submit">Login</button>
          <button type="button" onClick={handleRegister} style={{marginLeft:8}}>Register</button>
        </div>
      </form>
      <div style={{marginTop:12}}>
        <strong>Message:</strong>
        <div style={{whiteSpace:'pre-wrap'}}>{message}</div>
      </div>
    </div>
  )
}
