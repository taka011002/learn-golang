terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.38.0"
    }
  }
}

provider "google" {
  project = var.project
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-network"
}