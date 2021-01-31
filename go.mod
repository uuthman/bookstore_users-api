module github.com/uuthman/bookstore_users-api

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/uuthman/bookstore_oauth-go v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
)

replace github.com/uuthman/bookstore_oauth-go => /home/uthutu/workspaces/uuthman/bookstore_oauth-go
