# ==========================================
# Stage 1: Build Stage (بناء التطبيق)
# ==========================================
FROM golang:1.26-alpine AS builder

# تثبيت الأدوات الأساسية المطلوبة لتحميل الحزم
RUN apk add --no-cache git ca-certificates tzdata

# تحديد مجلد العمل داخل حاوية البناء
WORKDIR /app

# نسخ ملفات الاعتماديات أولاً للاستفادة من التخزين المؤقت (Caching)
COPY go.mod go.sum ./
RUN go mod download

# نسخ باقي ملفات المشروع
COPY . .

COPY app.env .

# بناء التطبيق بشكل ستاتيكي (Static Binary) مع إزالة معلومات التصحيح لتقليل الحجم والأمان
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main main.go

# ==========================================
# Stage 2: Final Security Stage (النسخة النهائية الآمنة)
# ==========================================
FROM alpine:3.19

# تثبيت شهادات الأمان الأساسية والوقت
RUN apk add --no-cache ca-certificates tzdata

# إنشاء مجموعة ومستخدم غير متميز (Non-root user) لأسباب أمنية صارمة
RUN addgroup -g 10001 appgroup && \
    adduser -D -u 10001 -G appgroup appuser

WORKDIR /app

# نسخ الملف التنفيذي فقط من مرحلة البناء (Builder Stage)
COPY --from=builder /app/main .

# منح صلاحية الملكية للمستخدم الجديد على مجلد التطبيق
RUN chown -R appuser:appgroup /app

# التحويل إلى وضع المستخدم العادي (مهم جداً للأمان)
USER appuser

# فتح البورت الذي يعمل عليه سيرفر Gin
EXPOSE 8080

# تشغيل التطبيق
CMD ["./main"]