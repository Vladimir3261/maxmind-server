## Maxmind server REST API
GET /?ip=8.8.8.8
response example:

```json
    {
        "Status": 200,
        "Data": {
            "Continent": {
                "Code": "EU",
                "GeoNameID": 6255148,
                "Names": {
                    "de": "Europa",
                    "en": "Europe",
                    "es": "Europa",
                    "fr": "Europe",
                    "ja": "ヨーロッパ",
                    "pt-BR": "Europa",
                    "ru": "Европа",
                    "zh-CN": "欧洲"
                }
            },
            "Country": {
                "GeoNameID": 2921044,
                "IsInEuropeanUnion": true,
                "IsoCode": "DE",
                "Names": {
                    "de": "Deutschland",
                    "en": "Germany",
                    "es": "Alemania",
                    "fr": "Allemagne",
                    "ja": "ドイツ連邦共和国",
                    "pt-BR": "Alemanha",
                    "ru": "Германия",
                    "zh-CN": "德国"
                }
            },
            "RegisteredCountry": {
                "GeoNameID": 6252001,
                "IsInEuropeanUnion": false,
                "IsoCode": "US",
                "Names": {
                    "de": "USA",
                    "en": "United States",
                    "es": "Estados Unidos",
                    "fr": "États-Unis",
                    "ja": "アメリカ合衆国",
                    "pt-BR": "Estados Unidos",
                    "ru": "США",
                    "zh-CN": "美国"
                }
            },
            "RepresentedCountry": {
                "GeoNameID": 0,
                "IsInEuropeanUnion": false,
                "IsoCode": "",
                "Names": null,
                "Type": ""
            },
            "Traits": {
                "IsAnonymousProxy": false,
                "IsSatelliteProvider": false
            }
        },
        "Error": ""
    }
```
Actions:
### How to run:

`./maxmind-server [action] [options]`


Avaliable actions:

- update - update maxmind database file
- run [database_file_path] if database fille is "-" server will update database then run using updated file

### Configuration
rename config.ini.example to config.ini and edit it:

    database_url - URL for database download
    database_dir - dirctory for store all downloads and database files
    default_db_file - Default db file, if it not provided in cli argumants
    http_server - http server listen port like :3000
## Apache benchmark results

```text
 Server Software:        Go HTTP server
 Server Hostname:        localhost
 Server Port:            3000

 Document Path:          /?ip=138.68.80.28/
 Document Length:        78 bytes

 Concurrency Level:      1000
 Time taken for tests:   7.889 seconds
 Complete requests:      100000
 Failed requests:        0
 Total transferred:      18600000 bytes
 HTML transferred:       7800000 bytes
 Requests per second:    12676.02 [#/sec] (mean)
 Time per request:       78.889 [ms] (mean)
 Time per request:       0.079 [ms] (mean, across all concurrent requests)
 Transfer rate:          2302.48 [Kbytes/sec] received

 Connection Times (ms)
               min  mean[+/-sd] median   max
 Connect:        7   40  98.1     31    1075
 Processing:    12   39   9.8     39      91
 Waiting:        5   28   9.1     27      79
 Total:         28   78  99.0     71    1121

 Percentage of the requests served within a certain time (ms)
   50%     71
   66%     75
   75%     77
   80%     78
   90%     82
   95%     86
   98%     92
   99%    115
  100%   1121 (longest request)
  ```