# Trip Tracker Backend

GO Version: 1.23.5

This backend application is written in Go and provides API endpoints for managing trips. The backend uses the built-in `net/http` package for handling HTTP requests and maintains a map of trips in memory.

## Resources

We use the aviationstack API for real time flight/airport information [Check their site here](https://aviationstack.com)

[API Docs](https://aviationstack.com/documentation)

---

## Prerequisites

1. **Go Installed**: Ensure Go is installed on your system. You can download it from [here](https://golang.org/dl/).

   - Verify installation:
     ```bash
     go version
     ```

2. **Environment**: You can set an optional `PORT` environment variable to specify the port number on which the backend will run.

---

## Database

I use sqlite for this project

### Installation

Assuming you're working on an Ubuntu 22.04, but this process is similar for most OS cases.

```bash
sudo apt update
```

```bash
sudo apt install sqlite3
```

Verify the installation

```bash
sqlite3 --version
```

Creates the shell and new database. I like to put this under the database folder

```bash
sqlite3 database.db
```

You can also load the schema.sql, which is store under the database folder

```bash
sqlite3 database.db < database/schema.sql
```

## Installing Tailwind

To generate the Tailwind style sheet, we use the Tailwind binary. To get started with TailWind CSS, make sure you have the correct binary in the root directory. follow the instructions in this guide. Make sure you download the correct binary for your operating system. https://tailwindcss.com/blog/standalone-cli

Generating the output.css file

```bash
./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
```

Add the href to the templ files

```html
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>My Trips</title>
  <script src="/static/js/htmx.min.js"></script>
  <link rel="stylesheet" href="/static/css/output.css" />
</head>
```

## Installing Templ

https://templ.guide/

Generate the files via

```bash
templ generate
```

## Installing Air

Air is useful to autoload and track your go code

https://github.com/cosmtrek/air

Refere to [Here](https://github.com/air-verse/air) for reference

First, init air if not already doe

```bash
air init
```

That wil create the .air.toml file. Then just run the air command.

```bash
air
```

## HTMX Example

```html
<button hx-get="/trips" hx-target="#trip-list" hx-swap="innerHTML">
  Load Trips
</button>

<div id="trip-list">
  <ul>
    for _, trip := range trips {
    <li>{ trip.Airline }</li>
    }
  </ul>
</div>
```

## How to Run the Backend

1. **Clone the Repository**:

```bash
git clone <repository_url>
cd trip-tracker/backend
```

2. **Run Air**

```bash
air
```

Air will auto track changes to the golang code and restart the code as specified in .air.toml

---

## Troubleshooting

### Common Issues

1. **Port Already in Use**:

   - Error: `listen tcp :3000: bind: address already in use`
   - Solution: Use a different port by setting the `PORT` environment variable:
     ```bash
     export PORT=4000
     go run main.go
     ```

2. **Invalid JSON Payload**:

   - Ensure the payload in `POST` or `PUT` requests is well-formed and includes all required fields.

3. **CSS Not Being Updated**

- Ensure that your browser cache isn't caching an old version of your output.css file

---

## Deployment

We use docker and docker compose to run deployment.

The following forces the docker to build the Dockerfile and run in detached mode

```bash
docker compose up --build -d
```

It will setup the following

- The trip_tracker service with replicas
- The traefik reverse proxy

### Setting up your own hosted server

For me I used Hostinger. They have good pricing and super easy to setup with some knowledge of LinuxOS (I personally used Ubuntu)

I also used this [Youtube Video](https://www.youtube.com/watch?v=F-9KWQByeU0&ab_channel=DreamsofCode). Basically tells you how to setup a VPS from scratch. Highly recommend.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---
