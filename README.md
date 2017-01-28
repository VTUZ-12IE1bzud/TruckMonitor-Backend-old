#Truck Monitor Backend

## Сборка

1. Скачать Gin

    ```sh
    $ go get gopkg.in/gin-gonic/gin.v1
    $ go get github.com/lib/pq
    $ go get github.com/dgrijalva/jwt-go
    ```

## Настройка сервера
Сервер - [Ubuntu Server 16.04.1 x64](http://releases.ubuntu.com/16.04/ubuntu-16.04.1-server-amd64.iso)

###База данны
1. Установка
    ```sh
    $ sudo apt-get install postgresql postgresql-contrib
    ```
  
2. Настройка
    ```sh
    $ sudo su postgres -c psql postgres
    
    $ postgres=# ALTER USER postgres WITH PASSWORD 'ПАРОЛЬ';
    $ postgres=# \q
    ```
    Изменить пароль у пользователя postgres (пароли должны быть одинаковыми):
    ```sh
    $ sudo passwd -d postgres
    $ sudo su postgres -c passwd
    ```
    Настройка доступа к БД:
    ```sh
    $ sudo nano /etc/postgresql/9.5/main/postgresql.conf
    ```
    Расскоментировать строчку и открыть доступ:
    ```
    listen_addresses = '*'
    ```
    ```sh
    $ sudo nano /etc/postgresql/9.5/main/pg_hba.conf
    ```
    Добавить строчку:
    ```
    host                     all                   all               0.0.0.0/0    md5
    ```
    Перезапустить postgres:
    ```sh
    $ sudo service postresql restart
    ```