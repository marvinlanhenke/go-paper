resource "aws_security_group" "ecs_service_sg" {
  name        = "ml-sa-go-paper-${var.environment}-ecs-sg"
  description = "Security group for ECS Fargate service"
  vpc_id      = var.vpc_id

  ingress {
    description     = "Allow traffic from ALB on container port ${var.container_port}"
    from_port       = var.container_port
    to_port         = var.container_port
    protocol        = "tcp"
    security_groups = [var.alb_sg_id]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "ml-sa-go-paper-${var.environment}-ecs-sg"
    Environment = var.environment
  }
}

resource "aws_ecs_cluster" "this" {
  name = "ml-sa-go-paper-${var.environment}-cluster"

  tags = {
    Name        = "ml-sa-go-paper-${var.environment}-cluster"
    Environment = var.environment
  }
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name = "ml-sa-go-paper-${var.environment}-ecs-task-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      },
      Action = "sts:AssumeRole"
    }]
  })

  tags = {
    Name        = "ml-sa-go-paper-${var.environment}-ecs-task-execution-role"
    Environment = var.environment
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_task_definition" "this" {
  family                   = "ml-sa-go-paper-${var.environment}-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.cpu
  memory                   = var.memory

  execution_role_arn = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([{
    name  = "ml-sa-go-paper-container"
    image = var.container_image
    portMappings = [{
      containerPort = var.container_port
      hostPort      = var.container_port
      protocol      = "tcp"
    }]
    essential = true

    environment = [
      for key, value in var.environment_variables : {
        name  = key
        value = value
      }
    ]

    logConfiguration = {
      logDriver = "awslogs"
      options = {
        "awslogs-group"         = aws_cloudwatch_log_group.ecs_logs.name
        "awslogs-region"        = "eu-central-1"
        "awslogs-stream-prefix" = "ecs"
      }
    }
  }])

  tags = {
    Name        = "ml-sa-go-paper-${var.environment}-task"
    Environment = var.environment
  }
}

resource "aws_cloudwatch_log_group" "ecs_logs" {
  name              = "ecs/ml-sa-go-paper-${var.environment}-backend"
  retention_in_days = 30

  tags = {
    Name        = "ml-sa-go-paper-${var.environment}-ecs-logs"
    Environment = var.environment
  }
}

resource "aws_ecs_service" "this" {
  name            = "ml-sa-go-paper-${var.environment}-ecs-service"
  cluster         = aws_ecs_cluster.this.id
  task_definition = aws_ecs_task_definition.this.arn
  desired_count   = var.desired_count
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.private_subnets
    security_groups  = [aws_security_group.ecs_service_sg.id]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = var.alb_target_group_arn
    container_name   = "ml-sa-go-paper-container"
    container_port   = var.container_port
  }

  tags = {
    Name        = "ml-sa-go-paper-${var.environment}-ecs-service"
    Environment = var.environment
  }
}
