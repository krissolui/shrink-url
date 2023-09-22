package server

import (
	"net/http"
	"net/url"
	"shrink-url/internal/config"
	"shrink-url/internal/util"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", Handler())

	return mux
}

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check path is exact match
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// check http method
		switch r.Method {
		case http.MethodGet:
			EncryptUrl(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func EncryptUrl(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.URL.Query().Get("url")
	if originalUrl == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	if !util.ValidateUrl(originalUrl) {
		http.Error(w, "url format invalid", http.StatusBadRequest)
		return
	}

	settings := config.Cfg.ShrinkUrlSettings
	encryptor := util.NewCryptography()
	cipher := encryptor.EncryptUrl(originalUrl, settings.MaxLength)
	encryptedUrl, _ := url.JoinPath(settings.BaseUrl, cipher)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(encryptedUrl))
}
