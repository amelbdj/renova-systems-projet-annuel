# ============================================================
#  Backend Go - build multi-stage (image finale legere)
# ============================================================

# ---- Etape 1 : compilation ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

# On telecharge d'abord les dependances (cache Docker efficace)
COPY go.mod go.sum ./
RUN go mod download

# Puis le reste du code
COPY . .

# On garantit l'existence des dossiers runtime (evite une COPY qui echoue plus bas)
RUN mkdir -p uploads static documents/doubles

# Binaire statique (pas de CGO) -> tourne sur une image minimale
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# ---- Etape 2 : image finale ----
FROM alpine:3.20

# Certificats (HTTPS Stripe / OneSignal) + fuseau horaire
RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -u 1000 appuser

WORKDIR /app

# On copie uniquement le binaire + les dossiers servis a l'execution
COPY --from=builder /app/server ./server
COPY --from=builder /app/uploads ./uploads
COPY --from=builder /app/static ./static

# On s'assure que tous les dossiers existent (PDF generes, etc.) + droits
RUN mkdir -p uploads static documents/doubles && chown -R appuser:appuser /app

USER appuser

# Le serveur Go ecoute sur 8081
EXPOSE 8081

CMD ["./server"]
