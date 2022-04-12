package middleware

import (
	"net"
	"net/http"
	"strings"
)

// TrustedSubnetMiddleware middleware метод ограничивающий доступ к ресурсу /api/internal/stats только
// с доверенных подсетей, переданных в заголовке запроса X-Real-IP,
// в противном случае возвращает статус ответа 403 Forbidden.
func TrustedSubnetMiddleware(network string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/api/internal/stats" {
				next.ServeHTTP(w, r)
			}
			remoteIP := getRemoteIPAddr(r)
			if remoteIP == nil {
				http.Error(w, "", http.StatusForbidden)
				return
			}
			_, subnet, err := net.ParseCIDR(network)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !subnet.Contains(remoteIP) {
				http.Error(w, "", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// getRemoteIPAddr метод-helper получения клиентского IP-адреса
func getRemoteIPAddr(r *http.Request) net.IP {
	xri := r.Header.Get("X-Real-IP")
	remoteIP := net.ParseIP(xri)
	if remoteIP == nil {
		ips := r.Header.Get("X-Forwarded-For")
		splitIPs := strings.Split(ips, ",")
		xri = splitIPs[0]
		remoteIP = net.ParseIP(xri)
	}
	return remoteIP
}
