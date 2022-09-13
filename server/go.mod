module github.com/Dittiya/go-music-scrapper/server

go 1.19

replace github.com/Dittiya/go-music-scrapper/config => ../config

require (
	github.com/Dittiya/go-music-scrapper/config v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.37.1
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.40.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)
