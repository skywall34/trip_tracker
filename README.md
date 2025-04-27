# Trip Tracker Project (Mia's Trips)

GO Version: 1.23.5

This project started as a side project to build a simple travel website for one particular user. What started as a simple CRUD app is now expanding to become a fantastic learning experience and passion project to understand deploying production ready application in Golang/HTMX.

This application is written in Go/HTMX/Templ and runs an application for managing trips. The backend uses the built-in `net/http` package for handling HTTP requests and maintains the trips in a `sqlite` database. The application also uses a real time API for flight/airport information in the future and allowing the user to create/edit/delete trips.

## Resources

We use the aviationstack API for real time flight/airport information [Check their site here](https://aviationstack.com)

[API Docs](https://aviationstack.com/documentation)

---

## Architecture

### Middleware

This application uses the following middleware

- Auth
- TextHTML Middleware (serving html from the backend)
- CSP Middleware
- Logging

The CSP Middleware is used as a security measure to prevent unexpected `<script>` tags, inline JS code, and external resources such as images, fonts, etc. Since HTMX dynamically swaps html into the page via AJAX this is especially important in production environment to prevent XSS (cross-site scripting) and similar attacks.

**Dev Note** I have tried using chaining middleware. For some reason though it seems to break CSP And TextHTML Middleware effectively making the app inoperable. It is something I wish to tackle in the future.

### Database

The sqlite database currently uses the following tables (all which can be found under `internal/database/schema.sql`)

- airports: Static list of all airports. Populated using csv files received from public websites
- users: Holds user information
- trips: Holder all trip information. Many queries will pair this with airports via a `JOIN` operation
- sessions: Holds session data of the user.
- password_reset_tokens: Holds 1 hour expiry reset tokens for users requesting forgot-password

All .go files under the database serve to run SQL queries on their respective tables and pass them to the handlers.

### Handlers

THe files in this folder represent the backend of the project. The naming convention for these files generally fall under this ruleset:

- `{get/post/update/delete}{resource_name}.go`

For example, to create a handler which is called by a GET request to retrieve the home page, we name the handler `gethome.go`

Each handler also follows the following file structure so that it can be easily called in the mux handler:

```go
// Base Handler Struct. Any database structs are defined here
type DeleteTripHandler struct {
	tripStore *db.TripStore
}

// The Handler Params struct
type DeleteTripHandlerParams struct {
	TripStore *db.TripStore
}

// This function allows the mux in main.go to pass in the actual database services
func NewDeleteTripHandler(params DeleteTripHandlerParams) (*DeleteTripHandler) {
	return &DeleteTripHandler{
		tripStore: params.TripStore,
	}

func (h *NewDeleteTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // Handler Logic Here
}
```

### Templates

The `templates` folder stores all the htmx/templ files. The general convention here is that if a component exist or will populate the same page they must live in the same file.

For example, all trip related components (RenderPastTrips, CreateTripPage) all live under `trips.templ`. This is to organize and easily find where all the components are.

In addition, for the page to appear under tha layout the layout must be instantiated in the handler servicing the page and then rendered with the component to go inside. Example:

`internal/handlers/gettrip.go`

```go
c := templates.TripsPage()
templates.Layout(c, "Trips").Render(r.Context(), w)
```

### Static JS

We have a few JS files in the application. Apart from the htmx.min.js file required to run the project there are few which fulfill a specific task in the application

convertTimes.js: Standardizes inputs into UTC for uniform time metrics (Handling time zones)
leaflet.js: Downloaded from Leaflet, this is used to render the world map (statistics page)
map.js: Uses the leaflet.js and runs configuration such as size and markers
modal.js: JS logic to show/hide the hidden trip-form element on the trips page
response-targets.js: HTMX logic to allow swapping of elements in different targets
tabs.js: CSS logic to mimic sliding animations

### External APIs

- Google Auth:
  - google_auth.go: Sets up the auth config
  - google_login.go: Sends a request to Google oauth
  - google_callback.go: Once Google processes the request, the callback validates the user and sets the session for login
- Flights
  - AviationStack has a free 100 calls/month plan which allows to call real time flight data. Here, we get the flight route using the iata_code and request the user to enter in the date.

## Prerequisites

1. **Go Installed**: Ensure Go is installed on your system (at least 1.23). You can download it from [here](https://golang.org/dl/).

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
     Another solution is to kill the process using port 3000

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
- Watchtower to check for new updates to docker images and then run rolling updates

Once you have it initially setup, watchtower will listen for any new updates to the docker image and update if there's a change to a tag

There is an example script `docker_build.sh` which will build the Dockerfile and then upload to this repository as a ghcr.io image.

### Setting up your own hosted server

For me I used Hostinger. They have good pricing and super easy to setup with some knowledge of LinuxOS (I personally used Ubuntu)

I also used this [Youtube Video](https://www.youtube.com/watch?v=F-9KWQByeU0&ab_channel=DreamsofCode). Basically tells you how to setup a VPS from scratch. Highly recommend.

## License

This project is licensed under the APACHE 2.0 License. See the `LICENSE` file for details.

---
