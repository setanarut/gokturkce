# gokturkce
gokturkce Türkçe -> Göktürkçe alfabe dönüştürme kütüphanesi

## Örnek Kullanım

```go
s := "Biz, ilhamlarımızı gökten ve gâipten değil, doğrudan doğruya hayattan almış bulunuyoruz"
fmt.Println(gokturkce.TR2GTR(s, false))
// 𐰋𐰃𐰔:𐰃𐰠𐰴𐰢𐰞𐰺𐰢𐰔𐰃:𐰏𐰜𐱅𐰤:𐰋𐰀:𐰍𐰀𐰃𐰯𐱅𐰤:𐰓𐰏𐰠:𐰑𐰍𐰺𐰑𐰣:𐰑𐰍𐰺𐰖𐰀:𐰚𐰖𐱃𐱅𐰣:𐰞𐰢𐱁:𐰉𐰞𐰣𐰖𐰺𐰔
```