# simple-bank

## steps to download mirgation on ubuntu ( [github-ref](https://github.com/golang-migrate/migrate))
    
**Add the GPG key for the repository:**

    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -

 **Add the repository:**

    sudo add-apt-repository "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -c | awk '{print $2}') main"

**Update your package list:**

    sudo apt-get update

**Install migrate:**

    sudo apt-get install migrate

** Check migration installed **

    migration -version

## Sqlc : [sqlc.dev](https://sqlc.dev/)
[postgresql sqlc documents](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html)

[github](https://github.com/sqlc-dev/sqlc)

install on ubuntu : ```sudo snap install sqlc```

commands : ```sqlc help```

sqlc configuration : https://docs.sqlc.dev/en/latest/reference/config.html#version-2

