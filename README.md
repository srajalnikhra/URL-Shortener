
# URL Shortener – Fast & Scalable | GoLang (Go 1.24+)

![Go](https://img.shields.io/badge/Go-1.24+-blue)
![API](https://img.shields.io/badge/Type-REST%20API-green)
![License](https://img.shields.io/badge/License-MIT-yellow)

🚀 **URL Shortener** is a lightweight and high-performance web service built using **GoLang**.  
It converts long URLs into short, shareable links and provides seamless redirection.  
The project demonstrates clean routing, structured data handling, and RESTful API design using Go’s powerful **`net/http`** package.

---

## 📌 Features  
- ✅ Generate short, unique, and shareable URLs 🔗  
- ✅ Fast redirection to original URLs ⚡  
- ✅ Simple RESTful API using **Go’s `net/http`** package 🛠️  
- ✅ In-memory storage using **Go maps** 🗄️  
- ✅ Lightweight and beginner-friendly project structure 🏗️  
- ✅ Written in **Go 1.24+** with clean and readable code ✍️  

---

## ⚙️ Tech Stack  
- **Backend:** GoLang (`net/http`, `crypto/md5`, `encoding/json`)  
- **Storage:** In-memory database using Go maps  
- **Architecture:** RESTful API with clean separation of routes and handlers  

---

## 📂 Project Structure  

```plaintext
URL-Shortener/
│── main.go             # Entry point - server setup, handlers, and routing
│── go.mod              # Go module configuration
│── go.sum              # Dependencies checksum (auto-generated)
│── README.md           # Project documentation
```

---

## 🚀 Getting Started  

### Prerequisites  
- **Go** (v1.24+)

### Clone the Repository  
```bash
git clone https://github.com/srajalnikhra/URL-Shortener.git
```

```bash
cd URL-Shortener
```

### Install Dependencies  
```bash
go mod tidy
```

### Run the Project  
```bash
go run main.go
```

---

## 📌 API Endpoints  

### **1. Generate Short URL**  
**POST** `/shorten`  
Request Body:
```json
{
  "url": "https://github.com/Prince-1501"
}
```
Response:
```json
{
  "short_url": "a1b2c3d4"
}
```

---

### **2. Redirect to Original URL**  
**GET** `/redirect/{id}`  

Example:
```bash
http://localhost:3000/redirect/a1b2c3d4
```
Redirects to:  
```
https://github.com/Prince-1501
```

---

## 💡 Future Enhancements  
- 🔹 Add **persistent storage** using SQLite/PostgreSQL 🗄️  
- 🔹 Add **user authentication** with JWT 🔐  
- 🔹 Implement **rate limiting** and analytics 📊  
- 🔹 Add Docker support 🐳  

---

## 📜 License  

This project is licensed under the **MIT License**.
