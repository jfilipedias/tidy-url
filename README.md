# Tidy URL

A URL shortener micro SaaS built with Go, chi, templ, Tailwind and MongoDB.

## Getting stated

To run the project you need Go 1.24.0.

Install the tools need to run the project:
```sh
go mod download
```

Install the Tailwind [standalone executable](https://github.com/tailwindlabs/tailwindcss/releases):
```sh
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mkdir bin
mv tailwindcss-macos-arm64 bin/tailwindcss
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
