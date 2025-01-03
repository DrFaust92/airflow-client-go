module github.com/drfaust92/airflow-client-go/airflow

go 1.23

replace github.com/apache/airflow-client-go/airflow => ./airflow

require golang.org/x/oauth2 v0.24.0
