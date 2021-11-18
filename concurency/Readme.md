Concurency ( eszamanlilik) :

    * Fonksiyonlarin sirayla degil ayni anda calismasi sistemi 
    * Go da concurency ler "go routines" ve "channels" kullanilir

    Go Routines :
        * Goroutine çalışma modeline göre belleğin { stack } adı verilen bölgesinde başlangıç için "2Kb" kadar yer ayrıldığı ve bu alanın gerektiğinde büyüdüğü ifade ediliyor
        * Bagimsiz ve es zamanli calisan ufak is parcaciklari 
        * Bunu yapmak icin sadece func basina "go" getirmelisiniz
        * Go routines birbirleri ile iletisimde degillerdir. 
        * Calismak icin baska fonksiyonlari beklemezler.

        - Race Condiiton : 
            * Race Condition birden çok thread'in paylaşılan bir hafıza alanına aynı anda ulaşıp o alanı değiştirmesiyle meydana gelir.
            * Farkli "goroutines'lerin ayni memory alanina erismesidir.

            * !Burdaki race-condition hatasini [ "Mutex ve "" ] ile giderilir.

    
    Go Channels :

        * Go routine'ler partiktir ama birbirleri arasinda sinyal kurmadan calisirlar.
        * Sessizce calisir isleri bittikleri zaman kullaniklari alanlari geri iade ederler
        * Cannels'lar sayesinde ise "go routine"ler birbirileri ile iletişim kurabilir ve birbirlerine sinyaller  göndererek senkronize çalışabilirler.
        * Go routine'ler herhangi bir channel'a veri gonderip alabilir bunu yaparken verinin akis yonunu 
        gosteren ( name <- "merhaba" ) kullanilir.  
        * Bir Channel go routine'ler tarafindan kullanilmaya baslamdan once tanimlanmali ( name := make(chan string)
        * Channel'lar okundugu zaman bloklanir ve yazildigi zamanda bloklanir. 

        * UnBuffered : make(chan string)
        * Buffered : make(chan string, 100)
            - Birden çok veri almaktadır ancak standart bloklama konusunda herhangi bir değişiklik yoktur. Kanala veri girişi ve çıkışı esnasında bloklama işlemi gerçekleşir.
            - Kanala veri girişi esnasında kanalın tamamı dolunca yalnızca yeni veri girişi bloklanır. Kanal boşsa yalnızca okuma yani veri çıkışı işlemi bloklanır.
            - FIFO (First-In-First-Out) yaklaşımı burada geçerlidir yani kanala ilk giren veri, okuma tarafında ilk çıkan veridir.

        Not : Bir channel hem "yazma hem okuma" isleminde bloklama yapar.
              eger okuma isleminde bloklama istemiyorsak "Select/Case" mekanizmasi kullanilir.

        * Select Case :
            - bir channel'a veri girisini bekletebiliyoruz.

    Deadlock(kitlenme) :

        * unbuffered bir channel'a hic dinleyici olmayinca write yapmaya calisirsask deadlock olur !

        * numerous wait group add() gelirse deadlock olur

        * ayni lock iki kere ust uste cagirirsak deadlock olur.

        * selec-case 'lerde default veya if condition kullanilmaz ise deadlock olur