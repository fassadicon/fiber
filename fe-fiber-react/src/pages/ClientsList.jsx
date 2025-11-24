import React from 'react'
import { Link } from 'react-router-dom'
import PaginatedList from '../components/PaginatedList'
import { listClients, deleteClient } from '../services/client'

export default function ClientsList({token}){
  async function fetcher({page, per_page, keyword}){
    return listClients({token, page, per_page, keyword})
  }

  function renderItem(c){
    return (
      <div>
        <strong>{c.name || `${c.first_name || ''} ${c.last_name || ''}`}</strong>
        <div style={{fontSize:12, color:'#666'}}>{c.email || c.phone_number}</div>
      </div>
    )
  }

  function actions(c){
    return (
      <div>
        <Link to={`/clients/${c.guid}`}>Edit</Link>
        <button onClick={async ()=>{
          if(!confirm('Delete this client?')) return
          try{
            await deleteClient(c.guid, token)
            alert('Deleted')
            window.location.reload()
          }catch(err){ alert(err.message || String(err)) }
        }} style={{marginLeft:8}}>Delete</button>
      </div>
    )
  }

  return (
    <div>
      <h2>Clients</h2>
      <div style={{marginBottom:8}}>
        <Link to="/clients/new"><button>Create client</button></Link>
      </div>
      <PaginatedList fetcher={fetcher} renderItem={renderItem} actions={actions} />
    </div>
  )
}
