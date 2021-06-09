// WYMAGANA INSTALACJA BIBLIOTEKI "github.com/gocolly/colly"
// ZA POMOCĄ go get "github.com/gocolly/colly"

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"github.com/gocolly/colly"
)


// DEKLARACJA FUKNCJI ZWRACAJĄCEJ MAPE DLA PODANEGO SLICA W FORMACJE : UNIKALNA WARTOSC I LICZBA JEJ WYSTAPIEŃ
func printUniqueValue( arr []int) map[int]int {
    dict:= make(map[int]int)
    for _ , num :=  range arr {
        dict[num] = dict[num]+1
    }
    return dict
}

func main() {
	// WCZYTYWANIE DANYCH ZE STRONY INTERNETOWEJ

	// USTAWIENIE DOMEN KTORE MOZEMY ODWIEDZAC
	c := colly.NewCollector(
		colly.AllowedDomains("pl.wikipedia.org", "stackoverflow.com"),
	)


	// DEKLARACJA ZEMIENNYCH DLA DANYCH Z SIECI I ANALOGICZNIE DANYCH Z PLIKU
	var digits []int
	var digits2 []int  
	var numbers []string  
	var numbers2 []string  


	// WCZYTAENIE DANYCH Z SIECI
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// STWORZENIE REGEXA KTORY ZWRACA WLACZNIE LICZBY/CYFRY, RESZTE POMIJA 
		re := regexp.MustCompile("[0-9]+")
		numbers = re.FindAllString(e.Text, -1)[:]
	})
	
	// USTAWIANIE ODWIEDZANEJ WITRYNY
	c.Visit("https://pl.wikipedia.org/wiki/Mecz_Gwiazd_MLB")


	// PĘTLA KTORA ZWRACA Z ZEBRANYCH DANYCH TYLKO PIERWSZE CYFRY I ZAPISUJE JE DO SLICE, NIE LICZĄC '0'
	for _, num := range numbers{
		intNum, _ := strconv.Atoi(num) 
		firstDigit := intNum / int(math.Pow10(len(num) - 1))
		if firstDigit != 0{
			digits = append(digits, firstDigit)
		}
	}
	// SKORZYSTANIE Z WCZESNIEJ ZADEKLAROWANEJ FUNCKJI, ZMAPOWANIE CYFR Z ILOŚCIA ICH WYSTAPIEŃ
	mapValues := printUniqueValue(digits)

	// STWORZENIE PLIKU DLA DANYCH Z SIECI
	F, _ := os.Create("dataWEB.dat")

	// STORZENIE SLICE DLA KLUCZY MAPY
	keys := make([]int, 0)
	

	// ZAPIS KLUCZY DO SLICE
	for k := range mapValues {
    	keys = append(keys, k)
	}

	// POSORTOWANIE KLUCZY tj. CYFR OD 1 do 9
	sort.Ints(keys)

	// ZAPIS DO PLIKU DANYCH W FORMACIE : 'KLUCZ' 'WARTOSC'
	for _, k := range keys {
		F.WriteString(fmt.Sprintf("%d %d\n", k, mapValues[k]))
	
	}

	// WCZYTYWANIE DANYCH Z PLIKU 

	// OTWARCIE PLIKU Z DANYMI 
	file, err := ioutil.ReadFile("pliki.txt")

	// WYLAPANIE EWENTUALNEGO ERRORA
	if err != nil {
        fmt.Println("BRAK PLIKU")
		return
    }

	// SAPIS DANYCH Z PLIKU  DO ZMIENNEJ JAKO STRING
    string_file := string(file)

	// STWORZENIE REGEXA KTORY ZWRACA WLACZNIE LICZBY/CYFRY, RESZTE POMIJA 
	re2 := regexp.MustCompile("[0-9]+")
	numbers2 = re2.FindAllString(string_file, -1)[:]


	// PĘTLA KTORA ZWRACA Z ZEBRANYCH DANYCH TYLKO PIERWSZE CYFRY I ZAPISUJE JE DO SLICE, NIE LICZĄC '0'
	for _, num := range numbers2{
		intNum, _ := strconv.Atoi(num) 
		firstDigit := intNum / int(math.Pow10(len(num) - 1))
		if firstDigit != 0{
			digits2 = append(digits2, firstDigit)
		}
	}

	// SKORZYSTANIE Z WCZESNIEJ ZADEKLAROWANEJ FUNCKJI, ZMAPOWANIE CYFR Z ILOŚCIA ICH WYSTAPIEŃ
	mapValues2 := printUniqueValue(digits2)

	// STWORZENIE PLIKU DLA DANYCH Z PLIKU
	F2, _ := os.Create("dataPLIK.dat")

	// STORZENIE SLICE DLA KLUCZY MAPY
	keys2 := make([]int, 0)

	// ZAPIS KLUCZY DO SLICE
	for k := range mapValues {
    	keys2 = append(keys2, k)
	}

	// POSORTOWANIE KLUCZY tj. CYFR OD 1 do 9
	sort.Ints(keys2)

	// ZAPIS DO PLIKU DANYCH W FORMACIE : 'KLUCZ' 'WARTOSC'
	for _, k := range keys2 {
		F2.WriteString(fmt.Sprintf("%d %d\n", k, mapValues2[k]))
	}
}