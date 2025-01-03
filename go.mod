module github.com/drfaust92/airflow-client-go

go 1.13

replace github.com/apache/airflow-client-go/airflow => ./airflow

require (
	github.com/apache/airflow-client-go/airflow v0.0.0-20230210234754-8ce0b39cfbb2
	github.com/stretchr/testify v1.6.1
)
