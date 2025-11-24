import React, {useEffect, useState, useRef} from 'react'

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
  const [keywordDebounced, setKeywordDebounced] = useState('')
  const [total, setTotal] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  // Fetch when page/perPage/debounced keyword changes
  const controllerRef = useRef(null)

  useEffect(()=>{ fetchPage() }, [page, perPage, keywordDebounced])

  // Debounce keyword input to avoid spamming backend
  useEffect(()=>{
    const t = setTimeout(()=>{
      // reset to first page on new search
      setPage(1)
      setKeywordDebounced(keyword)
    }, 350)

    return ()=> clearTimeout(t)
  }, [keyword])

  async function fetchPage(){
    setLoading(true)
    setError(null)
    try{
      // cancel previous request if any
      if(controllerRef.current){
        try{ controllerRef.current.abort() }catch(e){}
      }
      const controller = new AbortController()
      controllerRef.current = controller

      const res = await fetcher({page, limit: perPage, keyword: keywordDebounced, signal: controller.signal})
      if(Array.isArray(res)){
        setItems(res)
        setTotal(res.length)
      } else if(res && res.data && res.paginate){
        setItems(res.data)
        setTotal(res.paginate.total_rows ?? res.paginate.total ?? null)
      } else if(res && res.data){
        setItems(res.data)
        setTotal((res.data && res.data.length) || null)
      } else if(res && res.items){
        setItems(res.items)
        setTotal(res.total ?? null)
      } else {
        // unknown shape, try to use as array
        setItems(res ?? [])
        setTotal((res && res.length) || null)
      }
    }catch(err){
      // If request was aborted, ignore the error
      if (err.name === 'CanceledError' || err.name === 'AbortError') {
        return
      }
      setError(err.message || String(err))
    }finally{
      setLoading(false)
      // clear controller after completion
      controllerRef.current = null
    }
  }

  function onSearch(e){
    e.preventDefault()
    // immediate search when user submits form
    setPage(1)
    setKeywordDebounced(keyword)
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
