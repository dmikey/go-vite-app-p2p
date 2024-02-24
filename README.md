# go-vite

Starter repo to build small, portable, rich apps, which are by default Web Apps. 

This repo contains a starter base, which will help create a React/Typescript/Vite Front End, that is powered by a Golang backend. Tied together using google `proto3`. 

You can simply distribute your binary application, which will contain a front end, and the webserver to power it together. It wil automatically open a user's Web Browser as it's primary UI Window.

Using `--headless` and `--port` you can specify some options, to transport or build and distribute this as a working web app where you need it.

In addition there are (will be) a number of associated "containers" which all follow the same RPC interface, and allow you to distribute ✌ Native ✌ apps, so users get the easy of ownership with your product.

# dev

To build the basics : 

* go 1.21
* node 1.18

To build more : 

* Swift for OSX 14+

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