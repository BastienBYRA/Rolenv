# Rolenv
🚢📦 A simple way to run Docker containers by defining an environment variable file 

*Rolenv's sole purpose today is to launch Docker containers using a .env file and validate their configuration.*

## Installation

### Docker
Docker is easiest way to run out Rolenv.
```bash
docker pull ghcr.io/bastienbyra/rolenv-host-socket:latest
```

You can then run it
```bash
# Validate the configuration of your container
docker run --rm -v $(pwd)/path/to/rolenv.env:/rolenv.env -v /var/run/docker.sock:/var/run/docker.sock ghcr.io/bastienbyra/rolenv-host-socket:latest validate --config /rolenv.env

# Run a container
docker run --rm -v $(pwd)/path/to/rolenv.env:/rolenv.env -v /var/run/docker.sock:/var/run/docker.sock ghcr.io/bastienbyra/rolenv-host-socket:latest run --config /rolenv.env
```

### Binary 
#### Linux, WSL
```bash
export ROLENV_LATEST_VERSION=$(curl -L https://raw.githubusercontent.com/BastienBYRA/Rolenv/master/version)
curl -L https://github.com/BastienBYRA/Rolenv/releases/download/v{ROLENV_LATEST_VERSION}/rolenv-linux-amd64 -o rolenv
unset ROLENV_LATEST_VERSION
```

#### MacOS
```bash
export ROLENV_LATEST_VERSION=$(curl -L https://raw.githubusercontent.com/BastienBYRA/Rolenv/master/version)
curl -L https://github.com/BastienBYRA/Rolenv/releases/download/v{ROLENV_LATEST_VERSION}/rolenv-darwin-amd64 -o rolenv
unset ROLENV_LATEST_VERSION
```

#### Windows
```powershell
$ROLENV_LATEST_VERSION = Invoke-WebRequest -Uri "https://raw.githubusercontent.com/BastienBYRA/Rolenv/master/version" -UseBasicParsing | Select-Object -ExpandProperty Content
$ROLENV_LATEST_VERSION = $ROLENV_LATEST_VERSION.Trim()

$downloadUrl = "https://github.com/BastienBYRA/Rolenv/releases/download/v$ROLENV_LATEST_VERSION/rolenv-windows-amd64"
Invoke-WebRequest -Uri $downloadUrl -OutFile "rolenv"

Remove-Variable ROLENV_LATEST_VERSION
```

### From source (Go)
```bash
# Build
go get .
go build -o rolenv

# Run
chmod +x rolenv
./rolenv
```

## Configuration

Rolenv is launched by running the `rolenv` command. For proper operation, it expects an `*.env` file.

By default, Rolenv looks for a file named `rolenv.env` in the root of its execution directory.

It is possible to specify the path to the Rolenv configuration file using `rolenv --config-file /path/to/.env`.

### Configuring the Container

The container definition is done exclusively through the use of environment variables. Below is the list of environment variables:

| VARIABLE               | DEFAULT       | EXPECTED VALUES                                                                                   | DOCKER EQUIVALENT          |
|------------------------|---------------|---------------------------------------------------------------------------------------------------|----------------------------|
| ROLENV_NAME            | (required)    | string : `my-cont-name`, `whatever`                                                               | `--name`                   |
| ROLENV_IMAGE           | (required)    | string : valid image name                                                                         | `IMAGE`                    |
| ROLENV_PORTS           | ``            | Comma-separated list of key-value pairs : `8080:80`, `2222:22;8080:80`                            | `-p/--publish`             |
| ROLENV_NETWORK         | `default`     | string : `bridge`, `host`, `none`, custom network name                                            | `--network`                |
| ROLENV_HOSTS           | ``            | Comma-separated list of key-value pairs : `host1:192.168.1.1`, `host2:192.168.1.2;host1:192.168.1.1` | `--add-host`               |
| ROLENV_ENTRYPOINT      | ``            | Comma-separated list of strings : `/bin/bash`, `python;app.py`                                    | `--entrypoint`             |
| ROLENV_COMMAND         | ``            | Comma-separated list of strings : `arg1;arg2`                                                     | `COMMAND`                  |
| ROLENV_HOSTNAME        | (container ID)| string : `my-hostname`                                                                            | `--hostname`               |
| ROLENV_PRIVILEGED      | `false`       | boolean : `true`, `false`, `yes`, `no`                                                            | `--privileged`             |
| ROLENV_RESTART_POLICY  | `no`          | string : `no`, `on-failure`, `always`, `unless-stopped`                                           | `--restart`                |
| ROLENV_RESTART_POLICY_MAX_RETRIES | `0` | integer : positive number                                                                         | `--restart-max-attempts`   |
| ROLENV_USER            | (image default user) | string : `user`, `user:group`, `uid:gid`                                                   | `--user`                   |
| ROLENV_MEMORY_LIMIT    | `-1` (unlimited)| integer : memory in bytes                                                                         | `--memory`                 |
| ROLENV_CPU_CORE_LIMIT  | `-1` (unlimited)| integer : number of CPU cores                                                                     | `--cpus`                   |
| ROLENV_READONLY        | `false`       | boolean : `true`, `false`, `yes`, `no`                                                              | `--read-only`              |
| ROLENV_SECURITY_OPTS+   | ``            | Comma-separated list of strings : `no-new-privileges;seccomp=unconfined`                          | `--security-opt`           |
| ROLENV_VOLUMES         | ``            | Comma-separated list of key-value pairs : `./test:/tmp/test;data-rolenv-test:/a-folder`           | `-v/--volume`              |

### Application environment variables
All environment variables in the `*.env` file that do not start with ROLENV_ will be added to the container as application environment variables.

## Start the Program

To start the program, use `rolenv` in your terminal.

```bash
$ rolenv --help
```

## Commands

### Run

The `run` command is the main command. It runs a Docker container in detached mode using the provided configuration file.

```bash
# Run a container
$ rolenv run
# Providing a configuration file
$ rolenv run --config-file /path/to/my/rolenv.env/or/.env/or/whatever.env
```

### Validate

The `validate` command checks if the configuration is valid and can be used to run a container.

```bash
# Validate a container configuration
$ rolenv validate
# Providing a configuration file
$ rolenv validate --config-file /path/to/my/rolenv.env/or/.env/or/whatever.env
```

## License

The project is distributed under the **GNU AFFERO GENERAL PUBLIC LICENSE v3 (AGPLv3)**. The goal is to foster rapid and collaborative development of Rolenv by encouraging all users to share their improvements and contributions.