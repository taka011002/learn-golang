locals {
  services_enabled = toset([
    "compute.googleapis.com",
    "artifactregistry.googleapis.com",
    "run.googleapis.com",
  ])
}