package proxyapp

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Proxy struct {
	HostTarget map[string]string
	AppsCount  int
	HostProxy  map[string]*httputil.ReverseProxy
	CurrentApp int
}

func NewProxy() *Proxy {

	hostTarget := viper.GetStringMapString("proxy.apps")

	return &Proxy{
		HostTarget: hostTarget,
		AppsCount:  len(hostTarget),
		HostProxy:  map[string]*httputil.ReverseProxy{},
		CurrentApp: 0,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.CurrentApp++
	p.CurrentApp %= p.AppsCount
	host := strconv.Itoa(p.CurrentApp)

	log.Info().Msg("to -> " + p.HostTarget[host])

	if fn, ok := p.HostProxy[host]; ok {
		fn.ServeHTTP(w, r)
		return
	}

	if target, ok := p.HostTarget[host]; ok {
		remoteUrl, err := url.Parse(target)
		if err != nil {
			log.Err(err).Msg("target parse fail")
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		p.HostProxy[host] = proxy
		proxy.ServeHTTP(w, r)
		return
	}
	w.Write([]byte("403: Host forbidden " + host))
}
