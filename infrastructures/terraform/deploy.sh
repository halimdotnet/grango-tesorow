#!/bin/bash

set -e

ENVIRONMENT=$1
COMMAND=$2

if [ -z "$ENVIRONMENT" ] || [ -z "$COMMAND" ]; then
  echo "Usage: ./deploy.sh [production|staging|development] [plan|apply|destroy|init|validate|refresh]"
  echo ""
  echo "Examples:"
  echo "  ./deploy.sh production init"
  echo "  ./deploy.sh production validate"
  echo "  ./deploy.sh production plan"
  echo "  ./deploy.sh production apply"
  echo "  ./deploy.sh staging plan"
  echo "  ./deploy.sh development destroy"
  exit 1
fi

case $ENVIRONMENT in
  production|staging|development)
    ;;
  *)
    echo "Error: Unknown environment '$ENVIRONMENT'"
    echo "Valid environments: production, staging, development"
    exit 1
    ;;
esac

case $COMMAND in
  plan|apply|destroy|init|validate|refresh)
    ;;
  *)
    echo "Error: Unknown command '$COMMAND'"
    echo "Valid commands: plan, apply, destroy, init, validate, refresh"
    exit 1
    ;;
esac

cd "$(dirname "$0")/root"

echo "=========================================="
echo "Environment: $ENVIRONMENT"
echo "Command:     $COMMAND"
echo "=========================================="
echo ""

if [ "$COMMAND" = "init" ] || [ "$COMMAND" = "validate" ]; then
  terraform $COMMAND
else
  terraform $COMMAND -var-file=../environments/$ENVIRONMENT/terraform.tfvars
fi

echo ""
echo "=========================================="
echo "Completed: terraform $COMMAND for $ENVIRONMENT"
echo "=========================================="