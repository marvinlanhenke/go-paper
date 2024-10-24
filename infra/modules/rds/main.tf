resource "aws_security_group" "rds_sg" {
  name        = "ml-sa-go-paper-production-rds-sg"
  description = "Security group for RDS PostgreSQL instance"
  vpc_id      = var.vpc_id

  ingress {
    description     = "Allow PostgreSQL from ECS service"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = var.vpc_security_group_ids
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "ml-sa-go-paper-${var.db_name}-rds-sg"
    Environment = var.environment
  }
}

resource "aws_secretsmanager_secret" "db_credentials" {
  name        = "ml-sa-go-paper-${var.db_name}-credentials"
  description = "Credentials for the ml-sa-go-paper PostgreSQL database"

  tags = {
    Name        = "ml-sa-go-paper-${var.db_name}-credentials"
    Environment = var.environment
  }
}

resource "aws_secretsmanager_secret_version" "db_credentials_version" {
  secret_id = aws_secretsmanager_secret.db_credentials.id
  secret_string = jsonencode({
    username = var.db_username
    password = var.db_password
  })
}

resource "aws_db_subnet_group" "this" {
  name       = "ml-sa-go-paper-${var.db_name}-subnet-group"
  subnet_ids = var.private_subnets

  tags = {
    Name        = "ml-sa-go-paper-${var.db_name}-subnet-group"
    Environment = var.environment
  }
}

resource "aws_db_instance" "this" {
  identifier                 = "ml-sa-go-paper-${var.db_name}"
  db_name                    = var.db_name
  username                   = var.db_username
  password                   = var.db_password
  engine                     = var.engine
  engine_version             = var.db_engine_version
  instance_class             = var.db_instance_class
  allocated_storage          = var.allocated_storage
  max_allocated_storage      = var.max_allocated_storage
  storage_type               = var.storage_type
  db_subnet_group_name       = aws_db_subnet_group.this.name
  vpc_security_group_ids     = [aws_security_group.rds_sg.id]
  multi_az                   = var.multi_az
  publicly_accessible        = false
  auto_minor_version_upgrade = true
  license_model              = "postgresql-license"
  maintenance_window         = "Mon:05:00-Mon:06:00"
  backup_retention_period    = var.backup_retention_period
  delete_automated_backups   = true
  final_snapshot_identifier  = "ml-sa-go-paper-${var.db_name}-final-snapshot"
  skip_final_snapshot        = false

  tags = {
    Name        = "ml-sa-go-paper-${var.db_name}-db-instance"
    Environment = var.environment
  }

}
