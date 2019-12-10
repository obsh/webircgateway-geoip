package main

import (
	"github.com/kiwiirc/webircgateway/pkg/webircgateway"
	"github.com/oschwald/geoip2-golang"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var db *geoip2.Reader

func Start(gateway *webircgateway.Gateway, pluginsQuit *sync.WaitGroup) {
	gateway.Log(2, "GeoIP plugin loading")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		gateway.Log(3, err.Error())
		pluginsQuit.Done()
		return
	}

	ipdbFileName := dir + "/" + "GeoLite2-Country.mmdb"
	gateway.Log(1, "Looking for the IPDB file: "+ipdbFileName)
	db, err = geoip2.Open(ipdbFileName)
	if err != nil {
		gateway.Log(3, err.Error())
		pluginsQuit.Done()
		return
	}
	gateway.Log(1, "GeoIP DB opened")

	webircgateway.HookRegister("irc.connection.pre", hookIrcConnectionPre)
	webircgateway.HookRegister("gateway.closing", func(hook *webircgateway.HookGatewayClosing) {
		go func() {
			gateway.Log(1, "GeoIP DB closed")
			db.Close()
			pluginsQuit.Done()
		}()
	})
}

func hookIrcConnectionPre(hook *webircgateway.HookIrcConnectionPre) {
	ip := net.ParseIP(hook.Client.RemoteAddr)
	record, err := db.City(ip)
	if err != nil {
		hook.Client.Log(3, "Cannot find information about IP: "+ip.String())
		hook.Client.Log(3, err.Error())
		return
	}

	hook.Client.Gateway.Log(2, "GeoIP Plugin: "+record.Country.IsoCode)
	makeRealNameReplacements(hook.Client, record)
}

func makeRealNameReplacements(client *webircgateway.Client, record *geoip2.City) {
	realName := client.IrcState.RealName
	realName = strings.Replace(realName, "%country", record.Country.IsoCode, -1)
	client.IrcState.RealName = realName
}
