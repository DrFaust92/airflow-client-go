module github.com/drfaust92/airflow-client-go

go 1.23

replace github.com/drfaust92/airflow-client-go/airflow => ./airflow

require (
	github.com/drfaust92/airflow-client-go/airflow v0.0.0-20251117011734-cf1a24cd6aa4
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
