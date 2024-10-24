provider "aws" {
  region = var.region
}

module "vpc" {
  source      = "./modules/vpc/"
  environment = var.environment
  vpc_cidr    = "10.0.0.0/16"
}

module "rds" {
  source             = "./modules/rds/"
  environment        = var.environment
  db_name            = var.db_name
  db_username        = var.db_username
  db_password        = var.db_password
  vpc_id             = module.vpc.vpc_id
  private_subnets    = module.vpc.private_subnets
  security_group_ids = [module.ecs.ecs_security_group_id]
}

module "s3" {
  source      = "./modules/s3/"
  bucket_name = "ml-sa-s3-static-site"
}

module "alb" {
  source      = "./modules/alb/"
  environment = var.environment
  subnets     = module.vpc.public_subnets
  vpc_id      = module.vpc.vpc_id
}

module "ecr" {
  source          = "./modules/ecr/"
  environment     = var.environment
  repository_name = "ml-sa-go-paper-backend"
}

module "ecs" {
  source               = "./modules/ecs/"
  region               = var.region
  environment          = var.environment
  vpc_id               = module.vpc.vpc_id
  private_subnets      = module.vpc.private_subnets
  alb_sg_id            = module.alb.alb_sg_id
  alb_target_group_arn = module.alb.alb_target_group_arn
  container_image      = "${module.ecr.repository_url}:latest"
  environment_variables = {
    "DB_ADDR"             = "postgres://${var.db_username}:${var.db_password}@${module.rds.db_instance_endpoint}/gopaper?sslmode=require",
    "CORS_ALLOWED_ORIGIN" = "http://${module.s3.s3_static_site_endpoint}"
  }
}

