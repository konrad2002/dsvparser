# 📝 dsvparser

[![Go](https://github.com/konrad2002/dsvparser/actions/workflows/go.yml/badge.svg)](https://github.com/konrad2002/dsvparser/actions/workflows/go.yml)

<img src="dsvparser.png" align="right" alt="dsvparser logo" width="175">

Dieses Go "Package" enthält einen Parser für das Datenaustauschdateiformat des Deutschen Schwimm-Verband e.V. (DSV).
Die aktuelle Version soll die Standards DSV6 und DSV7 unterstützen.

## 💻 Entwicklung

Dieses Projekt befindet sich momentan in der Entwicklung und viele Funktionen fehlen aktuell noch.
Zum Mitwirken am Projekt können entweder Pull Requests gestellt werden oder @konrad2002 kontaktiert werden.

## 📥 Nutzung

Das Package darf von jedem uneingeschränkt für kommerzielle und nicht-kommerzielle Projekte verwendet werden.

### Installation

Importieren des Packages:

```sh
go get github.com/konrad2002/dsvparser@v1.1.0
```

### Beispiel

Verwendung zum Parsen von lokaler DSV7 Datei und Auslesen des Veranstalters:

```go
import (
	// ...
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/parser"
)

func VeranstaltungsortPlz() string {
	dat, err := os.ReadFile("definition.dsv7")
	if err != nil {
		panic(err)
	}
	
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	
	res, err := r.Read()
	if err != nil {
		panic(err)
	}
	
	def := res.(*model.Wettkampfdefinitionsliste)
	
	return def.Veranstaltungsort.PLZ
}

```

## 📋 Standard

Der DSV7 Standard wurde vom Deutschen Schwimm-Verband e.V. (DSV) mit Gültigkeit ab 1.1.2023 herausgegeben und ersetzt damit den vorherigen DSV6 Standard.
Dateien des DSV Standards sind UTF-8 kodierte Textdateien, welche ein Datenformat für den Informationsaustausch im deutschen Schwimmsport darstellen.
Die Spezifikationen das DSV7 Standards lassen sich [🔗 hier](https://www.dsv.de/fileadmin/dsv/documents/Amtliche_Mitteilungen/DSV_Standard07_rot_markiete_Aenderungen.pdf) nachlesen.

## 🏊‍♀️ SwimResults

Im Rahmen der Entwicklung von [SwimResults](https://swimresults.de) kam Bedarf für einen DSV7 Parser auf, welcher in diesem Repository unabhängig von SwimResults implementiert wird.

## 📄 Changelogs

### v1.1.0

- Behebt einen Bug, bei dem einige Methoden nicht `public` waren

### v1.0.0

- erste funktionsfähige Vollversion
- Unterstützt Wettkampfdefinitionslisten und Wettkampfergebnislisten
- kann dieser aus Buffer parsen
- keine Unterstützung für Staffeln
- keine Unterstützung für Vereinslisten

### v0.1.0

- erste Implementierung für Parser, Scanner und Reader
- kann nun Wettkampfdefinitionslisten und Wettkampfergebnislisten parsen

### v0.0.3

- make importable using `import "github.com/konrad2002/dsvparser"`

### v0.0.2

- Weitere Datentypen: Geschlecht, Kampfgericht Position, Meldegeld Typ und Wertungsklasse
- Model für "Elemente" des Formats "Wettkampfergebnisliste"
- vorerst KEINE Unterstützung für Staffelergebnisse

### v0.0.1

- initiale Implementierung der Datentypen des DSV7 Standards (Datum, Uhrzeit und Zeit).
- Model für alle "Elemente" des Formats "Wettkampfdefinitionsliste"
