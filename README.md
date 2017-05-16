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
