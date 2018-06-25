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
