const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080'

async function request(path, {method='GET', body, token} = {}){
  const headers = {'Content-Type':'application/json'}
  if(token) headers['Authorization'] = `Bearer ${token}`
  const opts = {method, headers}
  if(body) opts.body = JSON.stringify(body)

  const res = await fetch(`${API_BASE}${path}`, opts)
  let data
  try{ data = await res.json() }catch(e){ data = await res.text() }
  if(!res.ok){
    const err = new Error(data?.message || res.statusText || 'Request failed')
    err.response = data
    throw err
  }

  // The Go backend wraps responses as { success, status, message, data }
  // Unwrap and return the `data` field when present so callers get the payload directly.
  if (data && Object.prototype.hasOwnProperty.call(data, 'data')) {
    return data.data
  }

  return data
}

export async function login({email, password}){
  return request('/api/auth/login', {method:'POST', body:{email, password}})
}

export async function register(payload){
  // payload should include: email, first_name, last_name, password, phone_number (optional)
  return request('/api/auth/register', {method:'POST', body: payload})
}

export async function getProducts(token){
  return request('/api/product', {token})
}

export async function getNews(token){
  return request('/api/news', {token})
}

export async function getUsers(token){
  return request('/api/user', {token})
}

export default {login, register, getProducts, getNews, getUsers}
