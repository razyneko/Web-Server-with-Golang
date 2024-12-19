# Simple Web Server in Go

This repository contains a simple web server written in Go. It demonstrates basic functionalities like serving static files, handling forms, and routing HTTP requests to specific handlers. Below are the details of its implementation and usage.

## Features
- Serves static files from the `/static` directory.
- Handles a form submission at the `/form` endpoint.
- Responds with a greeting at the `/hello` endpoint.
- Error handling for unsupported paths and methods.

## Project Structure
```
.
├── main.go         # Go source file for the web server
├── static          # Directory containing static files
│   ├── index.html  # HTML file for the root route
│   └── form.html   # HTML file for the form
└── README.md       # Project documentation
```

## Endpoints
### 1. **Root Route (`/`)**
   - **Description:** Serves static files from the `/static` directory.
   - **Behavior:** 
     - Displays the content of `index.html` if accessed directly.
   - **Screenshot:** 
     ![Screenshot 2024-12-20 022539](https://github.com/user-attachments/assets/ef2238cb-60c7-4373-993d-41abff0f6698)


### 2. **Hello Route (`/hello`)**
   - **Description:** Responds with a simple `hello !` message.
   - **Behavior:**
     - Returns a `404` error if accessed with a method other than `GET`.
     - Returns a `404` error if the path is incorrect.
   - **Screenshot:** 
     ![Screenshot 2024-12-20 022603](https://github.com/user-attachments/assets/8e6795af-a847-458e-814d-aa5749af1846)


### 3. **Form Route (`/form`)**
   - **Description:** Processes form submissions using the `POST` method.
   - **Behavior:**
     - Parses the `name` and `address` fields from the form and displays them in the response.
     - Returns a `ParseForm()` error if the form data cannot be parsed.
   - **Screenshot:**
   - Before Submit
     ![Screenshot 2024-12-20 022646](https://github.com/user-attachments/assets/0b31172d-e282-4ccd-a34d-8e8bce898c01)
   - After Submit
     ![Screenshot 2024-12-20 023013](https://github.com/user-attachments/assets/3def6326-bff0-466e-ae46-fd28dfbdcd4f)


## Usage Instructions

### Prerequisites
- Install [Go](https://golang.org/doc/install).
- Set up your Go workspace.

### Steps to Run
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```
2. Run the server:
   ```bash
   go run main.go
   ```
3. Open your browser and navigate to:
   - **Root Route:** `http://localhost:8080`
   - **Hello Route:** `http://localhost:8080/hello`
   - **Form Route:** Submit the form at `http://localhost:8080/form`.

## Notes
- Ensure no other process is using port `8080` before starting the server.
- You can modify the port in the `main.go` file by changing:
  ```go
  http.ListenAndServe(":8080", nil)
  ```

