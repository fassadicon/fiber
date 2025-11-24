import React, {useEffect, useState} from 'react'

// Props:
// - fetcher: async ({page, per_page, keyword, filters}) => {items, meta} OR array
// - renderItem: (item) => ReactNode
// - perPageOptions: [10,25]
// - initialPerPage
// - actions: optional function(item) => buttons
export default function PaginatedList({fetcher, renderItem, perPageOptions=[10,25,50], initialPerPage=10, actions}){
  const [items, setItems] = useState([])
  const [page, setPage] = useState(1)
  const [perPage, setPerPage] = useState(initialPerPage)
  const [keyword, setKeyword] = useState('')
  const [total, setTotal] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  useEffect(()=>{ load() }, [page, perPage])

  async function load(){
    setLoading(true)
    setError(null)
    try{
      const res = await fetcher({page, per_page: perPage, keyword})
      if(Array.isArray(res)){
        setItems(res)
        setTotal(res.length)
      } else if(res && res.data){
        setItems(res.data)
        setTotal(res.meta?.total ?? res.meta?.total_records ?? null)
      } else if(res && res.items){
        setItems(res.items)
        setTotal(res.total ?? null)
      } else {
        // unknown shape, try to use as array
        setItems(res ?? [])
        setTotal((res && res.length) || null)
      }
    }catch(err){
      setError(err.message || String(err))
    }finally{
      setLoading(false)
    }
  }

  function onSearch(e){
    e.preventDefault()
    setPage(1)
    load()
  }

  const totalPages = total ? Math.max(1, Math.ceil(total / perPage)) : null

  return (
    <div>
      <form onSubmit={onSearch} style={{display:'flex', gap:8, marginBottom:12}}>
        <input placeholder="Search..." value={keyword} onChange={e=>setKeyword(e.target.value)} />
        <button type="submit">Search</button>
        <label style={{marginLeft:8}}>
          Per page:
          <select value={perPage} onChange={e=>{ setPerPage(Number(e.target.value)); setPage(1) }} style={{marginLeft:6}}>
            {perPageOptions.map(p => <option key={p} value={p}>{p}</option>)}
          </select>
        </label>
      </form>

      {loading ? <div>Loading...</div> : null}
      {error ? <div style={{color:'red'}}>{error}</div> : null}

      <div>
        {items.length === 0 ? <div>No items</div> : (
          <ul style={{paddingLeft:0, listStyle:'none'}}>
            {items.map(it => (
              <li key={it.id || it.guid} style={{borderBottom:'1px solid #eee', padding:'8px 0', display:'flex', justifyContent:'space-between', alignItems:'center'}}>
                <div>{renderItem(it)}</div>
                <div style={{marginLeft:12}}>{actions ? actions(it) : null}</div>
              </li>
            ))}
          </ul>
        )}
      </div>

      <div style={{marginTop:12, display:'flex', alignItems:'center', gap:8}}>
        <button onClick={()=>setPage(1)} disabled={page===1 || loading}>First</button>
        <button onClick={()=>setPage(p=>Math.max(1,p-1))} disabled={page===1 || loading}>Prev</button>
        <div>Page {page}{totalPages ? ` / ${totalPages}` : ''}</div>
        <button onClick={()=>setPage(p=>p+1)} disabled={totalPages && page>=totalPages || loading}>Next</button>
        <button onClick={()=>setPage(totalPages || 1)} disabled={totalPages && page>=totalPages || loading}>Last</button>
        {total !== null ? <div style={{marginLeft:12}}>Total: {total}</div> : null}
      </div>
    </div>
  )
}
