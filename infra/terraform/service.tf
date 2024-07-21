resource "google_project_service" "services_enabled" {
  for_each = local.services_enabled
  project  = var.project
  service  = each.value
}