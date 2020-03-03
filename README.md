# VIN Decoder

> :red_car: Decoding vehicle identification number

## Development

Build the binary

```sh
make
```

Start postgres

```sh
docker-compose up -Vd
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:@127.0.0.1:5432/vin-decoder_test\?sslmode=disable up
```

Run the web server

```sh
./bin/server
```

## Usage

For example, you get information about this amazing Tesla Model X

```sh
http http://localhost:8080/api/v1/vin-decoder/5YJXCCE40GF010543
```

```json
{
  "vehicle": {
    "check_digit": true,
    "country": "United States",
    "make": "Tesla",
    "manufacturer": "Tesla, Inc.",
    "region": "North America",
    "serial": "010543",
    "year": 2016
  },
  "vin": {
    "vds": "XCCE40",
    "vis": "GF010543",
    "wmi": "5YJ"
  }
}
```

## License

Project released under the terms of the MIT [license](./LICENSE).
