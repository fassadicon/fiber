import React, {useEffect, useState} from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getClient, createClient, updateClient } from '../services/client'

export default function ClientForm({token}){
  const {id} = useParams()
  const nav = useNavigate()
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [phone, setPhone] = useState('')
  const [message, setMessage] = useState('')

  useEffect(()=>{
    if(id && id !== 'new') load()
  }, [id])

  async function load(){
    try{
      const c = await getClient(id, token)
      setFirstName(c.first_name || '')
      setLastName(c.last_name || '')
      setEmail(c.email || '')
      setPhone(c.phone_number || '')
    }catch(err){ setMessage(err.message || String(err)) }
  }

  async function onSubmit(e){
    e.preventDefault()
    setMessage('...')
    try{
      const payload = { first_name: firstName, last_name: lastName, email, phone_number: phone }
      if(id && id !== 'new'){
        await updateClient(id, payload, token)
        setMessage('Updated')
      } else {
        await createClient(payload, token)
        setMessage('Created')
      }
      setTimeout(()=>nav('/clients'), 600)
    }catch(err){ setMessage(err.response?.message || err.message || String(err)) }
  }

  return (
    <div>
      <h2>{id && id !== 'new' ? 'Edit' : 'Create'} Client</h2>
      <form onSubmit={onSubmit}>
        <div>
          <label>First name</label>
          <input value={firstName} onChange={e=>setFirstName(e.target.value)} />
        </div>
        <div>
          <label>Last name</label>
          <input value={lastName} onChange={e=>setLastName(e.target.value)} />
        </div>
        <div>
          <label>Email</label>
          <input value={email} onChange={e=>setEmail(e.target.value)} />
        </div>
        <div>
          <label>Phone</label>
          <input value={phone} onChange={e=>setPhone(e.target.value)} />
        </div>
        <div style={{marginTop:8}}>
          <button type="submit">Save</button>
          <button type="button" onClick={()=>nav('/clients')} style={{marginLeft:8}}>Cancel</button>
        </div>
      </form>
      <div style={{marginTop:12}}>{message}</div>
    </div>
  )
}
