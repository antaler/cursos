FROM postgres:10.3

COPY up.sql /docker-entrypoint-intdb.d/1.sql

CMD ["postgres"]
