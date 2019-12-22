mkdir -p bin;
psql -c "DROP DATABASE IF EXISTS ${DBNAME}"
createdb ${DBNAME}
