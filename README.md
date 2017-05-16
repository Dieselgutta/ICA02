# ICA02


*Deltakere: Simen Fredriksen, Stian Blankenberg, Jone Manneråk, Kristian Hagberg, Tarjei Taxerås og Rasmus Sørby*


# Oppgave 1


![Tabell](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554656_10158500507685411_1574123725_n.png?oh=728f78c15a60cee0d61f2ddcf106fc6c&oe=591D7804)


Det er ikke mulig å finne nøyaktig antall siden det varierer hele tiden. Det meste som skjer på en pc krever en egen prosess, selv tastetrykk.
Prosessene som ikke kjører kommer opp som suspendert, dette betyr at de er lastet inn i hovedminnet, men ikke bruker noe prosessorkraft. De kan derfra enkelt endre status til kjørende. Er brukt for prosesser som “venter” på å bli startet.
Hos oss var det nettlesere som stod for hovedbruket av minne. Google Chrome er her godt kjent for å kreve store mengder minne, en av grunnene til dette er at den splitter det meste nettleseren består av i mindre deler. Dette forhindrer at resten av nettleseren crasher hvis en plug-in gjør det, men fører igjen til duplikater av prosesser. Det er mye informasjon som blir midlertidig lagret i minnet når man har en nettleser oppe, spesielt ved bruk av flere faner, her skjer det igjen duplisering av prosesser.
For at en prosess skal kjøre kreves det en prosessor, hovedminne, input-enhet og output-enhet. Prosessoren utfører kalkulasjonene som kreves for at et program skal kjøre. Det prosessoren kalkulerer blir underveis lagret i hovedminnet. Prosessoren kan igjen hente ting som er lagret i hovedminnet for å utføre videre kalkulasjoner. For å utføre prosesser kreves det en input-enhet som gir instruksjoner til prosessoren, dette er typisk fysiske ting som tastatur og datamus. Output-enheten er slik man mottar informasjonen som prosessoren produserer, dette kan være andre fysiske ting som en skjerm, en printer, eller høyttalere.


# Oppgave 2


For å skrive koden fant vi først ut hvor i ubuntu vi ville lage koden. Etter å ha valgt en mappe, skrev vi inn “nano hello.go” for å lage et hello.go-program. 


![Bilde2](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516180_10158500557310411_1191364368_n.png?oh=70151a0c10f8dbdbdd2c366e51781ded&oe=591D6232)

Dette er et bilde av koden. 

I dette eksempelet ville vi lage kode for Windows. Dette ble gjort ved å spesifisere OS=windows (GOOS=windows) og GOARCH=amd64. Dette er spesifikasjonene for en av 


PC’ene og kan variere for andre plattformer. Dette blir altså til en .exe-fil for windows. 
![Bilde3](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18578670_10158500561580411_981292444_n.png?oh=dc4d3db169aee318b78179c327bcaa11&oe=591D18F0)

Her ser vi kommandoen for å gjøre om filen til en exe-fil og en “ls”-kommando for å bekrefte at filen ligger der. Filen vi bruker er hello.exe.
Etter dette lastet vi opp koden via github. Vi har ikke tatt bilder av dette ettersom vi regner med at folk kjenner til den vanlige måten å laste opp/ned kode eller prosjekter fra Github. Dette er kanskje ikke den optimale måten å gjøre det på, spesielt med tanke på at filen ble en del større enn vanlig kode, men det fungerte greit og det var få problemer med det. Etter at filen var lastet ned, slettet vi den, fordi det generelt sett ikke er god praksis å ha slike filtyper på GitHub. 
En bedre måte å overføre på er gjennom SCP. SCP står for Secure Copy Protocol, og er en god måte å overføre fra en remote maskin til en local. Dette kan gjøres ved å bruke terminalen slik: scp -i ~downloads/Pem-filnavn ubuntu@IP:/home/ubuntu/filnavn. For meg vil det da se slik ut: “scp -i ~downloads/RRS.pem ubuntu@IP:home/ubuntu/Hello.exe .” Da vil Hello.exe-filen bli lastet ned i den mappen jeg befinner meg i, i terminalen. 

![Bilde4](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554964_10158500569170411_1482575662_n.png?oh=d41fb6c27e3c353f3f2209ef54037949&oe=591CF938)


# Oppgave 3

![Bilde5](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18578768_10158500577740411_759467107_n.png?oh=2e9b02a3092f7a8ed1d9463d1adf2c00&oe=591D84A6)


Her er et eksempel på hvordan man kan skrive det i ATOM. Dette er da int32. 
![Bilde6](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18555190_10158500578085411_274414608_n.png?oh=ca132452166d3f821532ed3e2ca60ff0&oe=591D0BEB]


Det vi har gjort her er at vi har laget tester som produserer feil. Det vi får som feilmelding på int32 er "constant 2147483648 overflows". Det vil da si at den er ikke innenfor int32 sin rekkevidde.
Settet av alle signerte 32-bits heltall er kun innenfor dette rekkevidde (-2147483648 til 2147483647), men vi gikk utenfor denne rekkevidden, fordi da produserer den feil. Slik som i eksempelet over. 


![Bilde7](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18518709_10158500587665411_708529554_n.png?oh=bdbd9f6210165620f29b4f5b6706d05e&oe=591CFE8D)

Her har vi da tatt inn to parametere og skrevet de ut i Powershell. Da vil det bli slik som dette: 

![Bilde8](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554597_10158500594140411_1483660565_n.png?oh=ad45ea2f6877b51bd4c8c470a9ebaa9b&oe=591D4899)

Det som skjer er at vi legger inn to argumenter (parametere) slik at den skal skrive det ut når vi prøver å kjøre denne i powershell. Det som blir gjort er at filen blir kjørt sammen med to variabler (tall) slik powershell kan regne ut “regnestykke” og dermed printer svaret rett under. 1+2=3.


# Oppgave 4

![Bilde](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554465_10158500603740411_303104132_n.png?oh=0b01ed12a8d64a83115ad7abc61ea1de&oe=591D19ED)

Logaritmisk graf som viser de tre forskjellige benchmark-testene opp mot hverandre. 
Benchmark-testene går gjennom en sortering av henholdsvis 100, 1000 og 10000 tall. Det kommer her tydelig fram effektiviteten til QuickSort sammenlignet med BubbleSort; ved 1000 tall ligger QuickSort på rundt 1 000 000 nanosekunder, mens begge BSort testene går forbi 50 000 000 ns (BSortModified ender her opp på langt over 300 000 000 ns).
Sammenlignet med Big-O regnearket får vi bekreftet det vi så ovenfor; til tross for at BSort  har potensial til å være bedre enn QSort, er den vanligvis suboptimal. Det kommer også fram at BSortModified bare er optimal sammenlignet med BSort opp til et viss punkt. Dette er sammenlignbart med andre sorteringsalgoritmer hvor vanligvis optimal algoritme blir suboptimal sammenlignet med en annen etter et viss punkt.
Selv om vi bruker relativt få verdier i testene, kan vi konkludere med at sorteringsfunksjonene er gjennomsnittlige i forhold til Big-O basert på resultatet.


# Oppgave 5
Først kjørte vi main_boring.go via bash. Denne kjører kontinuerlig, og printer ut index +1 hvert sekund. Vi lastet deretter ned Process Explorer. Man kan herfra også enkelt avslutte prosessen ved å bruke ctrl-c kommandoen.

![Bilde9](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18518560_10158500603735411_1005184835_n.png?oh=7ac168e3452d2215420f5f81a70a92e1&oe=591D0A81)

 
Her kan man se en liste over kjørende prosesser, sammen med utfyllende informasjon om hvor mye CPU de bruker, og hvor mange threads de er knyttet opp med. Man kan også enkelt avslutte prosessen herfra ved å høyreklikke på den og velge “kill process”.
Vi fant bash på denne listen, og valgte videre den kjørende boring prosessen.


![Bilde10](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516081_10158500603745411_1723717097_n.png?oh=3d70656d373706210f827fb9c327bf85&oe=591D5B0A)


Her kan man se 5 forskjellige threads som er knyttet opp til boring prosessen. Trykker man inn på disse står det at samtlige prosesser er sleeping, de venter altså på nye ordre.


![Bilde11](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516311_10158500603750411_1961678353_n.png?oh=7730a923e0ab7908d43ea4550fb16c24&oe=591D7934)


Mens vi kjørte main_boring.go på den virtuelle serveren gjennom et vindu, kunne vi se alle prosessene gjennom et nytt vindu. Her brukte vi et tillegg som het htop for lettere visualisering av prosesser.

![Bilde12](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18492951_10158500603725411_738716107_n.png?oh=0882267aa327486b1c66bc3ea2330a9d&oe=591CED53)


Man kan også bruke kommandoen “ps aux | grep main_boring;” for å finne prosesser som inneholder main_boring i informasjonen. Her får vi opp 3 forskjellige prosesser.


![Bilde13](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554529_10158500619075411_1011320544_n.png?oh=b3a66cfdf5ad6c08a79f316bf87f991d&oe=591D17DE)

Ved å bruke PID’en til den øverste prosessen (27628) kunne vi finne hvor mange tråder som er tilknyttet denne prosessen. Her kan vi også se hvilken status trådene har ved å se hvilke symbol de har under “S” (status). Vi kan da se at samtlige av trådene er sleeping.

Her er resultatet ved undersøkelse av main_boring_goroutine.go:

![Bilde14](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18492602_10158500619025411_498333445_n.png?oh=971f833c4e4b78dbf9f4362644529b9f&oe=591D7F05)


![Bilde15](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18555270_10158500619060411_719204135_n.png?oh=bf12679e65b8a37b9d8546d38eaeebbc&oe=591D55E8)


![Bilde16](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516200_10158500619055411_1529747607_n.png?oh=6e8cc3dba1ef571d85dfbcb74676f56d&oe=591D7861)


Boring10 funksjonen vil nok være mer dominerende med å få designert prosessorkraft ettersom den skriver ut tall i et mye høyere tempo.
