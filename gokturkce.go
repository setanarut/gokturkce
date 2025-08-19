// gokturkce Türkçe -> Göktürkçe alfabe dönüştürme kütüphanesi
// Bu kütüphane, Türkçe metinleri Göktürk alfabesine dönüştürmek için
// çeşitli kurallar ve dönüşümler uygular. Kutuplu ünsüzler, hece damgaları gibi
// Çeviri kusursuz değildir, ancak temel dönüşümleri sağlar.
package gokturkce

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

const (
	kalınSesliler = "aıuo"
	inceSesliler  = "eiüö"
)

var (
	// Özel hece damgaları (ligatürler) ve kutuplu olmayan ünsüzler için değiştirici.
	// Uzun eşleşmeler (örn. "ak") daha kısa eşleşmelerden (örn. "k") önce listelenmelidir.
	// "ük" kombinasyonu buradan çıkarılmıştır, çünkü "türk" kelimesinde 'ü'nün ayrı yazılması ve 'k'nin özel ligatüre dönüşmesi istenmektedir.
	ünlüKombinasyonDeğiştirici = strings.NewReplacer(
		// Hece damgaları (öncelikli ve spesifik olanlar)
		"ng", "𐰭",
		"ny", "𐰪",
		"nc", "𐰨",
		"nç", "𐰨",
		"nd", "𐰦",
		"nt", "𐰦",
		"ld", "𐰡",
		"lt", "𐰡",
		"ab", "𐰉",
		"eb", "𐰋",
		"ad", "𐰑",
		"ed", "𐰓",
		"ag", "𐰍",
		"eg", "𐰏",
		"ak", "𐰴",
		"ek", "𐰚",
		"ok", "𐰸",
		"uk", "𐰸",
		"ök", "𐰜",
		"ük", "𐰜",
		"ık", "𐰶",
		"al", "𐰞",
		"el", "𐰠",
		"an", "𐰣",
		// "en", "𐰤",
		"ar", "𐰺",
		"er", "𐰼",
		"as", "𐰽",
		"es", "𐰾",
		"at", "𐱃",
		"et", "𐱅",
		"ay", "𐰖",
		"ey", "𐰘",
		"iç", "𐰱",
		"ic", "𐰱",
		"eç", "𐰱",
		"eç", "𐰱",
		// Diğer kutupsuz ünsüz öbekleri ve tek ünsüzler
		"ç", "𐰲",
		"m", "𐰢",
		"p", "𐰯",
		"ş", "𐱁",
		"z", "𐰔",
	)
	ünlüDeğiştirici = strings.NewReplacer(
		"a", "𐰀",
		"e", "𐰀",
		"ı", "𐰃",
		"i", "𐰃",
		"o", "𐰆",
		"u", "𐰆",
		"ö", "𐰇",
		"ü", "𐰇",
	)
	olmayanSeslerDeğiştirici = strings.NewReplacer(
		"â", "a",
		"c", "ç",
		"f", "p",
		"ğ", "g",
		"h", "k",
		"j", "ç",
		"v", "b",
	)
	noktalamaTemizle = strings.NewReplacer(
		",", "",
		".", "",
		":", "",
		"?", "",
		"!", "",
		"&", "",
	)
)

// kutupluÜnsüzKalınMı sonraki veya en yakın ünlüye göre kalınlık döndürür
func kutupluÜnsüzKalınMı(idx int, liste []rune) bool {
	// Tüm karakterleri tek döngüde kontrol et
	for i := 1; i < len(liste); i++ {
		// Sağa bak
		if right := idx + i; right < len(liste) {
			if c := liste[right]; strings.ContainsRune(kalınSesliler, c) {
				return true
			} else if strings.ContainsRune(inceSesliler, c) {
				return false
			}
		}
		// Sola bak
		if left := idx - i; left >= 0 {
			if c := liste[left]; strings.ContainsRune(kalınSesliler, c) {
				return true
			} else if strings.ContainsRune(inceSesliler, c) {
				return false
			}
		}
	}
	return false
}

// Harf ve hece dönüşümlerini uygular (kutuplu ünsüzler)
func kutupluÜnsüzler(sözcük string) string {
	// Bu fonksiyon, sözcüğün hala Latin sesli harfler içerdiği varsayımıyla çalışır,
	// çünkü kalınlık/incelik belirlemesi için bu Latin sesli harfler gereklidir.
	// NOT: Buraya gelen sözcük, hece damgaları tarafından işlenmiş olabilir (örn. "ok" -> "𐰸").
	// İşlenmiş kısımlar artık Latin harfi içermediği için, kutuplu ünsüz tespiti
	// sadece kalan Latin harfleri üzerinde çalışacaktır.
	liste := []rune(sözcük)
	yListe := make([]rune, 0, len(liste))
	// n := len(liste) // Listeyi dönerken uzunluğa ihtiyacımız var

	for i, char := range liste {
		switch char {
		case 'b':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰉')
			} else {
				yListe = append(yListe, '𐰋')
			}
		case 'd':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰑')
			} else {
				yListe = append(yListe, '𐰓')
			}
		case 'g':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰍')
			} else {
				yListe = append(yListe, '𐰏')
			}
		case 'k':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰴')
			} else {
				yListe = append(yListe, '𐰚')
			}
		case 'l':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰞')
			} else {
				yListe = append(yListe, '𐰠')
			}
		case 'n':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰣')
			} else {
				yListe = append(yListe, '𐰤')
			}
		case 'r':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰺')
			} else {
				yListe = append(yListe, '𐰼')
			}
		case 's':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰽')
			} else {
				yListe = append(yListe, '𐰾')
			}
		case 't':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐱃')
			} else {
				yListe = append(yListe, '𐱅')
			}
		case 'y':
			if kutupluÜnsüzKalınMı(i, liste) {
				yListe = append(yListe, '𐰖')
			} else {
				yListe = append(yListe, '𐰘')
			}
		default:
			yListe = append(yListe, char)
		}
	}

	return string(yListe)
}

func ünlüDüşür(sözcük string) string {
	var builder strings.Builder
	runes := []rune(sözcük)
	n := len(runes)

	latinÜnlüMü := func(r rune) bool {
		return strings.ContainsRune("aeiouıöü", r)
	}

	// "baba" gibi kelimelerdeki ikinci hecedeki ünlüyü düşürme kuralı için bayrak.
	// Bu, C V C V -> C C V (örn. baba -> bba) yapısını hedefler.
	babaSyncopeApplied := false

	for i := 0; i < n; i++ {
		harf := runes[i]

		if !latinÜnlüMü(harf) {
			builder.WriteRune(harf)
		} else {
			// Bu bir Latin sesli harfidir. Ünlü düşürme kurallarını uygula.

			// Kural 1: 'ü' ve 'ö' sesli harfleri hiçbir koşulda düşürülmez.
			// Bu, "ötürü" ve "türk" gibi kelimelerin doğru yazılmasını sağlar.
			if harf == 'ü' || harf == 'ö' {
				builder.WriteRune(harf)
				continue
			}

			// Kural 2: "baba" gibi kelimelerde ikinci hecedeki 'a', 'ı', 'o', 'u' ünlülerinin düşürülmesi.
			// Şartlar:
			// 1. Kural daha önce uygulanmamış olmalı.
			// 2. Mevcut karakter bir Latin sesli olmalı ve 'a', 'ı', 'o', 'u' olmalı.
			// 3. Kelime en az 4 harfli olmalı (CVCV kalıbı için).
			// 4. Mevcut sesliyi önce ve sonra birer sessiz harf takip etmeli, ardından bir başka sesli harf gelmeli (C V C V kalıbı).
			//    Örneğin "baba" -> b A b A. İkinci 'A' hedeflenir.
			if !babaSyncopeApplied && i > 0 && i < n-1 && // Not ilk veya son karakter, yeterli uzunlukta
				!latinÜnlüMü(runes[i-1]) && // Önceki karakter sessiz olmalı
				strings.ContainsRune("aıou", harf) && // Bu ünlü düşürülecek türden (a,ı,o,u)
				i+1 < n && !latinÜnlüMü(runes[i+1]) && // Sonraki karakter sessiz olmalı
				i+2 < n && latinÜnlüMü(runes[i+2]) { // Ondan sonraki karakter sesli olmalı (CVCV kalıbı)

				babaSyncopeApplied = true // Kuralı bir kez uygulandı olarak işaretle
				continue                  // Bu ünlüyü düşür (yazma)
			}

			// Kural 3: Genel orta ünlü düşürme (ör: "asalaklar"daki ara 'a'lar).
			// Eğer harf 'a', 'e', 'ı', 'o', 'u' ise ve kelimenin ortasında sessiz harfler arasında ise düşür.
			// Bu kural, 'ü' ve 'ö' tarafından zaten korunmuş ünlülere uygulanmaz.
			// Ayrıca, bu kural "baba" gibi spesifik bir duruma takılmayan diğer orta ünlüleri düşürmek içindir.
			if i > 0 && i < n-1 && // Kelimenin başında veya sonunda değil
				!latinÜnlüMü(runes[i-1]) && // Önceki karakter sessiz
				strings.ContainsRune("aeıou", harf) && // Düşürülebilecek bir ünlü
				!latinÜnlüMü(runes[i+1]) { // Sonraki karakter sessiz
				continue // Bu ünlüyü düşür
			}

			// Kural 4: Sondan bir önceki ünlünün düşürülmesi (3 harften uzun kelimelerde)
			// Bu kural, yukarıdaki spesifik kurallar uygulanmadıysa veya farklı bir koşulu temsil ediyorsa geçerli olur.
			// 'ü' ve 'ö' zaten 1. kural tarafından korunmuştur.
			if i == n-1 {
				// Son harf her zaman yazılır
				builder.WriteRune(harf)
			} else if i == n-2 && n > 3 {
				// Sondan bir önceki harf düşer (3 harften uzun kelimelerde)
				continue
			} else {
				// Diğer ünlüler yazılır (eğer üstteki kurallardan birine uymuyorsa)
				builder.WriteRune(harf)
			}
		}
	}
	return builder.String()
}

func tersYaz(s, t string) {
	r := []rune(s)
	slices.Reverse(r)
	fmt.Println(string(r) + " - " + t)
}

// TersÇevir verilen stringi tersine çevirir
// Bu fonksiyon, verilen stringin karakterlerini ters sırada döndürür.
func TersÇevir(s string) string {
	r := []rune(s)
	slices.Reverse(r)
	return string(r)
}

func TR2GTR(s string, verbose bool) string {
	s = strings.ToLowerSpecial(unicode.TurkishCase, s)
	s = noktalamaTemizle.Replace(s)
	sözcükler := strings.Fields(s)
	sonuç := make([]string, 0, len(sözcükler))

	for _, sözcük := range sözcükler {
		temp := sözcük

		// Adım 1: Türkçe'de olmayan sesleri yaklaşık Türk seslerine dönüştür (Latin harfleri üzerinde).
		sözcük = olmayanSeslerDeğiştirici.Replace(sözcük)

		// Adım 2: Özel hece damgaları ve kutupsuz ünsüz öbeklerini dönüştür.
		// Bu adımda 'ak', 'ok', 'as' gibi hece ligatürleri ve 'ng' gibi ünsüz öbekleri işlenir.
		sözcük = ünlüKombinasyonDeğiştirici.Replace(sözcük)

		// Adım 3: Kutuplu ünsüzleri (sesli harf kalınlığına göre değişenleri) dönüştür.
		// Bu adım, Latin sesli harfler hala mevcutken yapılmalıdır, çünkü kalınlık/incelik
		// bilgisi bu Latin sesli harflerden alınır.
		// Bu aşamada kelime sonundaki 'k' için özel '𐰜' dönüşümü yapılır ("türk" için).
		sözcük = kutupluÜnsüzler(sözcük)

		// Adım 4: Ünlü düşürme kurallarını uygula.
		// Bu adım, Latin sesli harfleri Göktürk damgalarına dönüştürülmeden önce yapılmalıdır.
		// "baba" ve "asalaklar" gibi kelimelerdeki spesifik iç ünlü düşürme burada işlerken, "ötürü" ve "türk"deki ünlüler korunur.
		sözcük = ünlüDüşür(sözcük)

		// Adım 5: Kalan Latin sesli harfleri Göktürk sesli damgalarına dönüştür.
		// Bu, ünlü dönüşümünün son adımıdır.
		sözcük = ünlüDeğiştirici.Replace(sözcük)

		sonuç = append(sonuç, sözcük)

		if verbose {
			// tersYaz(sözcük, temp)
			fmt.Println(sözcük + " - " + temp)
		}
	}

	return strings.Join(sonuç, ":")
}
