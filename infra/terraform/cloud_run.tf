resource "google_cloud_run_service" "api" {
  name     = "learn-golang"
  location = var.region

  template {
    spec {
      # TODO: Cloud SQLに繋がらないので設定を確認する
      # https://cloud.google.com/sql/docs/postgres/connect-run?hl=ja
      containers {
        image = "asia-northeast1-docker.pkg.dev/playground-430113/learn-golang/api:latest"

        ports {
          container_port = 8080
        }
        env {
          name  = "POSTGRES_HOST"
          value = "/cloudsql/${google_sql_database_instance.postgres.connection_name}"
        }
        env {
          name  = "POSTGRES_PORT"
          value = "5432"
        }
        env {
          name  = "POSTGRES_DB"
          value = google_sql_database.default.name
        }
        env {
          name  = "POSTGRES_USER"
          value = google_sql_user.postgres_user.name
        }
        env {
          name  = "POSTGRES_PASSWORD"
          value = google_sql_user.postgres_user.password
        }
      }

      # 本当はcloudbuild等でデプロイ時にmigrationを実行したい
      # この方法はうまくいかなかったので、コンソールからDDLを実行した
      #       containers {
      #         image = "asia-northeast1-docker.pkg.dev/playground-430113/learn-golang/migrate:latest"
      #
      #         env {
      #           name  = "POSTGRES_HOST"
      #           value = "/cloudsql/${google_sql_database_instance.postgres.connection_name}"
      #         }
      #         env {
      #           name  = "POSTGRES_PORT"
      #           value = "5432"
      #         }
      #         env {
      #           name  = "POSTGRES_DB"
      #           value = google_sql_database.default.name
      #         }
      #         env {
      #           name  = "POSTGRES_USER"
      #           value = google_sql_user.postgres_user.name
      #         }
      #         env {
      #           name  = "POSTGRES_PASSWORD"
      #           value = google_sql_user.postgres_user.password
      #         }
      #       }
    }

    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.postgres.connection_name
      }
    }
  }
}