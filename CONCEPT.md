# Konzept

## Reader

Der Reader erhält als Eingabe eine Datei und erstellt abhängig vom Eingabeformat
einen Struct vom jeweiligen Typen (Wettkampfdefinitionsliste, Wettkampfergebnisliste).
Anschließend gibt er jede Zeile an den Parser, welcher ein `interface{}` zurück gibt.
Dieses Interface wird mit einem Switch einem der "Elemente" des Standards zugeordnet
und anschließend an der passenden Stelle im Ausgabe-Struct ergänzt.

## Parser

Vom Reader erhält der Parser die einzelnen Zeilen der DSV-Datei. Er übergibt diese dem
Scanner und bekommt einen Token mit einer Reihe von Werten zurück. Je nachdem welchem
"Element" des Standards sich der Token zuordnen lässt, werden die folgenden Werte
als Attribute interpretiert und anschließend in einem zum Element passenden Struct
geparst zurückgegeben.

## Scanner

Die übergebene Zeichenkette liest der Scanner Zeichen für Zeichen ein. Er extrahiert
den Token und reiht die folgenden Semikolon-getrennten Werte in ein Array.