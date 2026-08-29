$ErrorActionPreference = 'Stop'

$database = 'agriapp'
$user = 'postgres'
$password = 'postgres'
$host = 'localhost'
$port = '5432'

Write-Host "Creating PostgreSQL database '$database' if it does not exist..."

$env:PGPASSWORD = $password

$createDbSql = @"
SELECT 'CREATE DATABASE $database'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$database')\gexec
"@

& psql -h $host -p $port -U $user -d postgres -v ON_ERROR_STOP=1 -c $createDbSql

Write-Host "Database ready: postgresql://$user:$password@$host:$port/$database?sslmode=disable"
Write-Host "Run the app with:"
Write-Host "  go run ."
Write-Host "Optional manual override:"
Write-Host "  `$env:AGRI_DB_URL = 'postgres://postgres:postgres@localhost:5432/agriapp?sslmode=disable'"
Write-Host "  go run ."
