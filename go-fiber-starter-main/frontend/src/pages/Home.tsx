import React, { useState } from 'react'
import api from '../services/client'

export default function Home(){
  const [streamMsg, setStreamMsg] = useState<string | null>(null)
  const [products, setProducts] = useState<any[]>([])
  const [loading, setLoading] = useState(false)

  async function fetchStream(){
    try{
      setLoading(true)
      const res = await api.get('/import/products/stream')
      setStreamMsg(JSON.stringify(res.data))
    }catch(err:any){
      setStreamMsg(err?.response?.data || String(err))
    }finally{ setLoading(false) }
  }

  async function fetchProducts(){
    try{
      setLoading(true)
      const res = await api.get('/product')
      setProducts(res.data || [])
    }catch(err:any){
      console.error(err)
      setProducts([])
    }finally{ setLoading(false) }
  }

  return (
    <div>
      <h2>Home / API Tester</h2>
      <p>Backend API proxy: calls to <code>/api/*</code> are proxied to <code>http://127.0.0.1:8080</code>.</p>

      <div className="card">
        <h3>Import stream (public)</h3>
        <button onClick={fetchStream} disabled={loading}>Call /api/import/products/stream</button>
        <pre>{streamMsg}</pre>
      </div>

      <div className="card">
        <h3>List products (requires auth)</h3>
        <button onClick={fetchProducts} disabled={loading}>Call /api/product</button>
        <pre>{JSON.stringify(products, null, 2)}</pre>
      </div>
    </div>
  )
}
