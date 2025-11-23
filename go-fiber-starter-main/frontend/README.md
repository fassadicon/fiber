# Frontend (Vite + React)

This is a minimal React + Vite app intended to be used together with the Go Fiber backend in the repository root.

Goals:
- Provide a small UI to call the Go API endpoints
- Use Vite dev server proxy so `/api/*` requests go to `http://127.0.0.1:8080`

Quick start (PowerShell):

```powershell
cd frontend
npm install
npm run dev
```

- The dev server runs on `http://localhost:3000` by default and proxies `/api` to `http://127.0.0.1:8080`.
- Use the Home page to call `GET /api/import/products/stream` (public) and `GET /api/product` (requires auth token).
- Use the Auth page to `POST /api/auth/login` and `POST /api/auth/register` (whitelisted in the backend).

Notes:
- The backend must be running (e.g. `go run main.go --rollback --seed`) and listening on `127.0.0.1:8080` for the proxy to work.
- When you login, the app attempts to store a token from the login response in `localStorage` and uses it for subsequent requests.

If you'd like, I can:
- Add `react-router` and more pages
- Implement a nicer product listing and editing UI
- Commit the frontend folder into git

