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

module "ecs" {
  source               = "./modules/ecs/"
  vpc_id               = module.vpc.vpc_id
  private_subnets      = module.vpc.private_subnets
  alb_sg_id            = module.alb.alb_sg_id
  alb_target_group_arn = module.alb.alb_target_group_arn
  container_image      = "${module.ecr_repository.repository_url}:latest"
  environment_variables = {
    "DB_ADDR"             = "todo",
    "CORS_ALLOWED_ORIGIN" = module.s3_static_site.s3_static_site_enpoint
  }
}

