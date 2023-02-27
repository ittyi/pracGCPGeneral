provider "google" {
  credentials = "${file("<your-credential-file-path>")}"
  project     = "${lookup(var.project_name, "${terraform.workspace}")}"
  region      = "asia-northeast1"
}