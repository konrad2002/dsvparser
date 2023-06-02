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

## 📋 Standard

Der DSV7 Standard wurde vom Deutschen Schwimm-Verband e.V. (DSV) mit Gültigkeit ab 1.1.2023 herausgegeben und ersetzt damit den vorherigen DSV6 Standard.
Dateien des DSV Standards sind UTF-8 kodierte Textdateien, welche ein Datenformat für den Informationsaustausch im deutschen Schwimmsport darstellen.
Die Spezifikationen das DSV7 Standards lassen sich [🔗 hier](https://www.dsv.de/fileadmin/dsv/documents/Amtliche_Mitteilungen/DSV_Standard07_rot_markiete_Aenderungen.pdf) nachlesen.

## 🏊‍♀️ SwimResults

Im Rahmen der Entwicklung von [SwimResults](https://swimresults.de) kam Bedarf für einen DSV7 Parser auf, welcher in diesem Repository unabhängig von SwimResults implementiert wird.

## 📄 Changelogs

### v0.0.1

- initiale Implementierung der Datentypen des DSV7 Standards (Datum, Uhrzeit und Zeit).
- Model für alle "Elemente" der Formate "Wettkampfdefinitionsliste" und "Wettkampfergebnisliste"
