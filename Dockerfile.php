FROM php:8.1.6-fpm-alpine3.14

RUN docker-php-ext-install sockets

