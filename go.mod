module github.com/utunia/api

replace github.com/utunia/api => ./

go 1.21.4

require (
	github.com/gofiber/fiber/v2 v2.51.0
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/google/uuid v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/nedpals/supabase-go v0.4.0
	golang.org/x/crypto v0.7.0
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/nedpals/postgrest-go v0.1.3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.50.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
