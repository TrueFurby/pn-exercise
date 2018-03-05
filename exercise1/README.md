## Exercise 1

# Get weather info

```sh
export OPENWEATHER_API_KEY="xxxxx"
export CITY_NAME="Bratislava"
go run main.go
```

# Run with Docker

```sh
docker build -t getweather .
docker run --rm -e OPENWEATHER_API_KEY="xxxxx" -e CITY_NAME="Bratislava" getweather
```
