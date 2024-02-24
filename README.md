# go-vite

Build small, portable, rich apps.

* go 1.21
* node 1.18

# dev

```bash
make dev
```

Starts a development server.

* Vite Server http://localhost:5173/ (messages swallowed for clean)
* Development Server http://localhost:8080


* Install protoc-js 
```
npm install -g protoc-gen-js
```

# production

`make build` will produce a single executible containing the `vite` app, with the golang backed api server.

```bash
./myapp
```

## build

```
Usage of ./myapp:
  -headless
        Run in headless mode without opening the browser
  -port int
        Run in headless mode without opening the browser
```