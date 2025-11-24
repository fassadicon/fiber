import axios from 'axios'

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080'

async function request(path, {method='GET', body, token, signal} = {}){
  const opts = {
    url: `${API_BASE}${path}`,
    method,
    headers: {},
    data: body,
  }
  if(token) opts.headers.Authorization = `Bearer ${token}`
  if(signal) opts.signal = signal

  try{
    const res = await axios(opts)
    const data = res.data
    // If response contains pagination meta, return both data and paginate
    if (data && (Object.prototype.hasOwnProperty.call(data, 'paginate') || Object.prototype.hasOwnProperty.call(data, 'Paginate'))) {
      return { data: data.data, paginate: data.paginate || data.Paginate }
    }

    if (data && Object.prototype.hasOwnProperty.call(data, 'data')) {
      return data.data
    }

    return data
  }catch(e){
    const msg = e.response?.data?.message || e.message || 'Request failed'
    const err = new Error(msg)
    err.response = e.response?.data || null
    throw err
  }
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
// Note: Product CRUD removed per request

// Client CRUD helpers (client GUID-based in backend)
export async function listClients({token, page=1, per_page=10, keyword='', signal } = {}){
  const q = new URLSearchParams()
  if(page) q.set('page', page)
  if(per_page) q.set('limit', per_page)
  if(keyword) q.set('keyword', keyword)
  return request('/api/client?' + q.toString(), {token, signal})
}

export async function getClient(guid, token){
  return request(`/api/client/${guid}`, {token})
}

export async function createClient(payload, token){
  return request('/api/client', {method:'POST', body: payload, token})
}

export async function updateClient(guid, payload, token){
  return request(`/api/client/${guid}`, {method:'PUT', body: payload, token})
}

export async function deleteClient(guid, token){
  return request(`/api/client/${guid}`, {method:'DELETE', token})
}

export default {login, register, getProducts, getNews, getUsers,
  listClients, getClient, createClient, updateClient, deleteClient
}
