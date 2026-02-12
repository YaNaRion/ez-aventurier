# Colors for output
GREEN := \033[0;32m
YELLOW := \033[1;33m
RED := \033[0;31m
NC := \033[0m # No Color

.PHONY: dev staging prod build-frontend build-backend logs clean help \
        v1.0.0 v2.0.0 rollback status restart

## 🚀 Development
dev:
	@echo "$(GREEN)Starting development environment...$(NC)"
	@export ENVIRONMENT=dev && \
	 docker-compose --profile dev up --build

dev-detach:
	@echo "$(GREEN)Starting development environment (detached)...$(NC)"
	@export ENVIRONMENT=dev && \
	 docker-compose --profile dev up -d --build

## 🌍 Production
prod:
	@echo "$(RED)Starting production environment...$(NC)"
	@git checkout main || true
	@export ENVIRONMENT=prod && \
	 docker-compose --profile prod up -d --build

## 📦 Version deployment
v1.0.0:
	@echo "$(GREEN)Deploying version 1.0.0...$(NC)"
	@git checkout v1.0.0
	@$(MAKE) prod

v2.0.0:
	@echo "$(GREEN)Deploying version 2.0.0...$(NC)"
	@git checkout v2.0.0
	@$(MAKE) prod

## ⏪ Rollback
rollback:
	@echo "$(YELLOW)Rolling back to previous version...$(NC)"
	@git checkout HEAD~1
	@$(MAKE) prod

## 🏗️ Build frontend
build-frontend:
	@echo "$(GREEN)Building Dioxus frontend...$(NC)"
	@cd client_rust && dx build --release

## 🏗️ Build backend
build-backend:
	@echo "$(GREEN)Building Go backend...$(NC)"
	@docker-compose build backend

## 📊 Status
status:
	@docker-compose ps

logs:
	@docker-compose logs -f

restart:
	@docker-compose restart

## 🧹 Clean
clean:
	@echo "$(YELLOW)Cleaning up...$(NC)"
	@docker-compose down -v
	@docker system prune -f
	@echo "$(GREEN)Clean complete!$(NC)"

clean-all:
	@echo "$(RED)Full cleanup...$(NC)"
	@docker-compose down -v
	@docker system prune -a -f
	@docker volume prune -f
	@echo "$(RED)Full cleanup complete!$(NC)"

## ❤️ Health check
health:
	@echo "$(GREEN)Checking services...$(NC)"
	@curl -s http://localhost/api/health | jq . || echo "API not responding"
	@curl -s -o /dev/null -w "Frontend: %{http_code}\n" http://localhost

## 📖 Help
help:
	@echo "Available commands:"
	@echo "  $(GREEN)make dev$(NC)        - Start development environment"
	@echo "  $(YELLOW)make staging$(NC)    - Start staging environment"
	@echo "  $(RED)make prod$(NC)        - Start production environment"
	@echo "  make v1.0.0     - Deploy version 1.0.0"
	@echo "  make rollback   - Rollback to previous version"
	@echo "  make build-frontend - Build Dioxus frontend"
	@echo "  make logs       - View logs"
	@echo "  make clean      - Clean up containers"
