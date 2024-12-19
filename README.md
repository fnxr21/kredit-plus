# kredit+

Credit System for Conventional Store or Dealer:

I have set up user levels such as admin and customer to streamline the process. The goal is to ensure that the company can directly connect with customers, without needing to engage with our partners too much. This allows us to maintain strong customer connections while still keeping relevant partner information and connections intact. The focus is on ensuring a direct relationship with customers first and foremost.

## Tech Stack

**Server:** Golang,echo, Gorm,validator/v10 ,jwt/v4, disintegration ,**RSA**

## Link

- [ERD](https://drawsql.app/teams/team-1756/diagrams/kredit)
- [Postman Api](https://crimson-crater-385688.postman.co/workspace/Open~b2b49bd4-fdb3-472f-ae87-d815a933728a/collection/27541101-9d6c7a2b-bb01-4afd-a537-a78a8d7b6110?action=share&creator=27541101&active-environment=27541101-2db38100-c936-4952-887b-cc478f8892ec)
- [Docker Hub](https://hub.docker.com/repository/docker/nxrfandi/kredit-plus/general)

## folder structur

```plaintext
├── kredit/                   # Contains business logic or modules specific to the "kredit" feature
├── cmd/                      # Contains application entry points or command-line commands
│   └── server.go             # Main server setup and initialization
├── doc/                      # Documentation files (e.g., API specs, design notes)
├── internal/                 # Private application code, organized by domain
│   ├── database/             # Database setup and access (e.g., connection, migrations)
│   ├── dto/                  # Data Transfer Objects for structuring input and output data
│   ├── models/               # Data models or ORM entities for database operations
│   ├── repository/           # Interfaces and implementations for data persistence
│   ├── router/               # HTTP router setup, including API route definitions
│   ├── service/              # Business logic and core application services
├── pkg/                      # Utility and reusable libraries shared across the application
├── public                    # Imagesave
├── .env                      # Environment variables configuration file
├── docker-compose.yaml       # Docker Compose file for container orchestration
├── Dockerfile                # Dockerfile for building the application container
├── main.go                   # Application entry point
├── Makefile                  # Automation tasks (e.g., build, test, deploy)
└── README.md                 # Project overview and usage guide
```

## Task

### Persyaratan Minimum

1. **Adopsi Git Flow di git**
   I manage two Git branches (main and dev) using Git flow, ensuring proper version control and development cycles.

2. **Adopsi Clean Code Arhitecture**
   This project follows a **modular architecture** focused on **separation of concerns**, **scalability**, and **maintainability**. The folder structure is designed to keep business logic, services, and data access components organized and easy to extend, while ensuring the application can scale and be maintained over time

3. **Handling concurrent transaction di salah satu endpoint yang mnurut anda perlu handling concurrent transaction**
   I ensure proper handling of concurrent transactions, particularly in the context of payment transactions. During payment creation

4. **Adopsi minimal 3 pencegahan serangan keamanan yang masuk dalam TOP 10 OWASP**
   This document outlines the security measures implemented in the project, addressing key OWASP Top 10 risks and outlining the techniques used to ensure the system's integrity, confidentiality, and availability.

Session Management
**JWT Token Usage**

- **OWASP Top 10 Impact**: [A2: Broken Authentication](https://owasp.org/www-project-top-ten/#a2-broken-authentication)
- I utilize **JWT (JSON Web Token)** for session management, ensuring secure and efficient user authentication. Token-based authentication allows for stateless sessions, improving scalability and performance.
- **Logout Functionality**: The token is invalidated upon user logout, ensuring that sessions are properly terminated.

Output Encoding
**SQL Injection Prevention**

- **OWASP Top 10 Impact**: [A1: Injection](https://owasp.org/www-project-top-ten/#a1-injection)
- I rely on **GORM's built-in protection mechanisms** for preventing SQL injection. Additionally, parameterized queries and input sanitization are employed to mitigate SQL injection risks automatically.
- **XSS (Cross-Site Scripting) Prevention**
  - **OWASP Top 10 Impact**: [A7: Cross-Site Scripting (XSS)](https://owasp.org/www-project-top-ten/#a7-cross-site-scripting)
  - I avoid using plain text and ensure that all user inputs are sanitized. This prevents the execution of malicious code on user browsers.

Input Validation
**DTO Validation**

- **OWASP Top 10 Impact**: [A1: Injection](https://owasp.org/www-project-top-ten/#a1-injection)
- Input validation is handled using the **validation/v10** package. Validation occurs at the Data Transfer Object (DTO) level to ensure that only valid data enters the system.
- **String and Encoding**: I use **strconv** for string manipulation and ensure proper encoding/decoding, especially for image processing (e.g., base64 encoding/decoding).

File Management
**File Uploads**

- **OWASP Top 10 Impact**: [A5: Security Misconfiguration](https://owasp.org/www-project-top-ten/#a5-security-misconfiguration)
- For file uploads, I restrict file formats to only **JPEG** and **PNG** images. I also verify MIME types to ensure that only valid image files are uploaded, preventing the upload of malicious files.

Error Handling
**Logging and Middleware**

- **OWASP Top 10 Impact**: [A6: Security Misconfiguration](https://owasp.org/www-project-top-ten/#a6-security-misconfiguration)
- Custom middleware functions handle error logging, ensuring that sensitive error details are not exposed to users. Proper error responses are returned, and authentication errors are handled in the middleware.

Stored Procedures
**Using Stored Procedures**

- **OWASP Top 10 Impact**: [A1: Injection](https://owasp.org/www-project-top-ten/#a1-injection)
- I use **stored procedures** for optimized and secure database operations. This approach minimizes the risk of SQL injection by ensuring that queries are executed in a controlled and parameterized manner.

Cryptography
**RSA Implementation**

- **OWASP Top 10 Impact**: [A3: Sensitive Data Exposure](https://owasp.org/www-project-top-ten/#a3-sensitive-data-exposure)
- I have integrated an **RSA cryptographic package** to securely encrypt and decrypt sensitive data. This is particularly useful for handling sensitive data (e.g., QRIS data) that needs to be securely transferred across systems.

Access Tokens
**Token Management**

- **OWASP Top 10 Impact**: [A2: Broken Authentication](https://owasp.org/www-project-top-ten/#a2-broken-authentication)
- Access token management is handled through a dedicated package in the project. It includes token issuance, validation, and expiration management, ensuring secure access control for users.

---

## **OWASP Top 10 Risks Addressed**

- **A1: Injection**: Mitigated through GORM’s SQL injection protection, input validation, and stored procedures.
- **A2: Broken Authentication**: Handled with JWT token management, including secure logout functionality.
- **A3: Sensitive Data Exposure**: Prevented by using RSA encryption for sensitive data.
- **A5: Security Misconfiguration**: Prevented through secure file upload management and error handling.
- **A6: Security Misconfiguration**: Mitigated by logging errors securely and handling authentication failures in middleware.
- **A7: Cross-Site Scripting (XSS)**: Mitigated by sanitizing user inputs and avoiding plain text.

---

5. **Adopsi Unit Test**
   Unit Testing: Currently, unit testing is not yet implemented.

### Nilai Tambah

1. [Dockerize aplikasi yang dibangun](https://hub.docker.com/repository/docker/nxrfandi/kredit-plus/general)

### Hasil Pengerjaan

1. [Github Repository url](https://github.com/fnxr21/kredit-plus)
2. [File SQL](https://github.com/fnxr21/kredit-plus/tree/main/doc/sql)
3. Gambar Arsitektur Aplikasi yang diharapkan based on studi kasus diatas
4. [Gambar Desaign Database dalam bentuk Entity Relationship Diagram](https://drawsql.app/teams/team-1756/diagrams/kredit)

## .env

```.env
# JWT Secret Key - Used for signing and verifying JWT tokens
SECRET_KEY=SECRET_KEY # Replace with your actual secret key

# Application Configuration
APP_PORT=2001 # Port on which the service will be exposed

# Database Configuration
# Host settings for local or Dockerized environments
DB_HOST=host.docker.internal # Use this for Docker environments on Windows and macOS.
DB_HOST1=127.0.0.1 # Use this for local environment (localhost), adjust for different configurations.
DB_PORT=1306  # Database port
DB_USER=roo21! # Database username
DB_PASS=root21!Save # Database user password
DB_ROOTPASSWORD=root21!SaveMain # Superuser password for the database

# Database Charset and Timezone Configuration
DB_CHARNTIME=?charset=utf8mb4&parseTime=True&loc=Local # Charset for UTF-8 and time parsing with local timezone

# Database Relations (Schema and Ports)
DB_MST=kredit_mst # The main database schema used for operations
DB_PORT_EXPOSE=1306 # Exposed port for the database connection
DB_DEFAULTPORT=3306 # Default MySQL port, don't change this unless necessary
```

## Docker Compose

Clone the project

```bash
  git clone https://github.com/fnxr21/kredit-plus
```

Go to the project directory

```bash
  cd kredit-plus
```

To use Docker-compose.yaml run

```bash
  make dc
```

after build and success than "cntrl + c" ,than run

```bash
  make dc-up
```

## Run Locally

Start the server

```bash
  make run
```

## Authors

- [Fandi](https://www.github.com/fnxr21)
- [Github](https://www.github.com/fnxr21)
- [Linkind](https://www.linkedin.com/in/fandi-nur/)
