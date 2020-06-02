### Usage

#### setup postgresql orm -> pg
```
# restart postgresql
sudo /etc/init.d/postgresql restart
go get github.com/go-pg/pg/v10
```

##### setup air

```
go build demo.go
air -c .air.conf    # Live reload for Go apps https://github.com/cosmtrek/air.git
                    # ./.air.conf to see air config
```