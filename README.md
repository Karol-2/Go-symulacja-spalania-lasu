<!DOCTYPE html>
<html>
<head>
 
</head>
<body>
  <h1>Symulacja Spalania Lasu</h1>

<h2>Opis projektu</h2>
  <p>
    Ten program symuluje spalanie się lasu przy użyciu języka Go.
  </p>

<h3>Oznaczenia</h3>
  <ul>
    <li>0 - puste pole (brak drzewa)</li>
    <li>1 - drzewo (nie spalone)</li>
    <li>2 - ogień (drzewo spalone)</li>
  </ul>

<h2>Mozliwości programu</h2>
  <p>
    Program ma następujące możliwości:
  </p>
  <ol>
    <li>Pojedyncza symulacja 15x10</li>
    <li>Symulacja 15x10 dla podanej ilości drzew i ilości eksperymentów (testy po 10)</li>
    <li>Symulacja na wszystkie dowolne parametry</li>
    <li>Analiza danych z pliku data.csv</li>
    <li>Wyczyszczenie pliku data.csv</li>
  </ol>
<h2>Analiza</h2>
  <p>
    Program korzysta z pliku <em>data.csv</em>, który zawiera dane dotyczące symulacji.
    Dane te są grupowane według ilości drzew w lesie, a następnie obliczany jest średni procent drzew,
    które przetrwały po symulacji. Wyniki są podawane w losowej kolejności.
  </p>

<h2>Uruchomienie</h2>
  <ol>
    <li>Upewnij się, że masz zainstalowany język Go na swoim systemie.</li>
    <li>Pobierz wszystkie pliki projektu.</li>
    <li>Otwórz terminal i przejdź do katalogu projektu.</li>
    <li>Skompiluj program, wpisując w terminalu komendę: <code>go build</code>.</li>
    <li>Uruchom skompilowany plik, wpisując w terminalu komendę: <code>./nazwa_pliku</code> (gdzie <code>nazwa_pliku</code> to nazwa wygenerowanego pliku wykonywalnego).</li>
    <li>Alternatywnie, użyj <code> go run main.go</code></li>
    <li>Program rozpocznie symulację spalania lasu na podstawie danych z pliku <em>data.csv</em>.</li>
  </ol>

</body>
</html>
