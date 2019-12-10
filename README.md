# GeoIP plugin for Webircgateway

See kiwiirc/webircgateway https://github.com/kiwiirc/webircgateway

## Installation

```bash
# first clone webircgateway
git clone https://github.com/kiwiirc/webircgateway.git

# clone plugin
git clone https://github.com/obsh/webircgateway-geoip.git

# create folder for geoip plugin
mkdir webircgateway/plugins/geoip

# copy plugin files
cp webircgateway-geoip/plugin.go webircgateway/plugins/geoip
cp webircgateway-geoip/GeoLite2-Country.mmdb webircgateway/GeoLite2-Country.mmdb

# compile webircgateway with plugin
cd webircgateway && make
```

## Usage

- enable plugin in webircgateway config.conf:
```
[plugins]
plugins/geoip.so
```

The plugin substitues '%country' macro in realname with ISO 3166-2 two letter country code detected by client's IP address.

## Notes

This product includes GeoLite2 data created by MaxMind, available from
<a href="https://www.maxmind.com">https://www.maxmind.com</a>.