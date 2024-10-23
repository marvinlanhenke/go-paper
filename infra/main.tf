provider "aws" {
  region = var.region
}

module "vpc" {
  source   = "./modules/vpc/"
  vpc_cidr = "10.0.0.0/16"
}

module "alb" {
  source      = "./modules/alb/"
  environment = "Production"
  subnets     = module.vpc.public_subnets
  vpc_id      = module.vpc.vpc_id
}

module "s3_static_site" {
  source      = "./modules/s3_static_site/"
  bucket_name = "ml-sa-s3-static-site"
}

module "ecr_repository" {
  source          = "./modules/ecr_repository/"
  repository_name = "ml-sa-go-paper-backend"
}

