resource "google_cloud_run_service" "api" {
  name     = "learn-golang"
  location = var.region

  template {
    spec {
      containers {
        image = "httpd"
        ports {
          container_port = 80
        }
        #         env {
        #           name  = "DB_HOST"
        #           value = google_sql_database_instance.postgres.connection_name
        #         }
        #         env {
        #           name  = "DB_NAME"
        #           value = google_sql_database.default.name
        #         }
        #         env {
        #           name  = "DB_USER"
        #           value = google_sql_user.postgres_user.name
        #         }
        #         env {
        #           name  = "DB_PASSWORD"
        #           value = google_sql_user.postgres_user.password
        #         }
      }
    }
  }
}