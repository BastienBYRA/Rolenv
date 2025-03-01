# Rolenv
🚢📦 A simple way to run Docker containers by defining an environment variable file 

*Rolenv's sole purpose today is to launch Docker containers using a .env file and validate their configuration.*

## Installation

TODO

### Configuration

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
| ROLENV_PRIVILEGED      | `false`       | boolean : `true`, `false`                                                                         | `--privileged`             |
| ROLENV_RESTART_POLICY  | `no`          | string : `no`, `on-failure`, `always`, `unless-stopped`                                           | `--restart`                |
| ROLENV_RESTART_POLICY_MAX_RETRIES | `0` | integer : positive number                                                                         | `--restart-max-attempts`   |
| ROLENV_USER            | (image default user) | string : `user`, `user:group`, `uid:gid`                                                   | `--user`                   |
| ROLENV_ENV_LIST        | ``            | All environment variables that do not start with "ROLENV_"                                        | `-e/--env`                 |
| ROLENV_MEMORY_LIMIT    | `-1` (unlimited)| integer : memory in bytes                                                                         | `--memory`                 |
| ROLENV_CPU_CORE_LIMIT  | `-1` (unlimited)| integer : number of CPU cores                                                                     | `--cpus`                   |
| ROLENV_READONLY        | `false`       | boolean : `true`, `false`                                                                         | `--read-only`              |
| ROLENV_SECURITY_OPTS   | ``            | Comma-separated list of strings : `no-new-privileges;seccomp=unconfined`                          | `--security-opt`           |
| ROLENV_VOLUMES         | ``            | Comma-separated list of key-value pairs : `./test:/tmp/test;data-rolenv-test:/a-folder`           | `-v/--volume`              |

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