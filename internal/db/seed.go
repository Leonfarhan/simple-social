package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"math/rand"

	"github.com/Leonfarhan/simple-social/internal/store"
)

var usernames = []string{
	"aditya", "bima", "citra", "dewi", "eka", "farhan", "gita", "hani",
	"indra", "joko", "kirana", "luthfi", "maira", "nanda", "okta", "putri",
	"qori", "raka", "salsa", "tio", "ulfa", "vito", "wulan", "yoga",
	"zara", "bayu", "cahya", "dimas", "elisa", "fajar", "gina", "haris",
	"intan", "jihan", "kemal", "laras", "miko", "nabila", "osman", "prabu",
	"ratna", "satria", "tasya", "umar", "vania", "wahyu", "yudha", "zahra",
}

var titles = []string{
	"Catatan Pagi di Kedai Kopi", "Tips Menata Meja Kerja Kecil", "Belajar Konsisten Berolahraga",
	"Rute Jalan Kaki Favorit", "Resep Bekal Praktis", "Cara Mengatur Waktu Belajar",
	"Pengalaman Naik Kereta Malam", "Ide Liburan Hemat di Kota Sendiri", "Merawat Tanaman di Balkon",
	"Playlist Fokus untuk Bekerja", "Membaca Buku Sebelum Tidur", "Kebiasaan Finansial Sehat",
	"Menulis Jurnal Tanpa Ribet", "Mencoba Hobi Baru Akhir Pekan", "Menjaga Koneksi dengan Teman",
	"Membersihkan Rumah dalam 30 Menit", "Belajar Memasak dari Nol", "Mengurangi Distraksi Digital",
	"Menikmati Hujan dari Teras", "Rencana Kecil untuk Minggu Ini",
}

var contents = []string{
	"Pagi ini dimulai pelan dengan kopi hangat, buku catatan, dan rencana sederhana untuk menjalani hari.",
	"Meja kerja kecil tetap bisa nyaman kalau barang yang sering dipakai mudah dijangkau dan sisanya disimpan rapi.",
	"Olahraga terasa lebih ringan saat targetnya realistis, durasinya pendek, dan dilakukan pada jam yang sama.",
	"Jalan kaki di sekitar kompleks memberi jeda dari layar sekaligus membantu menemukan sudut kota yang sering terlewat.",
	"Bekal praktis tidak perlu rumit; nasi, lauk sederhana, dan sayur cepat tumis sudah cukup untuk hari sibuk.",
	"Waktu belajar lebih terkendali saat materi dipecah menjadi bagian kecil dan diberi jeda istirahat yang jelas.",
	"Perjalanan kereta malam selalu punya cerita sendiri, dari suara rel sampai obrolan singkat dengan penumpang lain.",
	"Liburan hemat bisa dimulai dari museum, taman kota, atau kedai kecil yang belum pernah dikunjungi sebelumnya.",
	"Tanaman balkon butuh cahaya cukup, jadwal siram yang konsisten, dan pot dengan drainase yang tidak tersumbat.",
	"Playlist yang tepat membantu masuk ke mode fokus tanpa harus memaksa diri bekerja terlalu keras di awal.",
	"Membaca beberapa halaman sebelum tidur menjadi cara sederhana untuk menutup hari tanpa terburu-buru.",
	"Mencatat pemasukan dan pengeluaran kecil membuat keputusan finansial harian lebih sadar dan tidak asal lewat.",
	"Jurnal tidak harus panjang; tiga kalimat tentang kejadian, perasaan, dan rencana besok sudah bisa membantu.",
	"Akhir pekan ini cocok dipakai mencoba hobi baru tanpa target harus mahir sejak percobaan pertama.",
	"Menjaga pertemanan kadang cukup dengan pesan singkat, ajakan makan, atau menanyakan kabar dengan tulus.",
	"Rumah terasa lebih lega setelah fokus membersihkan satu area kecil selama tiga puluh menit tanpa distraksi.",
	"Memasak dari nol dimulai dari resep sederhana, bahan yang mudah ditemukan, dan keberanian untuk gagal sedikit.",
	"Mengurangi distraksi digital bisa dimulai dengan mematikan notifikasi yang tidak penting selama jam produktif.",
	"Hujan sore di teras mengingatkan bahwa tidak semua waktu harus diisi dengan rencana besar.",
	"Minggu ini cukup dimulai dengan tiga prioritas kecil yang jelas, terukur, dan benar-benar mungkin dikerjakan.",
}

var tags = []string{
	"Keseharian", "Produktivitas", "Kesehatan", "Perjalanan", "Masakan",
	"Belajar", "Liburan", "Tanaman", "Musik", "Buku",
	"Keuangan", "Jurnal", "Hobi", "Pertemanan", "Rumah",
	"Memasak", "Digital", "Refleksi", "Rencana", "Kopi",
}

var comments = []string{
	"Relate banget, terima kasih sudah berbagi.",
	"Bagian ini paling kena buatku.",
	"Tulisannya ringan tapi tetap bermanfaat.",
	"Aku mau coba tips ini minggu ini.",
	"Setuju, kadang yang sederhana justru paling membantu.",
	"Pengalamanmu menarik, jadi kepikiran hal yang sama.",
	"Terima kasih, ini pas dengan yang sedang aku cari.",
	"Suka cara kamu menjelaskannya.",
	"Catatan kecil begini enak dibaca.",
	"Semoga bisa konsisten juga setelah baca ini.",
}

func Seed(store store.Storage, _ *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
