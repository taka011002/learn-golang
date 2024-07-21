resource "google_artifact_registry_repository" "learn-golang" {
  repository_id = "learn-golang"
  format = "DOCKER"

  location = var.region
}