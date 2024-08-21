package rest

import (
	"net/http"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
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
	mux.HandleFunc("/profile", h.methodHandler(h.usersHandler.Profile, http.MethodGet))
	mux.HandleFunc("/users", h.methodHandler(h.usersHandler.GetUserByName, http.MethodGet)) // Added route for GetUserByName

	// Initialize AuthMiddleware
	authMiddleware := NewAuthMiddleware(h.usersService)

	// Protected Books routes with token validation
	mux.HandleFunc("/books/create", authMiddleware.Middleware(h.methodHandler(h.booksHandler.Create, http.MethodPost)))
	mux.HandleFunc("/books", h.methodHandler(h.booksHandler.GetAll, http.MethodGet))
	mux.HandleFunc("/books/", h.methodHandler(h.booksHandler.GetByID, http.MethodGet))
	mux.HandleFunc("/books/update", authMiddleware.Middleware(h.methodHandler(h.booksHandler.Update, http.MethodPut)))
	mux.HandleFunc("/books/delete", authMiddleware.Middleware(h.methodHandler(h.booksHandler.Delete, http.MethodDelete)))

	// Apply CORS middleware
	return CORS(mux)
}

func (h *Handler) methodHandler(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}
