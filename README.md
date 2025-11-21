This project is a complete implementation of the Shopping Cart assignment. It includes a fully working backend built with Go (Gin + GORM) and a frontend built with React (Vite). The application follows the exact workflow defined in the assessment:

User ->Cart ->Order
It covers user authentication, creating and listing items, creating carts, adding items to carts, converting carts to orders, and viewing order history.
SECTION 1: FEATURES
Backend Features:
User registration
User login with token-based authentication
Single active token per user
Create items
List items
Create cart automatically for a user
Add items to cart
Convert cart into an order
List orders for the user
SQLite database with auto-migration

Frontend Features:
Login screen
Items listing page
Click on item to add to cart
Checkout button to convert cart to order
Cart button to view all cart items (alert)
Order history button to view all orders (alert)
Smooth connection with backend through API calls


SECTION 2: TECHNOLOGIES USED

Backend:
Go
Gin Web Framework
GORM ORM
SQLite

Frontend:
React
Vite
Fetch API

SECTION 3: PROJECT FOLDER STRUCTURE

shopping-cart
backend
go.mod
main.go
database/
database.go
models/
user.go
item.go
cart.go
order.go
controllers/
user_controller.go
item_controller.go
cart_controller.go
order_controller.go
middleware/
auth.go
routes/
routes.go

frontend
index.html
package.json
vite.config.js
src/
main.jsx
App.jsx
api.js
components/
LoginPage.jsx
ItemsPage.jsx

README.md


SECTION 4: HOW TO RUN THE PROJECT

BACKEND SETUP:

Open terminal and go inside the backend folder
cd backend

Download all Go dependencies
go mod tidy

Run the backend server
go run main.go

The backend will start on:
http://localhost:8080

A SQLite database file (shopping_cart.db) will be created automatically.

FRONTEND SETUP:

Open another terminal and go inside the frontend folder
cd frontend

Install dependencies
npm install

Start frontend server
npm run dev

The frontend will start on:
http://localhost:5173

SECTION 5: API ENDPOINTS USED IN THIS PROJECT

Users:
POST /users
GET /users
POST /users/login

Items:
POST /items
GET /items

Carts (requires token):
POST /carts
GET /carts

Orders (requires token):
POST /orders
GET /orders

Authentication:
Send token in the header as
Authorization: Bearer <token>

SECTION 6: FRONTEND USER FLOW
User logs in using username and password
Items page loads
Clicking an item adds it to cart
The top bar contains:
Checkout button: converts cart into an order
Cart button: displays cart items in an alert
Order history button: displays all previous order IDs
When checkout is successful, an alert shows "Order successful"

SECTION 7: PROJECT NOTES
Only one cart is active per user (status: "open")
After checkout, the cart status becomes "ordered" and a new cart will be created as needed
Tokens reset on every login
No inventory management is used (as per assessment instructions)
Fully aligned with the requirements provided in the PDF

SECTION 8: POSTMAN COLLECTION

A complete Postman collection is provided separately.
Import the JSON file into Postman to test all endpoints easily.

SECTION 9: COMPLETION STATUS

All assessment requirements are fully implemented:
Backend complete
Frontend complete
All endpoints implemented
Authentication implemented
Cart and order flow fully functional
UI flow matches the PDF exactly

END OF README
