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

module "rds" {
  source                 = "./modules/rds/"
  vpc_id                 = module.vpc.vpc_id
  vpc_security_group_ids = [module.ecs.ecs_security_group_id]
  private_subnets        = module.vpc.private_subnets
  db_name                = "gopaper"
  db_username            = "adminuser"
  db_password            = "adminpassword"
}

module "ecs" {
  source               = "./modules/ecs/"
  vpc_id               = module.vpc.vpc_id
  private_subnets      = module.vpc.private_subnets
  alb_sg_id            = module.alb.alb_sg_id
  alb_target_group_arn = module.alb.alb_target_group_arn
  container_image      = "${module.ecr_repository.repository_url}:latest"
  environment_variables = {
    "DB_ADDR"             = "postgres://adminuser:adminpassword@${module.rds.db_instance_endpoint}/gopaper?sslmode=require",
    "CORS_ALLOWED_ORIGIN" = module.s3_static_site.s3_static_site_enpoint
  }

  # - DB_ADDR=postgres://admin:admin@db:5432/gopaper?sslmode=disable
}

