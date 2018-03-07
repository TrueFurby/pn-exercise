## Exercise 1

### Programming
Retrieve weather data directly through API.

```sh
export OPENWEATHER_API_KEY="xxxxx"
export CITY_NAME="Bratislava"
go run main.go
```

### Ansible
Install the Docker service using Ansible and enable logging to Docker host's syslog file.

```sh
ansible-playbook -i "localhost," -c local site.yml
```

### Docker
Retrieve weather data by running the program in Docker container.

```sh
docker build -t getweather .
docker run --rm -e OPENWEATHER_API_KEY="xxxxx" -e CITY_NAME="Bratislava" getweather
```
