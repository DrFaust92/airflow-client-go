module github.com/drfaust92/airflow-client-go

go 1.23

replace github.com/apache/airflow-client-go/airflow => ./airflow

require (
	github.com/apache/airflow-client-go/airflow v0.0.0-20230210234754-8ce0b39cfbb2
	github.com/stretchr/testify v1.6.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
