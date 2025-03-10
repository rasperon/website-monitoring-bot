Tabii, işte projeniz için eğlenceli ve açıklayıcı bir **README.md** dosyası:

---

# Website Monitoring 🚀

**Website Monitoring** aracı, izlediğiniz web sitelerinin durumunu kontrol etmek için basit bir Go uygulamasıdır. Web sitelerinizin erişilebilirliğini düzenli aralıklarla kontrol eder ve eğer bir hata varsa **Discord Webhook** ile bildirim gönderir.

---

## 📦 **Kurulum ve Başlangıç**

### 1. **GitHub'dan İndirme**
Öncelikle bu projeyi bilgisayarınıza klonlayın:

```bash
git clone https://github.com/kullaniciAdi/website-monitor.git
cd website-monitor
```

### 2. **Bağımlılıkları Yükleme**
Go modüllerini yüklemek için şu komutu çalıştırın:

```bash
go mod tidy
```

### 3. **Gerekli Yapılandırmalar**
Kullanıcıdan site URL'leri, kontrol süresi ve Discord Webhook URL'sini almak için aşağıdaki komutla uygulamayı başlatın:

```bash
go run main.go
```

- İzlemek istediğiniz web sitelerinin URL'lerini girin (boş bırakabilirsiniz).  
- Kontrol aralığını saniye cinsinden belirtin.  
- Discord Webhook URL'sini girin (zorunlu değil).

---

## 🔧 **Kullanım**

Uygulama, belirlediğiniz aralıklarla her bir siteyi kontrol eder ve aşağıdaki durumları kontrol eder:

- **Erişim Hatası (HTTP 4xx/5xx)**: Eğer siteye erişilemiyorsa veya bir hata alındıysa, Discord’a bildirim gönderilir.  
- **Başarıyla Erişim (HTTP 2xx)**: Site doğru şekilde çalışıyorsa, herhangi bir bildirim gönderilmez.

Her statü değişikliğinde Discord kanalınıza bildirim gönderilecektir. Aynı hata tekrar meydana gelirse, bir bildirim sadece ilk defa hata oluştuğunda yapılacaktır.

---

## 🔨 **Proje Yapısı**

```bash
website-monitor/
│── main.go               # Ana uygulama
│── config/
│   ├── config.go         # Kullanıcı girişlerinin alındığı dosya
│── monitor/
│   ├── check.go          # Web sitesi kontrol mekanizması
│── notify/
│   ├── discord.go        # Discord Webhook bildirimleri
│── logs/
│   ├── monitor.log       # Uygulama logları
│── utils/
│   ├── logger.go         # Log dosyasına yazma fonksiyonu
│   ├── storage.go        # Kullanıcı yapılandırmalarını kaydetme
```

---

## 🛠 **Özellikler**

- **Discord Bildirimleri**: Web sitenizin durumu değiştiğinde, Discord kanalınıza bildirim gönderir.
- **Retry Mekanizması**: Erişim hatası durumunda siteyi 3 kez tekrar deneme mekanizması.
- **Yapılandırma Kaydetme**: Kullanıcı girdiği ayarları JSON dosyasına kaydeder, böylece her seferinde tekrar girmenize gerek yok.
- **Loglama**: Web sitesi durumları ve hata mesajları `monitor.log` dosyasına kaydedilir.

---

## 📲 **Discord Webhook**

Web siteniz hakkında bildirim almak için Discord'da bir webhook oluşturmanız gerekecek. Webhook oluşturduktan sonra, URL'yi uygulamaya girmeniz yeterlidir.

[Discord Webhook Nasıl Oluşturulur?](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)

---

## 💬 **Katkı**

Katkı sağlamak isterseniz, bir Pull Request göndererek önerilerinizi iletebilirsiniz. Ya da sorunları **Issues** bölümünde bildirebilirsiniz.

---

## 📜 **Lisans**

MIT Lisansı altında lisanslanmıştır. (Lisans dosyasını inceleyin.)

---

🛠 **İyi kodlamalar!** 🚀