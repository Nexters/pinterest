#!/bin/bash

# Step 1: Stop your services with Docker Compose
docker-compose down

# Step 2: Attempt to renew the certificate in dry run mode
if certbot renew --dry-run > /dev/null 2>&1
then
    echo "Dry run successful. Proceeding with actual renewal."
    certbot renew
else
    echo "Dry run failed. Not proceeding with actual renewal."
    exit 1
fi

# Step 3: Start your services again with Docker Compose
docker-compose up -d