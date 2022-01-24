# ICT-DEV-IP-18 - Pinautomaat
Eindopdracht voor module ICT-DEV-IP-18

# Introductie

Dit is de basis voor de eindopdracht van het vak Imperatief Programmeren. Op deze basis ga je voortbouwen. 

In week 7, 8 en 9 werk je aan deze eindopdracht (goed voor 3 * 14 = 42 uur!). Zorg dat je op tijd begint met de eindopdracht. De opdrachten staan hieronder beschreven. Alle opdrachten werk je uit. Per opdracht is er een moeilijkheidsgraad aangegeven, zodat je voor jezelf kunt inschatten wanneer je met de opdracht aan de slag kunt. Met deze opdrachten demonstreer je aan ons jouw vaardigheden doordat hierin alle aspecten aan bod komen zoals beschreven in de toetsmatrijs, waarin alle beoordelingscriteria staan uitgewerkt.

Vóór de deadline (zie ItsLearning) lever je jouw code in via ItsLearning.  

> **LET OP: In week 7 ga je in het college beginnen met deze basis. Je loopt door de hele applicatie heen en voorziet deze van commentaar. Dit is essentieel om de applicatie goed te begrijpen en hierop voort te bouwen!**

Let op de volgende zaken:
- je wordt niet beoordeeld op de opmaak / styling van je applicatie (css)
- maak zo vaak mogelijk (tenminste na elke opdracht) een commit in Git of een backup / kopie van je programma. Zo kom je niet voor verrassingen te staan bij het inleveren, en kun je altijd terug naar een werkende versie

> Tijdens de module Frontend Development heb je al kennis gemaakt met de basis van Git.

# Usecase

Het programma dat je gaat ontwikkelen is een online pinautomaat. Met deze pinautomaat is het mogelijk om te kijken wat het saldo is op jouw rekening, en kun je geld overmaken naar andere rekeningen. 

In deze versimpelde applicatie zitten een aantal **accounts**. Elk account bevat een aantal **transacties**.

# Terminologie

Hieronder vind je een lijst van termen die we in het programma gebruiken.

## Account

Het account is iemand die kan 'inloggen' op de applicatie. Het account heeft een uniek ID, en heeft een bepaald rekeningnummer.

Deze applicatie bevat geen uitgebreid login systeem, we houden het simpel. Op de `/login` kun je inloggen door een (bestaand) accountnummer in te vullen.

## Transacties

Een transactie (in de code gebruiken we de term `transaction`) behoort tot een bepaald account. Een betaling kan positief of negatief uitpakken:

- een **debet** transactie betekent dat het bedrag **er bij** komt
- een **credit** transactie betekent dat het bedrag **er af** gaat

# Installatie

Download de eindopdracht van ITSLearning, en open deze folder in Visual Studio Code.

Kopieer `.env.example` naar een nieuwe file `.env`. In `.env` pas je de database gegevens aan die voor jou relevant zijn. Hiervoor pas je waarschijnlijk `DB_USER`, `DB_PASSWORD` en `DB_NAME` aan.

Om het programma uit te voeren voer je het commando `go run .` uit. Hierna is je eigen webserver op poort `8080` beschikbaar. Om de webapplicatie te bezoeken ga je naar `http://localhost:8080`.

> **Let op:** wanneer je wijzigingen doet aan je go-code moet je de webserver opnieuw starten om deze nieuwe code te compileren en daardoor te kunnen gebruiken!

# Opdrachten

Dit zijn de opdrachten:

## Eindopdracht 1. Opvragen saldo
> Complexiteit: ★☆☆☆☆

Vanuit het menu kun je naar de pagina `/saldo` navigeren. Op deze pagina kun je snel zien wat het saldo is van het account waarmee je bent ingelogd.

Maak een functie in de `repositories` package om het saldo van het betreffende account te berekenen.

Een aantal nuttige gegevens:
- Deze functie heet `GetSaldoForAccount`
- De functie ontvangt het `account_id`
- De functie geeft een type `float64` terug
- Roep deze functie aan in `handlers/ShowSaldoHandler.go` zodat het juiste saldo wordt getoond

## Eindopdracht 2. Overzicht transacties
> Complexiteit: ★★☆☆☆

We kunnen nu het huidige saldo van het ingelogde account zien, maar we willen graag meer gegevens hebben. We willen namelijk een lijst van transacties hebben zodat we kunnen zien welke bedragen er uit zijn gegaan, en welke bedragen binnen zijn gekomen. De lijst van transacties is te zien op `/transactions`.

Maak een functie in de `repositories` package om alle transacties van het ingelogde account op te halen.

Roep vervolgens de functie aan in de betreffende handler, en geef alle transacties weer in het template in `templates/list-transactions.html`.

Een aantal nuttige gegevens:
- Deze functie heet `GetTransactionsForAccount`
- De functie ontvangt het `account_id`
- De functie geeft een slice met het type `Transaction` terug
- Roep deze functie aan in `handlers/ListTransactionsHandler.go` zodat de lijst van `Transactions` wordt doorgegeven aan het html template

## Eindopdracht 3. Geld opnemen
> Complexiteit: ★★★☆☆

Nu we kunnen zien wat ons saldo is, en welke transacties er hebben plaatsgevonden willen we geld kunnen opnemen van de rekening. Dit kan op `/withdrawal`.

Maak een functie `CanWithdrawFromAccount` in de `repositories` package om te controleren of het opgegeven bedrag wel kan worden afgeschreven van het betreffende account.

Maak daarnaast een functie `WithdrawFromAccount` in de `repositories` package om het opgegeven bedrag af te schrijven (credit) van het account.

Een aantal nuttige gegevens:
- Het formulier is al gegeven (zie `templates/create-withdrawal.html`)
- De validatiefunctie heet `CanWithdrawFromAccount`. Deze ontvangt het `account_id` en de `amount` dat is ingevuld in het formulier, en geeft een `boolean` type terug
- De actiefunctie heet `WithdrawFromAccount`. Deze ontvangt het `account_id` en de `amount` dat is ingevuld in het formulier, en geeft een `Transaction` type terug
- Roep deze functies aan in `handlers/CreateWithdrawalHandler.go` zodat de validatie wordt toegepast, en de opname wordt geregistreerd in het systeem
- Wanneer het bedrag niet kan worden afgeschreven wordt er een foutmelding toegevoegd aan `data.Errors`

## Eindopdracht 4. Geld overmaken
> Complexiteit: ★★★★☆

Naast het opnemen van geld willen we ook overboekingen kunnen laten doen. Dat betekent dat de ingelogd account een bepaald bedrag kan overmaken naar een ander account. 

Implementeer de eerder gebouwde functie `CanWithdrawFromAccount` in de `CreateTransferHandler` om te valideren of het gekozen bedrag kan worden afgeschreven.

Maak daarnaast een functie in de `repositories` package om het opgegeven bedrag af te schrijven (credit) van het huidige account, en bij te schrijven bij het gekozen account. Let ook op dat je de meegegeven beschrijving ook opslaat bij de transacties.

> **Let op: ** Het gaat hier dus om twee transacties! Een `credit` transactie, en een `debet` transactie

## Eindopdracht 5. Importeren
> Complexiteit: ★★★★★

In de _root_ van jouw project vind je `transactions-import.json`. Als je in de applicatie naar `/import-transactions` gaat, kun je hier een JSON file uploaden. De transacties worden dan geimporteerd in het systeem.

Maak een functie in de `repositories` package om het opgegeven bestand in te lezen en de transacties te importeren.

Een aantal nuttige gegevens:
- Deze functie heet `ImportTransactions`
- De functie ontvangt het path naar de geuploade file (`string`)
- In de package `helpers` vind je een method die voor jou het bestand kan inlezen en omzet naar een slice van het opgegeven type

## Eindopdracht 6. Voorbereiding assessment
Ingangseis: Jouw website functioneert volgens de bovenstaande specificaties.

Jouw uitwerking van de eindopdracht lever je in op ITS Learning, in de betreffende inleverbibliotheek (zie Bronnen > Afronding).

Neem de toetsmatrijs door en bereid jezelf voor op vragen die je kunt verwachten bij de verschillende leeruitkomsten (jouw toelichtingen zijn een belangrijk onderdeel van jouw eindcijfer). 

Je hebt een werkende demo en zorgt dat deze klaar staat als je aan de beurt bent. De database van de applicatie is teruggezet naar de beginsituatie (seeder.sql), en de import (eindopdracht 5) is nog NIET uitgevoerd.

Aan het begin van het assessment is jouw eigen laptop opgeladen is en is Visual Studio Code opgestart met jouw programmacode geladen en gereed voor gebruik. In het eerste deel van het assessment geef je een demonstratie van de applicatie aan de hand van een aantal usecases die je moet uitvoeren met jouw demo. Vervolgens doen wij een codereview. In totaal duurt de toets ongeveer 20 minuten.

Tijdens het assessment krijg je de kans om de assessor te overtuigen van de werkende code én van het feit dat jij die code gemaakt hebt. Let op het principe van plagiaat: Samenwerken aan één probleem, betekent niet dat je dezelfde oplossing inlevert!

## EXTENDED Eindopdracht 7. Rood staan
> Complexiteit: ★★★★★★

Deze opdracht is een EXTENDED opdracht, en dus niet verplicht voor het inleveren van jouw eindopdracht.

Op dit moment valideert de functie `CanWithdrawFromAccount` of je een bepaalde opname of transfer kan doen. Je mag dus niet rood staan.

Voeg een stuk extra logica toe aan de applicatie zodat een bepaald account een bepaald negatief saldo mag hebben. De hoogte van dit bedrag kan voor elk account anders zijn. Dat betekent dus dat je wijzigingen zult moeten doen in de database, het `Account` type, en de de go-code in de `repositories` en `handlers` namespace.
