package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/kamalshkeir/kmux"
	"github.com/kamalshkeir/ksmux"
	"github.com/labstack/echo/v4"
)

func BenchmarkKmux(b *testing.B) {
	app := kmux.New()
	app.Get("/test/smhmayboudi/hello/bye", func(c *kmux.Context) {
		c.Text("Hello")
	})

	req := httptest.NewRequest("GET", "/test/smhmayboudi/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmux(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/smhmayboudi/hello/bye", func(c *ksmux.Context) {
		c.Text("Hello")
	})

	req := httptest.NewRequest("GET", "/test/smhmayboudi/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChi(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/smhmayboudi/hello/bye", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	req := httptest.NewRequest("GET", "/test/smhmayboudi/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTP(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/smhmayboudi/hello/bye", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	req := httptest.NewRequest("GET", "/test/smhmayboudi/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/smhmayboudi/hello/bye", func(ctx *gin.Context) {
		ctx.String(200, "Hello")
	})

	req := httptest.NewRequest("GET", "/test/smhmayboudi/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEcho(b *testing.B) {
	app := echo.New()
	app.GET("/test/smhmayboudi/hello/bye", func(c echo.Context) error {
		return c.String(200, "Hello")
	})

	req := httptest.NewRequest("GET", "/test/smhmayboudi/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKmuxWithParam(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:something", func(c *kmux.Context) {
		c.Text("Hello " + c.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWithParam(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:something", func(c *ksmux.Context) {
		c.Text("Hello " + c.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWithParam(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{something}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + chi.URLParam(r, "something")))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWithParam(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{something}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + r.PathValue("something")))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWithParam(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:something", func(ctx *gin.Context) {
		ctx.String(200, "Hello "+ctx.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWithParam(b *testing.B) {
	app := echo.New()
	app.GET("/test/:something", func(c echo.Context) error {
		return c.String(200, "Hello "+c.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKmuxWith2Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:something/:another", func(c *kmux.Context) {
		c.Text("Hello " + c.Param("something") + c.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith2Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:something/:another", func(c *ksmux.Context) {
		c.Text("Hello " + c.Param("something") + c.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith2Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{something}/{another}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello, " + r.PathValue("something") + r.PathValue("another")))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith2Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:something/:another", func(ctx *gin.Context) {
		ctx.String(200, "Hello "+ctx.Param("something")+ctx.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith2Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:something/:another", func(c echo.Context) error {
		return c.String(200, "Hello "+c.Param("something")+c.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKmuxWith5Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *kmux.Context) {
		c.Text("Hello " + c.Param("first") + c.Param("second") + c.Param("third") + c.Param("fourth") + c.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith5Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *ksmux.Context) {
		c.Text("Hello " + c.Param("first") + c.Param("second") + c.Param("third") + c.Param("fourth") + c.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith5Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{first}/{second}/{third}/{fourth}/{fifth}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "Hello, "+r.PathValue("first")+r.PathValue("second")+r.PathValue("third")+r.PathValue("fourth")+r.PathValue("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith5Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:first/:second/:third/:fourth/:fifth", func(ctx *gin.Context) {
		ctx.String(200, "Hello "+ctx.Param("first")+ctx.Param("second")+ctx.Param("third")+ctx.Param("fourth")+ctx.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith5Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:first/:second/:third/:fourth/:fifth", func(c echo.Context) error {
		return c.String(200, "Hello "+c.Param("first")+c.Param("second")+c.Param("third")+c.Param("fourth")+c.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}
