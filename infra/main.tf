provider "aws" {
  region = var.region
}

module "s3_static_site" {
  source = "./modules/s3_static_site"
}
