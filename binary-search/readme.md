## Cara kerja Binary search :
Algoritma Binary search ini menggunakan metode devide and conquer dimana sebuah list akan dipecah menjadi 2 bagian dan kembali menentukan nilai tengah dan membandingkannya secara terus menerus, hingga ditemukan bahwa nilai tengahnya adalah angka yang dicari.

Untuk melakukan pembagian list tersebut, diperlukan sebuah angka tengah, bila angka yang dicari lebih kecil daripada angka ditengah list, maka algoritma ini akan fokus ke serangkaian index yang berada di sebelah kiri list, begitu juga sebaliknya.

Untuk menjalankan algoritma ini, kita harus mengurutkan terlebih dahulu value dari array yang ingin kita cari secara ascending. untuk masalah ini, kita bisa menggunakan algoritma sorting seperti bubble sort,insertion sort atau quick sort.
Penjelasan algoritma binary search
1. Pertama-tama diambil posisi awal 0 dan posisi akhir = N - 1, kemudian dicari posisi data tengah dengan rumus (posisi awal + posisi akhir) / 2. 
2. Kemudian data yang dicari dibandingkan dengan data tengah.Kemudian kita cari posisi data tengah dengan rumus posisi tengah yaitu = (posisi awal + posisi akhir ) div 2.

3. Lalu data yang di cari akan dibandingkan dengan data tengah
4. Jika sama, data ditemukan, Proses selesai.
5. Jika lebih kecil, maka proses akan dilakukan kembali tetapi, posisi akhir dianggap sama dengan posisi tengah -1.
6. Jika lebih besar pun proses akan dilakukan kembali tetapi posisi awal dianggap sama dengan posisi tengah +1.
7. Mengulang dari langkah kedua sampai data ditemukan, atau tidak ditemukan.
8. Searching biner ini akan berakhir jika data ditemukan posisi awal lebih besar dari pada posisi akhir. Jika posisi awal sudah lebih besar dari posisis akhir berarti data tidak ditemukan.