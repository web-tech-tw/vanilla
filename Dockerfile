# Vanilla
#    ---
# Directory based, virtual hosing manager and provider with PHP.
#
# (c) 2021 SuperSonic(https://github.com/supersonictw)

FROM alpine:3.13

# Copy Deploy Files
COPY ./script/ /script
COPY ./body /var/www/workplace

# Setup Requirement
RUN sh /script/require.sh
# Setup Nginx Gateway
RUN bash /script/nginx.sh
# Configure
RUN bash /script/configure.sh

# Copy Default Configuration
## Supervisor Config
COPY ./configs/supervisord.conf /etc/supervisord.conf
## NGINX Config
COPY ./configs/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf

# Clean
RUN rm -rf /script

# Set Entrypoint
EXPOSE 80
CMD supervisord
