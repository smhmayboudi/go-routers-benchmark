package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexedwards/flow"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gorilla/mux"
	"github.com/kamalshkeir/kmux"
	"github.com/kamalshkeir/ksmux"
	"github.com/labstack/echo/v4"
)

func BenchmarkKmux(b *testing.B) {
	app := kmux.New()
	app.Get("/test/first/second/third/fourth/fifth", func(c *kmux.Context) {
		c.Text("Hello")
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmux(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/first/second/third/fourth/fifth", func(c *ksmux.Context) {
		c.Text("Hello")
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChi(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/first/second/third/fourth/fifth", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTP(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/first/second/third/fourth/fifth", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/first/second/third/fourth/fifth", func(ctx *gin.Context) {
		ctx.String(200, "Hello")
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEcho(b *testing.B) {
	app := echo.New()
	app.GET("/test/first/second/third/fourth/fifth", func(c echo.Context) error {
		return c.String(200, "Hello")
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFlow(b *testing.B) {
	app := flow.New()
	app.HandleFunc("/test/first/second/third/fourth/fifth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}, "GET")

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFiber(b *testing.B) {
	app := fiber.New()
	app.Get("/test/first/second/third/fourth/fifth", func(c fiber.Ctx) error {
		return c.SendString("Hello")
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	handlerFunc := adaptor.FiberApp(app)
	for i := 0; i < b.N; i++ {
		handlerFunc.ServeHTTP(w, req)
	}
}

func BenchmarkGorilla(b *testing.B) {
	app := mux.NewRouter()
	app.HandleFunc("/test/first/second/third/fourth/fifth", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello")
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKmuxWith1Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:first", func(c *kmux.Context) {
		first := c.Param("first")
		c.Text(fmt.Sprintf("Hello %s", first))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith1Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:first", func(c *ksmux.Context) {
		first := c.Param("first")
		c.Text(fmt.Sprintf("Hello %s", first))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWith1Param(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{first}", func(w http.ResponseWriter, r *http.Request) {
		first := chi.URLParam(r, "first")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello %s", first)))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith1Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{first}/{$}", func(w http.ResponseWriter, r *http.Request) {
		first := r.PathValue("first")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello %s", first)))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith1Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:first", func(ctx *gin.Context) {
		first := ctx.Param("first")
		ctx.String(200, "Hello %s", first)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith1Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:first", func(c echo.Context) error {
		first := c.Param("first")
		return c.String(200, fmt.Sprintf("Hello %s", first))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFlowWith1Param(b *testing.B) {
	app := flow.New()
	app.HandleFunc("/test/:first", func(w http.ResponseWriter, r *http.Request) {
		first := flow.Param(r.Context(), "first")
		fmt.Fprintf(w, "Hello %s", first)
	}, "GET")

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}

}
func BenchmarkFiberWith1Param(b *testing.B) {
	app := fiber.New()
	app.Get("/test/:first", func(c fiber.Ctx) error {
		first := c.Params("first")
		return c.SendString(fmt.Sprintf("Hello %s", first))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	handlerFunc := adaptor.FiberApp(app)
	for i := 0; i < b.N; i++ {
		handlerFunc.ServeHTTP(w, req)
	}
}

func BenchmarkGorillaWith1Param(b *testing.B) {
	app := mux.NewRouter()
	app.HandleFunc("/test/{first}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		first := vars["first"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", first)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKmuxWith2Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:first/:second", func(c *kmux.Context) {
		first := c.Param("first")
		second := c.Param("second")
		c.Text(fmt.Sprintf("Hello %s %s", first, second))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith2Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:first/:second", func(c *ksmux.Context) {
		first := c.Param("first")
		second := c.Param("second")
		c.Text(fmt.Sprintf("Hello %s %s", first, second))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWith2Param(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{first}/{second}", func(w http.ResponseWriter, r *http.Request) {
		first := chi.URLParam(r, "first")
		second := chi.URLParam(r, "second")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello %s %s", first, second)))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith2Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{first}/{second}/{$}", func(w http.ResponseWriter, r *http.Request) {
		first := r.PathValue("first")
		second := r.PathValue("second")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello %s %s", first, second)))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith2Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:first/:second", func(ctx *gin.Context) {
		first := ctx.Param("first")
		second := ctx.Param("second")
		ctx.String(200, "Hello %s %s", first, second)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith2Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:first/:second", func(c echo.Context) error {
		first := c.Param("first")
		second := c.Param("second")
		return c.String(200, fmt.Sprintf("Hello %s %s", first, second))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFlowWith2Param(b *testing.B) {
	app := flow.New()
	app.HandleFunc("/test/:first/:second", func(w http.ResponseWriter, r *http.Request) {
		first := flow.Param(r.Context(), "first")
		second := flow.Param(r.Context(), "second")
		fmt.Fprintf(w, "Hello %s %s", first, second)
	}, "GET")

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFiberWith2Param(b *testing.B) {
	app := fiber.New()
	app.Get("/test/:first/:second", func(c fiber.Ctx) error {
		first := c.Params("first")
		second := c.Params("second")
		return c.SendString(fmt.Sprintf("Hello %s %s", first, second))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	handlerFunc := adaptor.FiberApp(app)
	for i := 0; i < b.N; i++ {
		handlerFunc.ServeHTTP(w, req)
	}
}

func BenchmarkGorillaWith2Param(b *testing.B) {
	app := mux.NewRouter()
	app.HandleFunc("/test/{first}/{second}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		first := vars["first"]
		second := vars["second"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s %s", first, second)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKmuxWith5Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *kmux.Context) {
		first := c.Param("first")
		second := c.Param("second")
		third := c.Param("third")
		fourth := c.Param("fourth")
		fifth := c.Param("fifth")
		c.Text(fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith5Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *ksmux.Context) {
		first := c.Param("first")
		second := c.Param("second")
		third := c.Param("third")
		fourth := c.Param("fourth")
		fifth := c.Param("fifth")
		c.Text(fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWith5Param(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{first}/{second}/{third}/{fourth}/{fifth}", func(w http.ResponseWriter, r *http.Request) {
		first := chi.URLParam(r, "first")
		second := chi.URLParam(r, "second")
		third := chi.URLParam(r, "third")
		fourth := chi.URLParam(r, "fourth")
		fifth := chi.URLParam(r, "fifth")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth)))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith5Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{first}/{second}/{third}/{fourth}/{fifth}/{$}", func(w http.ResponseWriter, r *http.Request) {
		first := r.PathValue("first")
		second := r.PathValue("second")
		third := r.PathValue("third")
		fourth := r.PathValue("fourth")
		fifth := r.PathValue("fifth")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth)))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith5Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:first/:second/:third/:fourth/:fifth", func(ctx *gin.Context) {
		first := ctx.Param("first")
		second := ctx.Param("second")
		third := ctx.Param("third")
		fourth := ctx.Param("fourth")
		fifth := ctx.Param("fifth")
		ctx.String(200, "Hello %s %s %s %s %s", first, second, third, fourth, fifth)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith5Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:first/:second/:third/:fourth/:fifth", func(c echo.Context) error {
		first := c.Param("first")
		second := c.Param("second")
		third := c.Param("third")
		fourth := c.Param("fourth")
		fifth := c.Param("fifth")
		return c.String(200, fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFlowWith5Param(b *testing.B) {
	app := flow.New()
	app.HandleFunc("/test/:first/:second/:third/:fourth/:fifth", func(w http.ResponseWriter, r *http.Request) {
		first := flow.Param(r.Context(), "first")
		second := flow.Param(r.Context(), "second")
		third := flow.Param(r.Context(), "third")
		fourth := flow.Param(r.Context(), "fourth")
		fifth := flow.Param(r.Context(), "fifth")
		fmt.Fprintf(w, "Hello %s %s %s %s %s", first, second, third, fourth, fifth)
	}, "GET")

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/forth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFiberWith5Param(b *testing.B) {
	app := fiber.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c fiber.Ctx) error {
		first := c.Params("first")
		second := c.Params("second")
		third := c.Params("third")
		fourth := c.Params("fourth")
		fifth := c.Params("fifth")
		return c.SendString(fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth))
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first", nil)
	w := httptest.NewRecorder()

	handlerFunc := adaptor.FiberApp(app)
	for i := 0; i < b.N; i++ {
		handlerFunc.ServeHTTP(w, req)
	}
}

func BenchmarkGorillaWith5Param(b *testing.B) {
	app := mux.NewRouter()
	app.HandleFunc("/test/{first}/{second}/{third}/{fourth}/{fifth}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		first := vars["first"]
		second := vars["second"]
		third := vars["third"]
		fourth := vars["fourth"]
		fifth := vars["fifth"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s %s %s %s %s", first, second, third, fourth, fifth)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/first/second/third/fourth/fifth", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}
