import React, {useState} from 'react'
import { getProducts, getNews, getUsers } from '../services/client'

export default function Home({token}){
  const [output, setOutput] = useState('')

  async function callProducts(){
    setOutput('...')
    try{
      const res = await getProducts(token)
      setOutput(JSON.stringify(res, null, 2))
    }catch(err){
      setOutput(err.message || String(err))
    }
  }

  async function callNews(){
    setOutput('...')
    try{
      const res = await getNews(token)
      setOutput(JSON.stringify(res, null, 2))
    }catch(err){
      setOutput(err.message || String(err))
    }
  }

  async function callUsers(){
    setOutput('...')
    try{
      const res = await getUsers(token)
      setOutput(JSON.stringify(res, null, 2))
    }catch(err){
      setOutput(err.message || String(err))
    }
  }

  return (
    <div>
      <h2>API Tester</h2>
      <div style={{marginBottom:8}}>
        <button onClick={callProducts}>GET /api/product</button>
        <button onClick={callNews} style={{marginLeft:8}}>GET /api/news</button>
        <button onClick={callUsers} style={{marginLeft:8}}>GET /api/user</button>
      </div>

      <div>
        <strong>Response:</strong>
        <pre style={{whiteSpace:'pre-wrap', marginTop:8}}>{output}</pre>
      </div>
    </div>
  )
}
