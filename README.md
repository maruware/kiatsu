kiatsu
==========

Kiatsu can show you pressure value.


## Usage

### Requirements

Regist OpenWeatherMap. And set api key to env

```
export OPEN_WEATHER_MAP_API_KEY=1111111111
```

### Execution

```
$ kiatsu
1001hPa@Tokyo
$ kiatsu Osaka
1006hPa@Osaka
$ kiatsu --save Nagoya
1006hPa@Nagoya
$ kiatsu
1006hPa@Nagoya
$ kiatsu --reset
$ kiatsu
1001hPa@Tokyo
```

## Install

```
$ go get github.com/maruware/kiatsu
```
