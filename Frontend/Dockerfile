# ============================================================
#  Frontend statique (HTML/CSS/JS) servi par Nginx
#  Nginx sert aussi de reverse-proxy vers l'API Go (/api/)
# ============================================================
FROM nginx:alpine

# On remplace la conf par defaut
RUN rm -rf /usr/share/nginx/html/*

# Le site statique
COPY . /usr/share/nginx/html

# Notre configuration Nginx (statique + proxy /api/)
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
