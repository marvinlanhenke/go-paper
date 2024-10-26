resource "aws_security_group" "alb_sg" {
  description = "Security group for the ALB"
  name        = "ml-sa-go-paper-alb-sg"
  vpc_id      = var.vpc_id

  ingress {
    description = "Allow HTTP traffic from anywhere"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow HTTPS traffic from anywhere"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "ml-sa-go-paper-alb-sg"
    Environment = var.environment
  }
}

resource "aws_lb" "this" {
  name                       = "ml-sa-go-paper-alb"
  internal                   = false
  load_balancer_type         = "application"
  security_groups            = [aws_security_group.alb_sg.id]
  idle_timeout               = 60
  subnets                    = var.subnets
  enable_deletion_protection = false

  tags = {
    Name        = "ml-sa-go-paper-alb"
    Environment = var.environment
  }
}

resource "aws_lb_target_group" "this" {
  name        = "ml-sa-go-paper-tg"
  port        = 8080
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "ip"

  health_check {
    path                = "/v1/health"
    protocol            = "HTTP"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }

  tags = {
    Name        = "ml-sa-go-paper-tg"
    Environment = var.environment
  }
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.this.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this.arn
  }
}
