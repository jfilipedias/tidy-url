# Tidy URL

A URL shortener micro SaaS built with Go, chi, templ, Tailwind and MongoDB.

## Getting stated

To run the project you need Go 1.24.0, or above, to run the project, and Node.js 20.11.0 and npm 10.5.0 to generate the CSS through Tailwind CLI.

Install the tools need to run the project:
```sh
go mod download
npm install
```

To run the project you will need to run 3 separeted process:

1. Generate the CSS files using the Tailwind CLI in watch mode:
```sh
make styles
```

2. Generate the templ files in watch mode:
```sh
make views
```

3. Run the application using air:
```sh
make run
```

## Tech Stack

- [Chi](https://go-chi.io/): A lightweight, idiomatic and composable router for building Go HTTP services.
- [templ](https://templ.guide/): A language for writing HTML user interfaces in Go.
- [Tailwind](https://tailwindcss.com/): A utility-first CSS framework
- [MongoDB](https://www.mongodb.com/): A NoSQL document database designed for ease of application development and scaling.
