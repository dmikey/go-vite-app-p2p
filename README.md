# go-vite

Starter repo to build small, portable, rich apps.

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

## app containers

This repo comes with some projects and developer workflows to help distribute ✌ Native ✌ apps. 

* OSX via Swift UI (available)
* iOS via Swift UI (coming)
* Android via Kotlin (coming)
* Windows (coming)
* GTK (coming)
* QT (coming)