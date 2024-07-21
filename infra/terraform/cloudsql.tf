resource "google_sql_database_instance" "postgres" {
  database_version = "POSTGRES_16"
  settings {
    tier = "db-f1-micro"
    availability_type = "ZONAL"
    disk_type = "PD_HDD"

    backup_configuration {
      enabled = false
    }
  }
  deletion_protection = false
}

resource "google_sql_database" "default" {
  name     = var.db_name
  instance = google_sql_database_instance.postgres.name
}

resource "google_sql_user" "postgres_user" {
  name     = var.db_user
  instance = google_sql_database_instance.postgres.name
  password = var.db_password
}