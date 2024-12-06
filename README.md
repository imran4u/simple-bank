# simple-bank

## steps to download mirgation on ubuntu
    
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