# Learning Go - HNews (HackerNews-style Blogging Platform)

## 📋 Loyiha Haqida (About the Project)

Bu loyiha **Go (Golang)** tilida yozilgan **HackerNews uslubidagi blogging platformasi** hisoblanadi. Foydalanuvchilar ro'yxatdan o'tishi mumkin, login qilishi mumkin, maqolalar yuborishlari mumkin va boshqa foydalanuvchilarning maqolalarini ko'rishlari mumkin.

This is a **HackerNews-style blogging platform** built with **Go (Golang)**. Users can register, log in, submit articles, and view other users' stories.

---

## 🏗️ Arxitektura Tarkibi (Architecture Structure)

Loyiha quyidagi file va papkalardan tashkil topgan:

```
learning-go-hnews/
├── main.go              # Asosiy entry point
├── server.go            # HTTP serverni sozlash
├── routes.go            # URL marshrutlari
├── handler.go           # HTTP handler funksiyalari
├── models.go            # Ma'lumot modellari (User, Profile)
├── user_repository.go   # Database operatsiyalari
├── middlewares.go       # Middleware'lar (logging, recovery, auth)
├── render.go            # Template rendering funksiyasi
├── renderer.go          # Template renderer klasi
├── helpers.go           # Yordamchi funksiyalar
├── go.mod               # Go moduli faylı
├── go.sum               # Go dependensiyalari
├── templates/           # HTML shablonlar
│   ├── layouts/base.html    # Asosiy layout
│   ├── partials/header.html # Header component
│   ├── partials/footer.html # Footer component
│   ├── index.html       # Home page
│   ├── login.html       # Login page
│   ├── register.html    # Register page
│   ├── about.html       # About page
│   ├── contact.html     # Contact page
│   └── submit.html      # Submit article page
├── public/              # Statik fayllar
│   └── css/style.css    # CSS stillar
└── users_database.db    # SQLite database
```

---

## 🔄 Jarayon Qanday Ketayotgani (Process Flow)

### 1️⃣ **APPLICATION STARTUP (Ilovani Ishga Tushirish)**

```go
main.go → connectToDatabase() → NewSQLUserRepository() 
       → NewTemplatesRenderer() → app.serve()
```

**Qadamlar:**
1. **SQLite Database Ulanish**: `users_database.db` faylidan database yaratiladi
2. **Session Yaratish**: 12 soatlik session lifetime bilan sessions object yaratiladi
3. **Application Struct Yaratish**: Barcha kerakli komponentlar o'z joyiga qo'yiladi
4. **Server Ishga Tushirish**: `:8080` portda HTTP server startiladi

---

### 2️⃣ **REQUEST PROCESSING (So'rovni Qayta Ishlash)**

```
CLIENT REQUEST (Brauzer)
        ↓
routes() - Router
        ↓
MIDDLEWARE STACK:
  1. recover() - Xatolarni tutish
  2. logger()  - Logging
  3. session.Enable - Session
  4. requireAuth (faqat ba'zi routelarda)
        ↓
HANDLER FUNCTION (handler.go)
        ↓
TEMPLATE RENDERING (renderer.go)
        ↓
HTML RESPONSE (Brauzerga qaytarish)
```

---

## 🔐 Database Struktura (Database Structure)

### **Users Jadval**
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
```

### **Profile Jadval**
```sql
CREATE TABLE profile (
    user_id INTEGER PRIMARY KEY,
    avatar TEXT,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
)
```

---

## 🔑 Komponentlarni Tafsiliy Tushuntirish (Component Breakdown)

---

### **1. MAIN.GO - Asosiy Entry Point**

```go
func main() {
    // Database ulanish
    db, err := connectToDatabase("users_database.db")
    
    // Session yaratish (12 soatlik lifetime)
    session := sessions.New(secret)
    session.Lifetime = 12 * time.Hour
    
    // Application struct
    app := &application{
        errorLog, infoLog,   // Logging
        userRepo,            // Database operatsiyalari
        templateDir,         // Template papkasi
        publicPath,          // Static file papkasi
        session,             // Session manager
    }
    
    // Server startlash
    app.serve()
}

// Database ulanish
func connectToDatabase(name string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", name)    // Driver registratsiya
    err = db.Ping()                         // Actual connection
    return db, nil
}
```

**Nima qiladi:**
- ✅ SQLite database bilan ulanishni vosita qiladi
- ✅ Session'ni 12 soat uchun sozlaydi
- ✅ Barcha logging va repository'ni initialized qiladi
- ✅ Server'ni `:8080` portda ishga tushiradi

---

### **2. SERVER.GO - HTTP Server Konfiguratsiyasi**

```go
func (app *application) serve() error {
    srv := &http.Server{
        Addr:         ":8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
        Handler:      app.routes(),
    }
    return srv.ListenAndServe()
}
```

**Nima qiladi:**
- ✅ HTTP serverni `:8080` portda sozlaydi
- ✅ Request/Response timeout'ni 5 sekundga o'rnatadi
- ✅ `app.routes()` handler'ni ishga tushiradi

---

### **3. ROUTES.GO - URL Marshrutlari**

```go
func (app *application) routes() http.Handler {
    mux := http.NewServeMux()
    
    // Middleware'lar
    defaultMiddleware := alice.New(app.recover, app.logger)
    secureMiddleware := alice.New(app.session.Enable)
    
    // Static files
    mux.Handle("/public/", 
        http.StripPrefix("/public/", 
            http.FileServer(http.Dir(app.publicPath))))
    
    // Routes
    mux.Handle("/", 
        secureMiddleware.ThenFunc(app.home))
    mux.Handle("/login", 
        secureMiddleware.ThenFunc(app.login))
    mux.Handle("/submit", 
        secureMiddleware.Append(app.requireAuth).ThenFunc(app.submit))
    mux.Handle("/register", 
        secureMiddleware.ThenFunc(app.register))
    
    return defaultMiddleware.Then(mux)
}
```

**Route Jadvali:**

| URL | Handler | Middleware | Maqsadi |
|-----|---------|-----------|---------|
| `/` | `home()` | session.Enable | Bosh sahifa |
| `/login` | `login()` | session.Enable | Kirish sahifasi |
| `/register` | `register()` | session.Enable | Ro'yxatdan o'tish |
| `/submit` | `submit()` | session + requireAuth | Maqola yuborish |
| `/about` | `about()` | - | Haqida sahifasi |
| `/contact` | `contact()` | - | Aloqa sahifasi |
| `/public/*` | FileServer | - | CSS va boshqa statik fayllar |

---

### **4. MODELS.GO - Ma'lumot Modellari**

```go
type User struct {
    ID             int       // Foydalanuvchi IDsi
    Name           string    // Ism
    Email          string    // Email
    HashedPassword string    // Shifrlangan parol
    CreatedAt      time.Time // Yaratilgan vaqti
    Profile        Profile   // Profil ma'lumotlari
}

type Profile struct {
    UserID  int       // User IDsi (Foreign Key)
    Avatar  string    // Avatar URL
    Created time.Time // Yaratilgan vaqti
}
```

**Nima uchun ikkita struct:**
- `User` - Asosiy autentifikatsiya ma'lumotlari
- `Profile` - Qo'shimcha foydalanuvchi ma'lumotlari (avatar)

---

### **5. USER_REPOSITORY.GO - Database Operatsiyalari**

Bu file Database'ga qo'l solish uchun ishlatiladi. **Repository Pattern** deb nomlanadi.

#### **Interface**
```go
type UserRepository interface {
    CreateUser(name, email, hashedPassword, avatar string) (int64, error)
    GetUserByEmail(email string) (*User, error)
    GetUsers() ([]User, error)
}
```

#### **A) CreateUser() - Yangi Foydalanuvchi Yaratish**

```go
func (r *SQLUserRepository) CreateUser(name, email, 
                                       hashedPassword, avatar string) (int64, error) {
    // TRANSACTION BOSHLANADI
    ctx := context.Background()
    tx, err := r.db.BeginTx(ctx, nil)
    defer tx.Rollback()  // Agar xato bo'lsa orqaga qaytaradi
    
    // USERS JADVALGA QO'SHISH
    stmt, err := tx.PrepareContext(ctx, 
        `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
    
    // PAROLNI SHIFRLASH (bcrypt)
    hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), 
                                          bcrypt.DefaultCost)
    result, err := stmt.Exec(name, email, string(hp))
    
    // USER IDI OLINADI
    userID, err := result.LastInsertId()
    
    // PROFILE JADVALGA QO'SHISH
    profileStm, err := tx.PrepareContext(ctx, 
        `INSERT INTO profile (user_id, avatar) VALUES(?, ?)`)
    _, err = profileStm.Exec(userID, avatar)
    
    // TRANSACTION COMMIT QO'YILADI
    err = tx.Commit()
    
    return userID, nil
}
```

**Qadamlar:**
1. Transaction boshlanadi (agar xato bo'lsa hammasini orqaga qaytarish uchun)
2. Parol bcrypt orqali shifrlangan bo'ladi
3. User va Profile jadvallariga ma'lumot qo'shiladi
4. Agar hammasiga xatolik bo'lmasa, tranzaksyon commit qo'yiladi

#### **B) GetUserByEmail() - Emaili Orqali Foydalanuvchini Topish**

```go
func (r *SQLUserRepository) GetUserByEmail(email string) (*User, error) {
    // User va Profile jadvallarni JOIN qila
    stmt := `SELECT u.id, u.name, u.email, u.hashed_password, 
             u.created_at, p.avatar 
             FROM users u 
             INNER JOIN profile p ON u.id = p.user_id 
             WHERE u.email = ?`
    
    row := r.db.QueryRow(stmt, email)
    var user User
    
    // Ma'lumotlarni User structga joylashtirish
    err := row.Scan(&user.ID, &user.Name, &user.Email, 
                    &user.HashedPassword, &user.CreatedAt, 
                    &user.Profile.Avatar)
    
    user.Profile.UserID = user.ID
    return &user, nil
}
```

#### **C) GetUsers() - Barcha Foydalanuvchilarni Olish**

```go
func (r *SQLUserRepository) GetUsers() ([]User, error) {
    stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
    rows, err := r.db.Query(stmt)
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email, 
                        &user.HashedPassword, &user.CreatedAt)
        users = append(users, user)
    }
    
    return users, nil
}
```

---

### **6. HANDLER.GO - HTTP Handler Funksiyalari**

Handler'lar brauzer so'rovlarini qabul qiladi va response qaytaradi.

```go
// 1. HOME PAGE
func (app *application) home(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("Home page visited")
    app.infoLog.Printf("Session data: %s", app.session.GetString(r, "UserID"))
    app.render(w, "index.html", nil)
}

// 2. LOGIN PAGE
func (app *application) login(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("Login page visited")
    app.session.Put(r, "UserID", "russel")  // Test session
    app.render(w, "login.html", nil)
}

// 3. REGISTER PAGE
func (app *application) register(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("Register page visited")
    app.render(w, "register.html", nil)
}

// 4. ABOUT PAGE
func (app *application) about(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("About page visited")
    app.render(w, "about.html", nil)
}

// 5. CONTACT PAGE
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("Contact page visited")
    app.render(w, "contact.html", nil)
}

// 6. SUBMIT ARTICLE PAGE (Protected - requires login)
func (app *application) submit(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("Submit page visited")
    app.infoLog.Printf("Session data: %s", app.session.GetString(r, "UserID"))
    app.render(w, "submit.html", nil)
}
```

**Handler'larning vazifasi:**
- ✅ Logging - har bir sahifani kim tashrif buyurganini qayd etadi
- ✅ Session - foydalanuvchi ma'lumotlarini saqlab turadi
- ✅ Template rendering - HTML sahifasini yaratib brauzerga yuboradi

---

### **7. MIDDLEWARES.GO - Middleware'lar (So'rovni Filtrlash)**

Middleware'lar request'larni handler'ga tushishidan oldin filtrlaydi.

#### **A) Logger Middleware - So'rovlarni Qayd Etish**

```go
func (app *application) logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // So'rovning ma'lumotlarini logga yoza
        app.infoLog.Printf("%s %s %s %s", 
            r.RemoteAddr,  // Mijoz IP address
            r.Proto,       // HTTP protokoli (HTTP/1.1)
            r.Method,      // GET, POST, DELETE, va hokazo
            r.URL)         // URL path
        next.ServeHTTP(w, r)  // Keyingi middleware/handler'ga o'tish
    })
}
```

**Natija:**
```
INFO: 127.0.0.1 HTTP/1.1 GET /
INFO: 127.0.0.1 HTTP/1.1 POST /login
```

#### **B) Recover Middleware - Xatolarni Tutish**

```go
func (app *application) recover(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {  // Agar panic bo'lsa
                w.Header().Set("Connection", "close")
                app.serverError(w, fmt.Errorf("%s", err))  // Error yuboradI
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

**Nima qiladi:**
- Agar code'da xatolik bo'lsa, brauzer "500 Internal Server Error" xabari oladi
- Serverning boshqa so'rovlari to'xtamaydi, faqat shu so'rovni tekislaydi

#### **C) RequireAuth Middleware - Foydalanuvchi Kirishi Kerakligini Tekshirish**

```go
func (app *application) requireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !app.isAuthenticated(r) {
            // Agar autentifikatsiya qilingan bo'lmasa
            http.Redirect(w, r, 
                fmt.Sprintf("/login?redirectTo=%s", r.URL.Path), 
                http.StatusSeeOther)
            return
        }
        
        // Cache'ni o'zgarmasligi uchun
        w.Header().Set("Cache-Control", "no-store")
        next.ServeHTTP(w, r)
    })
}

// Autentifikatsiyani tekshirish
func (app *application) isAuthenticated(r *http.Request) bool {
    isAuth, ok := r.Context().Value(contextAuthKey).(bool)
    if !ok {
        return false
    }
    return isAuth
}
```

**Nima qiladi:**
- `/submit` sahifasiga faqat login qilgan foydalanuvchilar kira oladi
- Agar login qilmagan bo'lsa, `/login` sahifasiga yo'naltiradi

---

### **8. RENDERER.GO - Template Rendering Engine**

Go'ning o'rnatilgan `html/template` paketi HTML fayllarini bilan ishlaydi. Lekin biz uni sozlashtirish uchun custom renderer yaratdik.

```go
type TemplatesRenderer struct {
    cache       map[string]*template.Template  // Template kesh
    mutex       sync.RWMutex                   // Thread safety uchun
    dev         bool                           // Development mode
    templateDir string                         // Template papkasi
}

// Template renderer yaratish
func NewTemplatesRenderer(templateDir string, isDev bool) *TemplatesRenderer {
    return &TemplatesRenderer{
        templateDir: templateDir,
        cache:       make(map[string]*template.Template),
        dev:         isDev,
    }
}

// Template faylni render qilish
func (t *TemplatesRenderer) Render(w http.ResponseWriter, templateName string, data interface{}) {
    tmpl, err := t.getTemplate(templateName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // "base.html" shablon bilan render qilish
    if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Template olish (kesh bilan)
func (t *TemplatesRenderer) getTemplate(name string) (*template.Template, error) {
    if !t.dev {  // Production mode'da
        t.mutex.Lock()
        if tmpl, ok := t.cache[name]; ok {  // Cache'dan olish
            t.mutex.Unlock()
            return tmpl, nil
        }
        t.mutex.Unlock()
    }
    
    // Template parse qilish
    tmpl, err := t.parseTemplate(name)
    
    if !t.dev {  // Production mode'da
        t.mutex.Lock()
        t.cache[name] = tmpl  // Cache'ga saqlash
        t.mutex.Unlock()
    }
    
    return tmpl, nil
}

// Template fayllarni parse qilish
func (t *TemplatesRenderer) parseTemplate(name string) (*template.Template, error) {
    // Shu template faylni olish
    templatePath := path.Join(t.templateDir, name)
    files := []string{templatePath}
    
    // layouts/ papkasidagi barcha .html fayllarni olish
    layoutPath := path.Join(t.templateDir, "layouts/*.html")
    layouts, err := filepath.Glob(layoutPath)
    if err == nil {
        files = append(files, layouts...)
    }
    
    // partials/ papkasidagi barcha .html fayllarni olish
    partialPath := path.Join(t.templateDir, "partials/*.html")
    partials, err := filepath.Glob(partialPath)
    if err == nil {
        files = append(files, partials...)
    }
    
    // Hammasi birga parse qilish
    tmpl, err := template.ParseFiles(files...)
    return tmpl, nil
}
```

**Nima qiladi:**
1. **Development mode'da**: Har safar shablonlar qayta parse qilinadi (tez o'zgarish uchun)
2. **Production mode'da**: Shablonlar birinchi marta parse qilinadi va cache'ga saqlanadi (tez ishlash uchun)
3. **Shablonlarni birlashtirish**: Main page + layout + partials hammasi birga parse qilinadi

---

### **9. RENDER.GO - Oddiy Render Wrapper**

```go
func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {
    if app.tp == nil {
        http.Error(w, "template render is not rendered", http.StatusInternalServerError)
        return
    }
    app.tp.Render(w, filename, data)
}
```

**Nima qiladi:**
- TemplatesRenderer'ni qulayroq ishlatish interfeysi

---

### **10. HELPERS.GO - Yordamchi Funksiyalar**

```go
func (app *application) serverError(w http.ResponseWriter, err error) {
    // Error va stack trace qayd etish
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    app.errorLog.Output(2, trace)
    
    // Brauzerga error xabari yuborish
    http.Error(w, http.StatusText(http.StatusInternalServerError), 
               http.StatusInternalServerError)
}
```

**Nima qiladi:**
- Server xatolarini qayd etadi va brauzerga xabari yuboradi

---

## 🎨 Template Struktura (Template Structure)

### **Base Layout (base.html)**
```html
<!DOCTYPE html>
<html>
<head>
    <title>Our Simple Blog</title>
    <link rel="stylesheet" href="/public/css/style.css">
</head>
<body>
    {{template "header.html" .}}
    
    <main class="main-content">
        {{template "content" .}}
    </main>
    
    {{template "footer.html" .}}
</body>
</html>
```

### **Template Ierarxiyasi**
```
base.html (Bosh layout)
├── header.html (Navigation)
│   ├── Logo
│   └── Navigation Links
├── [Page Template] (index.html, login.html, register.html, va hokazo)
│   └── Defines "content"
└── footer.html (Footer)
```

**Qanday ishlaydi:**
1. `/` request'i qabul qilinadi
2. `home()` handler `index.html` render qiladi
3. `index.html` faylida `{{define "content"}}` bor
4. `renderer.go` `base.html` bilan birga parse qiladi
5. `base.html` `{{template "content" .}}` qiladi va `index.html` content'i o'rniga qo'yiladi

---

## 📦 Dependencies (Dependensiyalar)

```go
github.com/mattn/go-sqlite3        // SQLite database driver
github.com/golangcollege/sessions  // Session management
github.com/justinas/alice          // Middleware chaining
golang.org/x/crypto/bcrypt         // Password hashing
```

---

## ✅ Hozirgi Vaqtda Tayyorlanganlar (What's Implemented)

| Xususiyat | Status | Izohlar |
|-----------|--------|---------|
| **Database Connection** | ✅ Tayyor | SQLite bilan ulanish |
| **User Registration** | 🔄 Qisman | Form bor, lekin POST handler yo'q |
| **User Login** | 🔄 Qisman | Form bor, lekin autentifikatsiya yo'q |
| **Session Management** | ✅ Tayyor | 12 soatlik session |
| **Password Hashing** | ✅ Tayyor | bcrypt orqali |
| **User Repository** | ✅ Tayyor | CRUD operatsiyalari |
| **Template Rendering** | ✅ Tayyor | Layout + Partials + Pages |
| **Middleware Stack** | ✅ Tayyor | Logging, Recovery, Auth |
| **Static Files** | ✅ Tayyor | CSS va boshqa fayllar |
| **Error Handling** | ✅ Tayyor | Server error middleware |
| **Navigation** | 🔄 Qisman | Header/Footer shablonlari |
| **Article Submit** | 🔄 Qisman | Sahife bor, lekin backend yo'q |
| **Article Display** | 🔄 Qisman | Template bor, lekin ma'lumot yo'q |

---

## 🔧 Kerak Qilinadigan Ishlar (TODO)

### 1. **Login Handler Implement Qilish**
```go
POST /login qabul qilish
Email va parolni tekshirish
bcrypt.CompareHashAndPassword() ishlat
Session'ni sozlash
Redirect qilish
```

### 2. **Register Handler Implement Qilish**
```go
POST /register qabul qilish
Email va parolni validate qilish
CreateUser() qo'llanish
Set session
Redirect qilish
```

### 3. **Article Model Qo'shish**
```go
type Article struct {
    ID        int
    Title     string
    URL       string
    UserID    int
    Points    int
    Comments  int
    CreatedAt time.Time
}
```

### 4. **Article Repository Qo'shish**
```go
CreateArticle()
GetArticles()
GetArticleByID()
DeleteArticle()
UpdateArticle()
```

### 5. **Article Handler'larini Implement Qilish**
```go
POST /submit - maqola yuborish
GET / - maqolalarni ko'rsatish
GET /article/:id - maqolani ko'rsatish
```

### 6. **Database Migrations**
```go
Users jadvalini yaratish
Profile jadvalini yaratish
Articles jadvalini yaratish
```

---

## 📊 Data Flow Diagram

```
┌─────────────┐
│   Browser   │
└──────┬──────┘
       │ HTTP Request
       │
       ▼
┌─────────────────┐
│   routes()      │ ◄─── URL matching
└────────┬────────┘
         │
         ▼
┌──────────────────────┐
│  MIDDLEWARE STACK    │
├──────────────────────┤
│ 1. recover()        │
│ 2. logger()         │
│ 3. session.Enable   │
│ 4. requireAuth      │
└────────┬─────────────┘
         │
         ▼
┌─────────────────┐
│     HANDLER     │ ◄─── handler.go
└────────┬────────┘
         │
         ▼
┌─────────────────────┐
│ render()            │ ◄─── render.go
│  └─ TemplatesRenderer
└────────┬────────────┘
         │
         ▼
┌───────────────────────────────────┐
│ Template System (renderer.go)     │
├───────────────────────────────────┤
│ 1. getTemplate()                 │
│ 2. parseTemplate()               │
│    └─ Combine:                   │
│       • Main page (index.html)   │
│       • Layout (base.html)       │
│       • Partials (header/footer) │
│ 3. ExecuteTemplate()             │
└────────┬────────────────────────┘
         │
         ▼
     HTML Response
         │
┌────────┴────────┐
│   Browser       │
└─────────────────┘
```

---

## 🧪 Session Flow Misoli

```go
1. Foydalanuvchi /login sahifasiga kiradi
   → handler.go: login()
   → app.session.Put(r, "UserID", "russel")
   → Session cookie brauzerga yuboriladi

2. Foydalanuvchi /submit linkiga bosadi
   → middleware: requireAuth()
   → app.isAuthenticated(r) tekshiriladi
   → Agar autentifikatsiya qilingan bo'lsa, submit()
   → Agar yo'q bo'lsa, /login?redirectTo=/submit ga yo'naltiradi

3. /submit ler mos qilingan:
   → handler.go: submit()
   → app.session.GetString(r, "UserID") orqali session data olinadi
   → submit.html render qilinadi
```

---

## 🚀 Loyihani Ishga Tushirish

```bash
# Database bilan test qilish
go test -v

# Serverni ishga tushirish
go run main.go server.go routes.go handler.go \
    models.go user_repository.go middlewares.go \
    render.go renderer.go helpers.go

# Brauzerda
http://localhost:8080
```

---

## 📝 Xulosa (Summary)

Bu loyiha aslida **Go veb-ilovasi arxitekturasi**ning to'liq misolidir:

✅ **Backend**: Go HTTP server, database, middleware, handlers  
✅ **Database**: SQLite bilan CRUD operatsiyalari  
✅ **Frontend**: HTML shablonlar va CSS stillar  
✅ **Security**: Password hashing, session management, middleware  
✅ **Architecture**: Pattern-based design (Repository, Middleware)

Loyiha hozirda foundation tulamiz va qo'shimcha xususiyatlar qo'shish uchun tayyor. Authentication, article management va comment system qo'shimcha deb hisoblash mumkin.

---

**Yaratuvchi**: Learning Go - Udemy  
**Til**: Go (Golang)  
**Framework**: `net/http` (Standart library)  
**Database**: SQLite  
**Last Updated**: 2026-05-24

