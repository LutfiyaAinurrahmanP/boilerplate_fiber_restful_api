# User Management API - Complete CRUD with Role-Based Access Control

RESTful API lengkap untuk manajemen user dengan fitur authentication, authorization, CRUD operations, pagination, search, dan filter.

## 🚀 Fitur Lengkap

### Authentication & Authorization
- ✅ Register (public, default role: user)
- ✅ Login (JWT token)
- ✅ Logout
- ✅ Role-based access control (Admin & User)

### User Management (Admin)
- ✅ Create user (admin dapat pilih role)
- ✅ List users (pagination, search, filter by role)
- ✅ Get user by ID
- ✅ Update user
- ✅ Soft delete user
- ✅ Admin dashboard

### User Self-Management
- ✅ User dashboard
- ✅ View own profile
- ✅ Update own profile

## 📡 API Endpoints

### Base URL
```
http://localhost:3000
```

---

## 🌐 Public Endpoints

### 1. Home / Health Check

**Endpoint:** `GET /`

**Access:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "User Management API is running",
  "data": {
    "version": "1.0.0",
    "status": "healthy"
  }
}
```

---

### 2. Register User (Public)

**Endpoint:** `POST /register`

**Access:** Public (no authentication required)

**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "phone": "+6281234567890",
  "password": "password123",
  "confirm_password": "password123"
}
```

**Success Response (201):**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "phone": "+6281234567890",
      "role": "user",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:00:00Z"
    }
  }
}
```

**Error Response (400) - Validation:**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "username": "username must be at least 3 characters",
    "email": "email must be a valid email address",
    "password": "password must be at least 6 characters"
  }
}
```

**Error Response (409) - Duplicate:**
```json
{
  "success": false,
  "message": "username already exists"
}
```

---

### 3. Login

**Endpoint:** `POST /login`

**Access:** Public

**Request Body:**
```json
{
  "username": "johndoe",
  "password": "password123"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "phone": "+6281234567890",
      "role": "user",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:00:00Z"
    }
  }
}
```

**Error Response (401) - Invalid Credentials:**
```json
{
  "success": false,
  "message": "Invalid username or password"
}
```

---

### 4. Logout

**Endpoint:** `POST /logout`

**Access:** Public (client-side operation)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Logout successful. Please remove token from client.",
  "data": null
}
```

---

## 👑 Admin Endpoints

**All admin endpoints require admin role authentication**

**Headers:**
```
Authorization: Bearer <admin-jwt-token>
```

---

### 1. Admin Dashboard

**Endpoint:** `GET /admin/dashboard`

**Access:** Admin only

**Success Response (200):**
```json
{
  "success": true,
  "message": "Dashboard data retrieved successfully",
  "data": {
    "total_users": 150,
    "total_admins": 5,
    "total_regular_users": 145,
    "new_users_today": 12,
    "new_users_this_week": 45,
    "new_users_this_month": 120
  }
}
```

---

### 2. List All Users with Pagination

**Endpoint:** `GET /admin/users`

**Access:** Admin only

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)
- `search` (optional): Search in username, email, phone
- `role` (optional): Filter by role (user/admin)
- `sort` (optional): Sort order (asc/desc, default: desc)
- `sort_by` (optional): Sort by field (id/username/email/created_at, default: id)

**Example Request:**
```bash
GET /admin/users?page=1&limit=10&search=john&role=user&sort=desc&sort_by=created_at
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": {
    "users": [
      {
        "id": 1,
        "username": "johndoe",
        "email": "john@example.com",
        "phone": "+6281234567890",
        "role": "user",
        "created_at": "2025-10-08T10:00:00Z",
        "updated_at": "2025-10-08T10:00:00Z"
      },
      {
        "id": 3,
        "username": "johnsmith",
        "email": "johnsmith@example.com",
        "phone": "+6281234567892",
        "role": "user",
        "created_at": "2025-10-08T10:10:00Z",
        "updated_at": "2025-10-08T10:10:00Z"
      }
    ]
  },
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total": 2,
    "total_pages": 1
  }
}
```

---

### 3. Create New User

**Endpoint:** `POST /admin/users/create`

**Access:** Admin only

**Request Body:**
```json
{
  "username": "janedoe",
  "email": "jane@example.com",
  "phone": "+6281234567891",
  "password": "password123",
  "confirm_password": "password123",
  "role": "admin"
}
```

**Success Response (201):**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "user": {
      "id": 2,
      "username": "janedoe",
      "email": "jane@example.com",
      "phone": "+6281234567891",
      "role": "admin",
      "created_at": "2025-10-08T10:05:00Z",
      "updated_at": "2025-10-08T10:05:00Z"
    }
  }
}
```

**Error Response (409) - Duplicate:**
```json
{
  "success": false,
  "message": "email already exists"
}
```

---

### 4. Get User by ID

**Endpoint:** `GET /admin/users/:id`

**Access:** Admin only

**Example Request:**
```bash
GET /admin/users/1
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "phone": "+6281234567890",
      "role": "user",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:00:00Z"
    }
  }
}
```

**Error Response (404):**
```json
{
  "success": false,
  "message": "user not found"
}
```

---

### 5. Update User

**Endpoint:** `PUT /admin/users/update/:id`

**Access:** Admin only

**Request Body (All fields optional):**
```json
{
  "username": "johndoe_updated",
  "email": "john_new@example.com",
  "phone": "+6281234567899",
  "role": "admin"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe_updated",
      "email": "john_new@example.com",
      "phone": "+6281234567899",
      "role": "admin",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:30:00Z"
    }
  }
}
```

**Error Response (409) - Duplicate:**
```json
{
  "success": false,
  "message": "username already exists"
}
```

---

### 6. Delete User (Soft Delete)

**Endpoint:** `DELETE /admin/users/:id`

**Access:** Admin only

**Example Request:**
```bash
DELETE /admin/users/2
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "User deleted successfully",
  "data": null
}
```

**Error Response (400) - Self Delete:**
```json
{
  "success": false,
  "message": "You cannot delete your own account"
}
```

**Error Response (404):**
```json
{
  "success": false,
  "message": "user not found"
}
```

---

### 7. Get Admin Profile

**Endpoint:** `GET /admin/profile`

**Access:** Admin only

**Success Response (200):**
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "user": {
      "id": 2,
      "username": "adminuser",
      "email": "admin@example.com",
      "phone": "+6281234567891",
      "role": "admin",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:00:00Z"
    }
  }
}
```

---

### 8. Update Admin Profile

**Endpoint:** `PUT /admin/profile/update`

**Access:** Admin only

**Request Body (All fields optional):**
```json
{
  "username": "adminuser_updated",
  "email": "admin_new@example.com",
  "phone": "+6281234567899"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "user": {
      "id": 2,
      "username": "adminuser_updated",
      "email": "admin_new@example.com",
      "phone": "+6281234567899",
      "role": "admin",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:30:00Z"
    }
  }
}
```

**Note:** Admin cannot change their own role through profile update.

---

## 👤 User Endpoints

**All user endpoints require user role authentication**

**Headers:**
```
Authorization: Bearer <user-jwt-token>
```

---

### 1. User Dashboard

**Endpoint:** `GET /user/dashboard`

**Access:** User only

**Success Response (200):**
```json
{
  "success": true,
  "message": "Dashboard data retrieved successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "phone": "+6281234567890",
      "role": "user",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:00:00Z"
    },
    "stats": {
      "account_age_days": 30,
      "last_login": "2025-10-11T17:07:23Z"
    }
  }
}
```

---

### 2. Get User Profile

**Endpoint:** `GET /user/profile`

**Access:** User only

**Success Response (200):**
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "phone": "+6281234567890",
      "role": "user",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:00:00Z"
    }
  }
}
```

---

### 3. Update User Profile

**Endpoint:** `PUT /user/profile/update`

**Access:** User only

**Request Body (All fields optional):**
```json
{
  "username": "johndoe_updated",
  "email": "john_new@example.com",
  "phone": "+6281234567899"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe_updated",
      "email": "john_new@example.com",
      "phone": "+6281234567899",
      "role": "user",
      "created_at": "2025-10-08T10:00:00Z",
      "updated_at": "2025-10-08T10:30:00Z"
    }
  }
}
```

**Error Response (409) - Duplicate:**
```json
{
  "success": false,
  "message": "email already exists"
}
```

**Note:** Regular users cannot change their own role.

---

## 🔐 Authentication & Authorization

### JWT Token

Setelah login, Anda akan menerima JWT token yang harus disertakan di setiap request ke protected endpoints.

**Header Format:**
```
Authorization: Bearer <your-jwt-token>
```

**Token Claims:**
```json
{
  "sub": 1,                    // User ID
  "username": "johndoe",       // Username
  "role": "user",              // User role (user/admin)
  "exp": 1728387600,           // Expiration time (24 hours)
  "iat": 1728301200            // Issued at time
}
```

### Role-Based Access Control

| Endpoint | Admin | User | Public |
|----------|-------|------|--------|
| GET / | ✅ | ✅ | ✅ |
| POST /register | ✅ | ✅ | ✅ |
| POST /login | ✅ | ✅ | ✅ |
| POST /logout | ✅ | ✅ | ✅ |
| **Admin Routes** | | | |
| GET /admin/dashboard | ✅ | ❌ | ❌ |
| GET /admin/users | ✅ | ❌ | ❌ |
| POST /admin/users/create | ✅ | ❌ | ❌ |
| GET /admin/users/:id | ✅ | ❌ | ❌ |
| PUT /admin/users/update/:id | ✅ | ❌ | ❌ |
| DELETE /admin/users/:id | ✅ | ❌ | ❌ |
| GET /admin/profile | ✅ | ❌ | ❌ |
| PUT /admin/profile/update | ✅ | ❌ | ❌ |
| **User Routes** | | | |
| GET /user/dashboard | ❌ | ✅ | ❌ |
| GET /user/profile | ❌ | ✅ | ❌ |
| PUT /user/profile/update | ❌ | ✅ | ❌ |

**Legend:**
- ✅ Allowed
- ❌ Forbidden

---

## 🛠️ Setup dan Instalasi

### 1. Prerequisites

- Go 1.21 atau lebih tinggi
- MySQL 8.0 atau lebih tinggi
- Git

### 2. Clone Repository

```bash
git clone <repository-url>
cd user-management-api
```

### 3. Install Dependencies

```bash
go mod download
```

### 4. Setup Database

Buat database MySQL:

```sql
CREATE DATABASE user_management_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. Setup Environment Variables

Copy `.env.example` ke `.env`:

```bash
cp .env.example .env
```

Edit `.env`:

```env
APP_NAME=User Management API
APP_ENV=development
APP_PORT=3000

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=user_management_db

JWT_SECRET=your-super-secret-key-change-this-in-production
JWT_EXPIRE=24h
```

### 6. Run Application

```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:3000`

---

## 🧪 Testing dengan cURL

### 1. Health Check

```bash
curl -X GET http://localhost:3000/
```

### 2. Register User (Default Role: User)

```bash
curl -X POST http://localhost:3000/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "phone": "+6281234567890",
    "password": "password123",
    "confirm_password": "password123"
  }'
```

### 3. Login

```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "password123"
  }'
```

**Save the token from response!**

### 4. Create Admin User (Need Admin Token)

First, you need to manually create an admin user in database or update existing user role to 'admin':

```sql
-- Update existing user to admin
UPDATE users SET role = 'admin' WHERE username = 'johndoe';
```

Then login as admin and use the token:

```bash
# Admin Dashboard
curl -X GET http://localhost:3000/admin/dashboard \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"

# Create New User
curl -X POST http://localhost:3000/admin/users/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -d '{
    "username": "janedoe",
    "email": "jane@example.com",
    "phone": "+6281234567891",
    "password": "password123",
    "confirm_password": "password123",
    "role": "admin"
  }'
```

### 5. List Users with Pagination and Filter

```bash
curl -X GET "http://localhost:3000/admin/users?page=1&limit=10&role=user&sort=desc" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### 6. Search Users

```bash
curl -X GET "http://localhost:3000/admin/users?search=john" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### 7. Get User by ID (Admin)

```bash
curl -X GET http://localhost:3000/admin/users/1 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### 8. Update User (Admin)

```bash
curl -X PUT http://localhost:3000/admin/users/update/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -d '{
    "username": "johndoe_updated",
    "email": "john_updated@example.com"
  }'
```

### 9. Delete User (Admin - Soft Delete)

```bash
curl -X DELETE http://localhost:3000/admin/users/2 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### 10. Admin Profile Management

```bash
# Get Admin Profile
curl -X GET http://localhost:3000/admin/profile \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"

# Update Admin Profile
curl -X PUT http://localhost:3000/admin/profile/update \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -d '{
    "username": "admin_updated",
    "email": "admin_new@example.com"
  }'
```

### 11. User Dashboard and Profile

```bash
# User Dashboard
curl -X GET http://localhost:3000/user/dashboard \
  -H "Authorization: Bearer YOUR_USER_TOKEN"

# Get User Profile
curl -X GET http://localhost:3000/user/profile \
  -H "Authorization: Bearer YOUR_USER_TOKEN"

# Update User Profile
curl -X PUT http://localhost:3000/user/profile/update \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_USER_TOKEN" \
  -d '{
    "username": "johndoe_updated",
    "email": "john_updated@example.com"
  }'
```

---

## 🗄️ Database Schema

### Users Table

```sql
CREATE TABLE users (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  phone VARCHAR(20) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(20) NOT NULL DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  
  INDEX idx_username (username),
  INDEX idx_email (email),
  INDEX idx_role (role),
  INDEX idx_deleted_at (deleted_at)
);
```

**Indexes:**
- Primary key pada `id`
- Unique index pada `username`, `email`, `phone`
- Index pada `role` untuk filter by role yang cepat
- Index pada `deleted_at` untuk soft delete queries

---

## 📊 Validation Rules

### Register & Create User:
- **username**: Required, min 3, max 50, unique
- **email**: Required, valid email format, unique
- **phone**: Required, min 10, max 20, unique
- **password**: Required, min 6
- **confirm_password**: Required, must match password
- **role**: Optional for register (default: user), Required for create user, must be 'user' or 'admin'

### Update User:
- **username**: Optional, min 3, max 50, unique
- **email**: Optional, valid email format, unique
- **phone**: Optional, min 10, max 20, unique
- **role**: Optional, must be 'user' or 'admin' (admin only)

### Login:
- **username**: Required
- **password**: Required

### List Users Query:
- **page**: Optional, min 1, default 1
- **limit**: Optional, min 1, max 100, default 10
- **search**: Optional, max 100 characters
- **role**: Optional, must be 'user' or 'admin'
- **sort**: Optional, must be 'asc' or 'desc', default 'desc'
- **sort_by**: Optional, must be 'id', 'username', 'email', or 'created_at', default 'id'

---

## 🏗️ Arsitektur Aplikasi

### Clean Architecture Layers:

```
┌─────────────────────────────────────┐
│   Presentation Layer (Handlers)     │  ← HTTP Request/Response
├─────────────────────────────────────┤
│   Business Logic Layer (Services)   │  ← Business Rules & Validation
├─────────────────────────────────────┤
│   Data Access Layer (Repositories)  │  ← Database Operations
├─────────────────────────────────────┤
│   Domain Layer (Models)             │  ← Entities & Domain Logic
└─────────────────────────────────────┘
```

### Dependency Flow:

```
main.go
  ├─> Config
  ├─> Database Connection
  ├─> Repositories (depends on DB)
  ├─> Services (depends on Repositories)
  ├─> Handlers (depends on Services)
  └─> Routes (depends on Handlers)
```

### Request Flow:

```
Client Request
  ↓
Middlewares (Recover, Logger, CORS)
  ↓
Auth Middleware (JWT Verification)
  ↓
Role Middleware (Authorization)
  ↓
Handler (Parse & Validate Request)
  ↓
Service (Business Logic)
  ↓
Repository (Database Query)
  ↓
Database (MySQL)
  ↓
Repository (Return Data)
  ↓
Service (Process Data)
  ↓
Handler (Format Response)
  ↓
Client Response
```

### Routes Structure:

```
/
├── / (GET) - Health check
├── /register (POST) - Public registration
├── /login (POST) - Authentication
├── /logout (POST) - Logout
│
├── /admin/* (Admin only routes)
│   ├── /admin/dashboard (GET)
│   ├── /admin/users (GET)
│   ├── /admin/users/create (POST)
│   ├── /admin/users/:id (GET)
│   ├── /admin/users/update/:id (PUT)
│   ├── /admin/users/:id (DELETE)
│   ├── /admin/profile (GET)
│   └── /admin/profile/update (PUT)
│
└── /user/* (User only routes)
    ├── /user/dashboard (GET)
    ├── /user/profile (GET)
    └── /user/profile/update (PUT)
```

---

## 🔒 Security Features

1. **Password Security**
   - Bcrypt hashing with cost 10
   - Never store plain text passwords
   - Password minimum 6 characters

2. **Authentication**
   - JWT token with 24 hours expiration
   - Stateless authentication
   - Token includes user ID, username, and role

3. **Authorization**
   - Role-based access control (RBAC)
   - Admin and User roles
   - Strict route separation (/admin/* vs /user/*)

4. **Input Validation**
   - Comprehensive validation rules
   - SQL injection prevention (GORM parameterized queries)
   - XSS prevention (JSON response only)

5. **Database Security**
   - Soft delete for data recovery
   - Unique constraints on username, email, phone
   - Connection pooling with limits

6. **API Security**
   - CORS configuration
   - Error handling without exposing sensitive info
   - Rate limiting ready (to be implemented)

---

## 📈 Fitur Query & Filter

### Pagination

```bash
# Page 1, 10 items per page
GET /admin/users?page=1&limit=10

# Page 2, 20 items per page
GET /admin/users?page=2&limit=20
```

### Search

```bash
# Search in username, email, and phone
GET /admin/users?search=john

# Search akan mencari di:
# - username yang mengandung "john"
# - email yang mengandung "john"
# - phone yang mengandung "john"
```

### Filter by Role

```bash
# Get only users with role "user"
GET /admin/users?role=user

# Get only users with role "admin"
GET /admin/users?role=admin
```

### Sorting

```bash
# Sort by created_at descending (newest first)
GET /admin/users?sort=desc&sort_by=created_at

# Sort by username ascending (A-Z)
GET /admin/users?sort=asc&sort_by=username

# Sort by email descending (Z-A)
GET /admin/users?sort=desc&sort_by=email
```

### Kombinasi Query

```bash
# Search "john", filter role "user", page 1, 10 items, sort by created_at desc
GET /admin/users?search=john&role=user&page=1&limit=10&sort=desc&sort_by=created_at
```

---

## 📝 Response Format

### Success Response (Without Pagination)

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    // response data
  }
}
```

### Success Response (With Pagination)

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    // array of items
  },
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total": 25,
    "total_pages": 3
  }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "errors": {
    // validation errors (optional)
  }
}
```

---

## 🎯 Use Cases

### Scenario 1: User Registration and Login
```bash
# 1. Register new user
curl -X POST http://localhost:3000/register \
  -H "Content-Type: application/json" \
  -d '{"username":"newuser","email":"new@example.com","phone":"+6281234567890","password":"password123","confirm_password":"password123"}'

# 2. Login
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{"username":"newuser","password":"password123"}'

# 3. Access user dashboard
curl -X GET http://localhost:3000/user/dashboard \
  -H "Authorization: Bearer <token>"
```

### Scenario 2: Admin Managing Users
```bash
# 1. Login as admin
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"adminpass"}'

# 2. View dashboard
curl -X GET http://localhost:3000/admin/dashboard \
  -H "Authorization: Bearer <admin-token>"

# 3. List all users
curl -X GET http://localhost:3000/admin/users \
  -H "Authorization: Bearer <admin-token>"

# 4. Create new admin
curl -X POST http://localhost:3000/admin/users/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <admin-token>" \
  -d '{"username":"newadmin","email":"newadmin@example.com","phone":"+6281234567891","password":"password123","confirm_password":"password123","role":"admin"}'

# 5. Update user
curl -X PUT http://localhost:3000/admin/users/update/5 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <admin-token>" \
  -d '{"role":"admin"}'

# 6. Delete user
curl -X DELETE http://localhost:3000/admin/users/10 \
  -H "Authorization: Bearer <admin-token>"
```

### Scenario 3: User Profile Management
```bash
# 1. View own profile
curl -X GET http://localhost:3000/user/profile \
  -H "Authorization: Bearer <user-token>"

# 2. Update own profile
curl -X PUT http://localhost:3000/user/profile/update \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <user-token>" \
  -d '{"email":"newemail@example.com","phone":"+6281234567899"}'
```

---

## 🚧 Future Enhancements

- [ ] Email verification
- [ ] Password reset via email
- [ ] Refresh token mechanism
- [ ] Rate limiting
- [ ] API documentation (Swagger/OpenAPI)
- [ ] User avatar upload
- [ ] Audit logging (track who changed what)
- [ ] Export users to CSV/Excel
- [ ] Bulk user import
- [ ] Advanced filtering (date range, multiple roles, etc)
- [ ] User activity history
- [ ] Two-factor authentication (2FA)
- [ ] OAuth integration (Google, Facebook, etc)
- [ ] WebSocket support for real-time notifications
- [ ] Admin activity logs

---

## 🐛 Troubleshooting

### Database Connection Error

```
Failed to connect to database
```

**Solution:**
- Check if MySQL is running
- Verify database credentials in `.env`
- Ensure database exists: `CREATE DATABASE user_management_db;`

### Validation Error

```
Validation failed: username must be at least 3 characters
```

**Solution:**
- Check request body format
- Ensure all required fields are provided
- Verify field values meet validation rules

### Unauthorized Error

```
Missing authorization header
```

**Solution:**
- Include JWT token in header: `Authorization: Bearer <token>`
- Ensure token is not expired (24 hours validity)
- Login again to get new token if expired

### Forbidden Error

```
Forbidden: insufficient permissions
```

**Solution:**
- Ensure you have correct role for the operation
- Admin token required for `/admin/*` endpoints
- User token required for `/user/*` endpoints
- Check if you're accessing the correct route for your role

### Route Not Found

```
Cannot GET /api/v1/users
```

**Solution:**
- This API does NOT use `/api/v1` prefix
- Use direct routes: `/admin/users`, `/user/profile`, etc.
- Check the endpoint list in this README

---

## 📚 Postman Collection

Anda dapat mengimport collection Postman untuk testing:

**Struktur Folder:**
```
User Management API/
├── Public/
│   ├── Health Check (GET /)
│   ├── Register (POST /register)
│   ├── Login (POST /login)
│   └── Logout (POST /logout)
├── Admin/
│   ├── Dashboard (GET /admin/dashboard)
│   ├── List Users (GET /admin/users)
│   ├── Create User (POST /admin/users/create)
│   ├── Get User (GET /admin/users/:id)
│   ├── Update User (PUT /admin/users/update/:id)
│   ├── Delete User (DELETE /admin/users/:id)
│   ├── Get Profile (GET /admin/profile)
│   └── Update Profile (PUT /admin/profile/update)
└── User/
    ├── Dashboard (GET /user/dashboard)
    ├── Get Profile (GET /user/profile)
    └── Update Profile (PUT /user/profile/update)
```

---

## 📄 License

MIT License

## 👨‍💻 Author

Lutfiya PR - [@lutfiyapr](https://github.com/lutfiyapr)

---

## 📞 Support

Jika Anda menemukan bug atau memiliki pertanyaan:
- Open issue di GitHub
- Email: lutfiyapr@example.com

---

**Happy Coding! 🚀**

**Made with ❤️ using Go, Fiber, GORM, and MySQL**