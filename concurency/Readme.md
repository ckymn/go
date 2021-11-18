Concurency ( eszamanlilik) :

    * Fonksiyonlarin sirayla degil ayni anda calismasi sistemi 
    * Go da concurency ler "go routines" ve "channels" kullanilir

    Go Routines :

        * Bagimsiz ve es zamanli calisan ufak is parcaciklari 
        * Bunu yapmak icin sadece func basina "go" getirmelisiniz
        * Go routines birbirleri ile iletisimde degillerdir. 
        * Calismak icin baska fonksiyonlari beklemezler.
    
    Go Channels :

        * Go routine'ler partiktir ama birbirleri arasinda sinyal kurmadan calisirlar.
        * Sessizce calisir isleri bittikleri zaman kullaniklari alanlari geri iade ederler
        * Cannels'lar sayesinde ise "go routine"ler birbirileri ile iletişim kurabilir ve birbirlerine sinyaller  göndererek senkronize çalışabilirler.
        * Go routine'ler herhangi bir channel'a veri gonderip alabilir bunu yaparken verinin akis yonunu 
        gosteren ( name <- "merhaba" ) kullanilir.  
        * Bir Channel go routine'ler tarafindan kullanilmaya baslamdan once tanimlanmali ( name := make(chan string)