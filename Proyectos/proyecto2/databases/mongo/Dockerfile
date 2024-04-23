FROM mongo:4.4

COPY ./init-db.js /docker-entrypoint-initdb.d/

RUN chmod +x /docker-entrypoint-initdb.d/init-db.js

EXPOSE 27017