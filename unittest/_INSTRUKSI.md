Umum:
1. Fork repo ini ke github masing2
2. kerjakan pada repo Masing-masing
3. Buat PR ke repo sumber (repo ini), arahkan ke branch master
4. Batas akhir pengumpulan: 6 Mei 2023, Pukul 13.00 WIB


Homework:
1. Buat unit test untuk fungsi yang sudah di implementasikan logic2nya.
2. Kerjakan sesuai dengan CAMP ID masing-masing.
3. Pada file campIDsiswa.go (sesuaikan dgn campID masing2), uncomment semua code
4. Buat unit test sehingga semua code yang telah di uncomment tadi tercover semuanya
5. Untuk melihat coverage per-fungsi, bisa menggunakan command `go test ./... -coverprofile coverage.out && go tool cover -func coverage.out`.

Note:
- tidak boleh ada perubahan logic pada file dimana implementasi fungsinya berada
- tidak boleh meng-uncomment file milik siswa lain.
- penambahan file baru hanya file untuk unit test saja


Penilaian: 
suplemen ke nilai keaktifan:
rumus: (coverage/100) * 6 