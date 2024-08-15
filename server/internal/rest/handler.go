package rest

import (
	"net/http"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
	"github.com/gorilla/handlers"
)

type Handler struct {
	booksHandler *BooksHandler
	usersHandler *UsersHandler
	usersService *service.Users
}

func NewHandler(booksService *service.Books, usersService *service.Users) *Handler {
	return &Handler{
		booksHandler: NewBooksHandler(booksService),
		usersHandler: NewUsersHandler(usersService),
		usersService: usersService,
	}
}

func (h *Handler) InitRouter() http.Handler {
	mux := http.NewServeMux()

	// Users routes
	mux.HandleFunc("/signup", h.methodHandler(h.usersHandler.SignUp, http.MethodPost))
	mux.HandleFunc("/signin", h.methodHandler(h.usersHandler.SignIn, http.MethodPost))

	// Initialize AuthMiddleware
	authMiddleware := NewAuthMiddleware(h.usersService)

	// Protected Books routes with token validation
	mux.HandleFunc("/books/create", authMiddleware.Middleware(h.methodHandler(h.booksHandler.Create, http.MethodPost)))
	mux.HandleFunc("/books", h.methodHandler(h.booksHandler.GetAll, http.MethodGet))   // For getting all books
	mux.HandleFunc("/books/", h.methodHandler(h.booksHandler.GetByID, http.MethodGet)) // For getting a book by ID
	mux.HandleFunc("/books/update", authMiddleware.Middleware(h.methodHandler(h.booksHandler.Update, http.MethodPut)))
	mux.HandleFunc("/books/delete", authMiddleware.Middleware(h.methodHandler(h.booksHandler.Delete, http.MethodDelete)))

	// Apply CORS middleware
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(mux)
}

// methodHandler is a helper function to restrict handler to a specific HTTP method.
func (h *Handler) methodHandler(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}
