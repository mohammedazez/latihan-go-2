# latihan array dan slice


## soal 1 : sum number

1. hitunglah hasil dari array dengan tipe data integer, seperti contoh berikut :
```
contoh :
1. [1, 2, 3, 4, 5] = 1 + 2 + 3 + 4 + 5 = 15 
2. [1, 2, 3] = 1 + 2 + 3 = 6 
3. [] = 0
```

## soal 2 : reverse data

lakukan konversi dari string menjadi sebuah array / slice yang berisi tiap karakter string , dan jika bertemu spasi maka jangan ditampilkan. Contohnya sebagai berikut:

```
contoh: 
1.afista => [a, t, s, i, f, a]
2 skilvul => [l, u, v, l, i, k, s]
3 impact byte => [e, t, y, b, t, c, a, p, m, i]
```

## soal 3 : sum max and min

buatlah sebuah eksekusi kode untuk melakukan penjumlahan dari tipe data array / slice yang berisi baris angka, dengan dua data berikut : 
1. angka terbesar ditambah angka terkecil 
2. angka terbesar dikurangi angka terkecil

tampilkan hasil penjumlahan dari nomer 1 dan nomer 2 tanpa perlu dimasukkan ke dalam array, atau ditampilkan biasa saja

```
contoh : 
1. [10, 120, 14,50, 5] => (120 + 5) + (120 - 5) = 125 + 115 = 240
2. [0, 2, 3, 4, 5] => (5 + 0) + (5 - 0) = 5 + 5 = 10
```

## soal 4 : matching product

Ada banyak tumpukan kaus kaki yang harus dicocokkan dengan warna. Diberikan serangkaian bilangan bulat yang mewakili warna setiap kaus kaki, tentukan berapa banyak pasang kaus kaki dengan warna yang serasi.

Ketika warna yang didapat serasi, maka kita sudah menganggap ada 1 produk yang sudah cocok.

```
contoh : 
1. [10, 10, 20, 20, 10, 30, 50, 10, 20] => terdapat 3 produk yang match yaitu 10 & 10 , 20 & 20, 10 & 10
output adalah 3

2. [1, 3, 3, 4, 5, 6, 1] => terdapat 2 produk yg match yaitu 1 & 1, 3 & 3
output adalah 2

3. [67, 213, 1, 0, 98, 2] => tidak ada produk yang match
output adalah 0
```


## soal 5 : slice in the middle 

buatlah sebuah tipe data slice yang menampung data dari sebuah array dengan output sebagai berikut :
1. jika jumlah data array adalah genap, maka tampung 2 data tengah ke dalam slice dan tampilkan
2 jika jumlah data array adalah ganjil, maka tampung data tengah kemudian 1 data sebelum tengah dan 1 data Setelah tengah, kemudian tampilkan

```
contoh : data := [5]string{a, i, u, e, o} => [i,u,e]
contoh : data := [4]string{b, a, c, a} => [a, c]
contoh : data := [2]string{j, o, m, b, l, o} => [m, b]
contoh : data := [5]string{k, e, r, e, n} => [e, r, e]
```


## jika sudah dikerjakan semua, silahkan kerjakan soal latihan javascript menggunakan bahasa pemrograman go
```
https://github.com/Excellent-Echo/latihan
```
