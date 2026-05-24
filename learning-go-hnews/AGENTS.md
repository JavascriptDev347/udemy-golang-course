# AGENTS.md - AI Coding Guide for Udemy Go Course

## Project Overview

This is a **Udemy Go learning course repository** with 11+ educational sections plus a practical HNews blogging platform application (`learning-go-hnews/`). Focus AI assistance on the main application unless explicitly asked about specific educational sections.

---

## Architecture & Key Components

### Application Structure (learning-go-hnews/)

The main application follows a **layered architecture** with clear separation of concerns:

```
main.go            → Application bootstrap + dependency injection
├── server.go      → HTTP server setup (port :8080, timeouts)
├── routes.go      → Middleware chain + route definitions
├── handler.go     → HTTP request handlers
├── models.go      → Data structs (User, Profile)
├── user_repository.go  → Database access layer (Repository pattern)
├── middlewares.go  → Middleware implementations (recover, logger, auth)
├── renderer.go    → Template rendering engine
├── render.go      → Render helper wrapper
├── helpers.go     → Utility functions (error handling)
└── templates/     → HTML templates with layout composition
```

### Critical Data Flows

**Request Processing Pipeline:**
```
1. Request → routes() router
2. defaultMiddleware: recover(), logger()
3. secureMiddleware: session.Enable()
4. Optional: requireAuth middleware (for protected routes)
5. Handler function → app.render(w, templateName, data)
6. Renderer: combines base.html + page template + partials
7. Response sent to client
```

**User Registration Flow:**
```
1. CreateUser() → transaction begins
2. Password hashed with bcrypt.GenerateFromPassword()
3. User inserted into users table
4. UserID retrieved with LastInsertId()
5. Profile record created with foreign key
6. Transaction committed (or rolled back on error)
```

---

## Project-Specific Conventions

### 1. **Repository Pattern for Database Access**

All database operations go through the `UserRepository` interface:

```go
type UserRepository interface {
    CreateUser(name, email, hashedPassword, avatar string) (int64, error)
    GetUserByEmail(email string) (*User, error)
    GetUsers() ([]User, error)
    Authenticate(email, password string) (int, error)
}
```

**Key points:**
- Interface defined in `user_repository.go`
- `SQLUserRepository` is the concrete implementation
- Always use transactions for multi-table operations
- Use `context.Background()` for context in transactions
- Always `defer tx.Rollback()` before committing

### 2. **Middleware Chaining with justinas/alice**

Routes are wrapped with middleware using the alice library:

```go
defaultMiddleware := alice.New(app.recover, app.logger)
secureMiddleware := alice.New(app.session.Enable)
mux.Handle("/", secureMiddleware.ThenFunc(app.home))
mux.Handle("/admin", secureMiddleware.Append(app.requireAuth).ThenFunc(app.admin))
```

**Convention:**
- `recover()` and `logger()` are always in `defaultMiddleware` (outermost)
- `session.Enable` wraps routes that need session data
- `requireAuth` is appended only to protected routes
- Order matters: recovery → logging → session → app logic

### 3. **Application Struct Dependency Injection**

Core application struct holds all dependencies:

```go
type application struct {
    errorLog    *log.Logger
    infoLog     *log.Logger
    userRepo    UserRepository      // Interface, not concrete
    templateDir string
    publicPath  string
    tp          *TemplatesRenderer
    session     *sessions.Session
}
```

**When modifying:** Add new dependencies here, not as global variables. Inject via constructor-like pattern in `main()`.

### 4. **Template Rendering with Composition**

Templates use Go's `html/template` with a composition pattern:

```
base.html (master layout)
├── {{template "header.html" .}}
├── {{template "content" .}}     ← Placeholder for page-specific content
└── {{template "footer.html" .}}
```

Each page (e.g., `index.html`) defines what `"content"` template yields:
```html
{{define "content"}}
    <!-- page-specific markup -->
{{end}}
```

**Renderer behavior:**
- **Dev mode** (`isDev: true`): Templates re-parsed on every request (but now always in dev)
- **Production mode** (`isDev: false`): Templates cached in memory with mutex protection
- Automatically parses: main template + `layouts/*.html` + `partials/*.html`

### 5. **Error Handling Patterns**

Custom error variables and custom error types:

```go
var ErrInvalidCredentials = errors.New("invalid credentials")

// In handlers:
func (app *application) serverError(w http.ResponseWriter, err error) {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    app.errorLog.Output(2, trace)
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
```

**Pattern:** Errors are logged with stack trace, generic message sent to client.

---

## External Dependencies & Integration Points

### Core Dependencies (in go.mod)

- **github.com/mattn/go-sqlite3** v1.14.44 → SQLite driver
- **golang.org/x/crypto** v0.51.0 → bcrypt for password hashing
- **github.com/golangcollege/sessions** v1.2.0 → Session management (12-hour lifetime)
- **github.com/justinas/alice** v1.2.0 → Middleware chaining

### Database Connection

```go
// Always use this pattern:
db, err := sql.Open("sqlite3", "users_database.db")  // Registers driver, doesn't connect
err = db.Ping()  // Actually verifies connection
```

Database file: `learning-go-hnews/users_database.db` (SQLite)

### Session Configuration

```go
session := sessions.New(secret)
session.Lifetime = 12 * time.Hour
session.Secure = true
session.SameSite = http.SameSiteLaxMode
```

Session key for user ID: `"UserID"` (used in `requireAuth` middleware)

---

## Build, Run & Development Commands

### Running the Main Application

```powershell
# From project root
cd learning-go-hnews
go run main.go
# Server listens on http://localhost:8080
```

### Running Educational Sections

Each section folder has standalone examples:

```powershell
cd section_2/1-foor-loop
go run main.go
```

### Database & Dependencies

```powershell
# Download dependencies (from learning-go-hnews or any section)
go mod download

# Verify no security issues
go list -json -m all | nancy sleuth  # if nancy installed
```

---

## File-Specific Modification Patterns

### Adding a New Route

Edit `routes.go`:
1. Define handler in `handler.go`
2. Add middleware wrapper in `routes.go` using alice chain
3. Create template in `templates/` directory
4. Reference in handler with `app.render(w, "page.html", data)`

### Adding Database Operations

Edit `user_repository.go`:
1. Add method to `UserRepository` interface
2. Implement in `SQLUserRepository` struct
3. Use patterns: transactions for writes, queries for reads
4. Always handle `context.Background()` for context

### Templates

- Base layout: `templates/layouts/base.html`
- Page templates: `templates/*.html` (each defines `{{define "content"}}...{{end}}`)
- Reusable components: `templates/partials/header.html`, `footer.html`

---

## Common Scenarios & Solutions

**Scenario: Add authentication to a route**
- Set `session.Put(r, "UserID", userID)` in login handler
- Append `app.requireAuth` middleware in `routes.go`
- Check `app.isAuthenticated(r)` in handlers (middleware does this)

**Scenario: Add a new user field**
1. Update `User` struct in `models.go` with JSON tags
2. Update database schema (add column to users table)
3. Update repository methods' SQL queries
4. Update template renders that display user data

**Scenario: Database error in production**
- `app.serverError(w, err)` logs with stack trace, sends generic 500 to client
- Check `app.infoLog` / `app.errorLog` (configured in `main.go` to stderr/stdout)

---

## Testing & Debugging Notes

- **No tests in current codebase** - testing infrastructure not set up
- **Logging:** Two loggers configured (`errorLog` to stderr, `infoLog` to stdout)
- **Common debug points:**
  - Middleware chain order (recover → logger → session → auth)
  - Template parsing in dev mode vs production
  - Transaction rollback in repository (always `defer tx.Rollback()`)
  - Session cookie presence on protected routes

---

## Educational Sections (Reference Only)

The `section_X/` folders contain:
- **section_1**: Go basics
- **section_2**: Control flow & loops
- **section_3**: Data structures (arrays, slices, maps, pointers)
- **section_4**: Functions, errors, defer, panic/recovery
- **section_5**: Structs, methods, interfaces, generics
- **section_6**: Composition & embedding
- **section_7**: Strings, formatting, unicode, regex
- **section_8**: Goroutines, channels, concurrency
- **section_9**: File I/O
- **section_10**: JSON marshaling/unmarshaling
- **section_11**: Database operations (progressively from basic to repository pattern)

Unless task specifies a section, focus on `learning-go-hnews/` only.

---

## Key Files Reference Map

| When You Need To... | Edit This File |
|---|---|
| Add a new HTTP endpoint | `routes.go` + `handler.go` + template |
| Add a database operation | `user_repository.go` + `models.go` if needed |
| Modify authentication logic | `middlewares.go` (requireAuth) + `user_repository.go` (Authenticate) |
| Add application-wide settings | `main.go` (application struct) |
| Change template layout | `templates/layouts/base.html` |
| Add reusable template component | `templates/partials/*.html` |
| Fix rendering issues | `renderer.go` (template compilation) |
| Handle new errors | `helpers.go` (serverError) or add custom error in relevant file |

