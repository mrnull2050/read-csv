# User CSV Processor

این پروژه یک ابزار ساده برای خواندن فایل CSV شامل کاربران، پاک‌سازی داده‌ها (Trim، Lowercase و Validation ایمیل و نام) و ذخیره در دیتابیس SQLite است.

## ویژگی‌ها
- پاک‌سازی و استانداردسازی نام و ایمیل کاربران
- اعتبارسنجی ایمیل با regex
- ذخیره اطلاعات در SQLite (یا driver pure Go)
- استفاده از Context برای کنترل timeout و concurrency
- Pipeline ساده برای پردازش داده‌ها

## نصب و اجرا
1. Clone پروژه:

git clone <repo-url>
cd <repo-folder>
