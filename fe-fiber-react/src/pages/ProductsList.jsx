import React from 'react'
import { Link, useNavigate } from 'react-router-dom'
import PaginatedList from '../components/PaginatedList'
import { listProducts, deleteProduct } from '../services/client'

export default function ProductsList({token}){
  const nav = useNavigate()

  async function fetcher({page, per_page, keyword}){
    return listProducts({token, page, per_page, keyword})
  }

  function renderItem(p){
    return (
      <div>
        <strong>{p.name}</strong>
        <div style={{fontSize:12, color:'#666'}}>{p.description}</div>
      </div>
    )
  }

  function actions(p){
    return (
      <div>
        <Link to={`/products/${p.id}`}>Edit</Link>
        <button onClick={async ()=>{
          if(!confirm('Delete this product?')) return
          try{
            await deleteProduct(p.id, token)
            alert('Deleted')
            window.location.reload()
          }catch(err){ alert(err.message || String(err)) }
        }} style={{marginLeft:8}}>Delete</button>
      </div>
    )
  }

  return (
    <div>
      <h2>Products</h2>
      <div style={{marginBottom:8}}>
        <Link to="/products/new"><button>Create product</button></Link>
      </div>
      <PaginatedList fetcher={fetcher} renderItem={renderItem} actions={actions} />
    </div>
  )
}
