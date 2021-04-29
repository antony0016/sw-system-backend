docker build -t sw-system-backend .

docker run -d --name sw-system-backend -p 3000:3000 sw-system-backend

docker compose up -d