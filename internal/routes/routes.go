package routes

//SET UP ALL PRODUCT AND HEALTH RELATED ROUTES , also enable logger, timout, cors.
import ( //import files from proj, and libraries

	ServerConfig "github.com/aleksander-sienkiewicz/dynamodb-go-crud/config"
	HealthHandler "github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/handlers/health"
	ProductHandler "github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/handlers/product"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/repository/adapter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	config *Config
	router *chi.Mux //chi.mux
}

func NewRouter() *Router { //make new router
	return &Router{ //return router from this func, we say &router cuz we refere to struct above
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

/*
takes params
returns the chi mux
*/
func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigsRouters()       //call func below
	r.RouterHealth(repository)  //func from below
	r.RouterProduct(repository) //func from below as well, pass repository

	return r.router //return router
}

/*
this and func above are struct methods.
*/
func (r *Router) setConfigsRouters() {
	r.EnableCORS() //YASSSSS
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP() //yuppppp
}

/*
 */
func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := HealthHandler.NewHandler(repository)
	//here we create our routes w/ /health
	r.router.Route("/health", func(route chi.Router) {
		route.Post("/", handler.Post)       //handle post operations
		route.Get("/", handler.Get)         //get operations
		route.Put("/", handler.Put)         //put
		route.Delete("/", handler.Delete)   //del
		route.Options("/", handler.Options) //options
	})
}

/*
all our routes will be product
*/
func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := ProductHandler.NewHandler(repository)
	//define all routes w/ /product
	r.router.Route("/product", func(route chi.Router) {
		route.Post("/", handler.Post) //same same
		route.Get("/", handler.Get)
		route.Get("/{ID}", handler.Get) //ur always handling routes eh?
		route.Put("/{ID}", handler.Put) //good thing im learning this i guess!
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

// doesnt take anything, still struct method, thats all of them
func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger) //chi gives a logger to us, thats what we use here k
	return r                        //router
}

/*ENABLETIMEOUT
 */
func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r //return router
}

/*ENABLECORS
 */
func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r //return router
}

/*ENABLE RECOVER
 */
func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r //return router
}

/*ENABLERequestid
 */
func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r //return router
}

/*ENABLErealip
 */
func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r //return router
}
