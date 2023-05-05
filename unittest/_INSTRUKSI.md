Umum:
1. Clone repo ini dan buat branch sesuai dengan nama masing2
2. kerjakan pada repo Masing-masing
3. Commit dan push perubahan hasil pengerjaan
4. Buat PR ke bracnh master
5. Batas akhir pengumpulan: 6 Mei 2023, Pukul 13.00 WIB


Homework:
1. Buat unit test untuk fungsi yang sudah di implementasikan logic2nya.
2. Kerjakan sesuai dengan CAMP ID masing-masing.
3. Pada file campIDsiswa.go (sesuaikan dgn campID masing2), uncomment semua code
4. Buat unit test sehingga semua code yang telah di uncomment tadi tercover semuanya
5. Untuk melihat coverage per-fungsi, bisa menggunakan command:

    `go test ./... -coverprofile coverage.out && go tool cover -func coverage.out`
    
    hasilnya akan menunjukkan coverage per fungsi, seperti contoh berikut:
    ```bash
    ok      unittest        0.001s  coverage: 47.1% of statements
    unittest/BE4284458.go:3:        HomeworkBE4284458       0.0%
    unittest/converter.go:9:        ToMeter                 100.0%
    unittest/converter.go:27:       main                    0.0%
    total:                          (statements)            47.1%
    ```

Note:
- tidak boleh ada perubahan logic pada file dimana implementasi fungsinya berada
- tidak boleh meng-uncomment file milik siswa lain.
- penambahan file baru hanya file untuk unit test saja


Penilaian: 
suplemen ke nilai keaktifan:
rumus: (coverage/100) * 6 
