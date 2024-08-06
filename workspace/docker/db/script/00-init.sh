#!/bin/bash
# cbaciliod
set -e
# Ejecutar cada script SQL:

for script in /docker-entrypoint-initdb.d/*.sql; do
    echo "Running $script"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -f "$script"
done

echo "execute success"