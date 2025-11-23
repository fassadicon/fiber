FE Fiber React (JS)

Minimal Vite + React frontend to test the Go Fiber backend (`go-fiber-starter-main`).

Quick start

1. Install dependencies:

```powershell
cd fe-fiber-react
npm install
```

2. (Optional) Set `VITE_API_BASE` in a `.env` file if your backend is not at `http://localhost:3000`:

```
VITE_API_BASE=http://localhost:3000
```

3. Run the dev server:

```powershell
npm run dev
```

Usage

- Open the dev server (Vite will show the URL, usually `http://localhost:5173`).
- Use the `Auth` tab to `register` or `login` (the Go backend whitelists `/api/auth/login` and `/api/auth/register`).
- After login, the token is stored in `localStorage` and the `Home` tab can call secured endpoints: `/api/product`, `/api/news`, `/api/user`.

Notes

- This is a small testing UI — no TypeScript, no heavy CSS.
- Update `VITE_API_BASE` if your Go server runs on a different host/port.
