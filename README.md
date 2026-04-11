<img width="991" height="1057" alt="Screenshot 2026-03-24 at 5 22 33 PM" src="https://github.com/user-attachments/assets/da886b49-cf70-4e6d-98b3-be29b6549a29" />

in .env file:
`API_URL=http://localhost:3005`

see `/internal/db/connect.go` for me db var info

run 
`go mod tidy`

then
`cd web`
`bun install`

then in 1 tab: (server/docker/go/postgress)
`make dev`

and in another (ts/react/vite)
`bun run dev`
