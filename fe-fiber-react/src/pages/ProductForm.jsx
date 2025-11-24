import React, {useEffect, useState} from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getProduct, createProduct, updateProduct } from '../services/client'

export default function ProductForm({token}){
  const {id} = useParams()
  const nav = useNavigate()
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [price, setPrice] = useState('')
  const [message, setMessage] = useState('')

  useEffect(()=>{
    if(id && id !== 'new') load()
  }, [id])

  async function load(){
    try{
      const p = await getProduct(id, token)
      setName(p.name||'')
      setDescription(p.description||'')
      setPrice(p.price || '')
    }catch(err){ setMessage(err.message || String(err)) }
  }

  async function onSubmit(e){
    e.preventDefault()
    setMessage('...')
    try{
      const payload = { name, description, price }
      if(id && id !== 'new'){
        await updateProduct(id, payload, token)
        setMessage('Updated')
      } else {
        await createProduct(payload, token)
        setMessage('Created')
      }
      setTimeout(()=>nav('/products'), 600)
    }catch(err){ setMessage(err.response?.message || err.message || String(err)) }
  }

  return (
    <div>
      <h2>{id && id !== 'new' ? 'Edit' : 'Create'} Product</h2>
      <form onSubmit={onSubmit}>
        <div>
          <label>Name</label>
          <input value={name} onChange={e=>setName(e.target.value)} />
        </div>
        <div>
          <label>Description</label>
          <textarea value={description} onChange={e=>setDescription(e.target.value)} />
        </div>
        <div>
          <label>Price</label>
          <input value={price} onChange={e=>setPrice(e.target.value)} />
        </div>
        <div style={{marginTop:8}}>
          <button type="submit">Save</button>
          <button type="button" onClick={()=>nav('/products')} style={{marginLeft:8}}>Cancel</button>
        </div>
      </form>
      <div style={{marginTop:12}}>{message}</div>
    </div>
  )
}
