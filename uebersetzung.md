# effective_go

Dies ist eine inoffizielle Übersetzung von [Effective Go](https://golang.org/doc/effective_go.html). Die Code Beispiele wurden aus der Originalfassung verwendet.

## Formatierung

Die Formatierung des Quellcodes wird meistens am stärksten Diskutiert, aber am inkonsquentesten umgesetzt. Jeder kann sich unterschiedliche Formattierungsregeln aneignen, aber es ist besser, wenn das nicht notwendig ist und jeder den gleichen Stil verwendet. Dadurch kann viel Zeit gespart werden, da dieses Thema nicht jedes mal neu diskutiert werden muss. Das Problem an der Stelle ist, wie man dieses Ziel ohne umfangreichen und komplizierten Style Guide umsetzen kann.

Mit Go fahren wir einen unüblichen Ansatz und lassen die Maschine über die Formatierung bestimmen. Das `gofmt` Programm (auch über den Befehl `go fmt` zugänglich) liest das Go Programm ein und formatiert den Quellcode in einem Standart-Stil. Dabei werden Einrückungen gesetzt und eine vertikale Ausrichtung vorgenommen. Wenn notwendig werden dabei auch Kommentare neu formattiert. Wenn unklar ist, wie gewisse Passagen formatiert werden sollen, kann einfach `gofmt` ausgeführt werden. Wenn das Ergebnis nicht korrekt aussieht, sollte das Programm angepasst werden (oder ein Bug zu `gofmt` gemeldet werden). An der Stelle sollte jedoch kein Workaround erstellt werden.

Hier ein kleines Beispiel, welches die Funktionsweise zeigen soll. An der Stelle ist es nicht notwendig die Kommentare zu den Feldern eines Struct auszurichten. `gofmt` erledigt das automatisch. Aus dem folgenden Code

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

`gofmt`richtet hier die einzelnen Spalten aus:

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

Der komplette Go Code der Standard Pakete wurde mit `gofmt` formatiert.

Ein paar Details sind jedoch weiterhin offen:

* Einrückung: Wir verwenden Tabs für Einrückungen und `gofmt` erzeugt diese im Standardfall. Spaces sollen nur verwendet werden, wenn diese wirklich notwendig ist
* Zeilenlänge: Go hat keine Einschränkung bezüglich der Zeilenlänge. Aber auch keine Angst vor unendlich langen Zeilen. Wenn eine Zeile gefühlt zu lange wird, dann kann diese einfach umgebrochen werden und mit einem extra Tab eingerückt werden.
* Klammern: Go verwendet weniger Klammern als C oder Java. Die Kontroll Struktueren (`if`, `for`, `switch`) benötigen laut Syntax keine Klammern. Auch die Rangordnung der Operatoren ist kürzer und klarer. Das Beispiel

```go
x<<8 + y<<16
```

berechnet sich, wie es die Leerzeichen definieren.

## Kommentare

Go verwendet C-Stil /* */ Block Kommentare und C++ Stil // Zeilen-Kommentare. Zeilen-Kommentare sind die Norm, wobei Block-Kommentare meistens als Paket Dokumentation verwendet werden. Diese sind jedoch auch innerhalb eines Ausdruckes nützlich oder um größere Bereiche von Code zu deaktivieren.

Das Programm - und Webserver - `godoc` geht durch die Go Quelldateien und erzeugt daraus eine Dokumentation für das jeweilige Paket. Der Kommentar, welcher vor der jeweiligen Deklaration von exportierten Elementen steht, wird als Beschreibung für das Element in der Dokumentation verwendet. Die Art und Weise wie diese Kommentare gestaltet wurden ist ausschlaggebend für die Qualität der Dokumentation welche `godoc` erzeugt.

Für jedes Paket sollte einen _Paket Kommentar_ erstellt werden, welcher vor der Paket Deklaration steht. Für Pakete mit mehr als einer Datei muss der Paket Kommentar nur in einem File definiert werden, dabei ist es egal in welcher Datei dieser steht. Der Paket Kommentar soll das Paket vorstellen und Informationen zu der Idee hinter dem Paket geben. Dieser wird als erstes auf der `godoc` Seite dargestellt und soll eine Einleitung für die folgende Detailierte Dokumentation der einzelnen Funktionen sein.

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

Bei einfachen Paketen kann dieser Kommentar auch wie folgt aussehen:

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

Kommentare benötigen keine extra Formatelemente wie zum Beispiel ein Banner aus Sternen. Der generierte Output hat auch keine feste Breite, so dass keine Leerzeichen für eine Ausrichtung der Kommentare verwendet werden sollten. `godoc` kümmert sich darum. Die Kommentare werden als reiner Text interpretiert, deshalb sollte kein HTML oder andere Notierungen wie `_this_` benutzt werden, da diese exakt so wieder gegeben werden. Es gibt an der Stelle eine Ausnahme, welche durch `godoc`durchgeführt wird. Eingerückter Text wird mit einer nichtproportionalien Schriftart (Festbreitenschrift) dargestellt, welche für Code Schnipsel verwendet werden kann. Der Paket Kommentar für das [fmt package](https://golang.org/pkg/fmt/) verwendet diesen Effekt.

Abhängig von dem jeweiligen Kontext kann es passieren, dass die Kommentare durch `godoc` nichtmal neu formatiert werden, deshalb sollte sichergestellt sein, dass diese gut aussehen. Das bedeutet: verwende eine korrekte Rechtschreibung, Punktuation, Satzstruktur und breche längere Zeilen um.

Innerhalb eines Pakets wird jeder Kommentar überhalb einer Deklaration als doc Kommentar (Dokumentation Kommentar) dargestellt. Jedes exportierte Element sollte deshalb ein doc Kommentar besitzen.

Doc Kommentare sollten am besten als ganze Sätze formuliert werden, wodurch diese vielfältig automatisiert verwendet werden können. Der erste Satz sollte eine Zusammenfassung sein, welches mit dem Namen des Elements beginnt.

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

Wenn jeder doc Kommentar mit dem Namen des Elements beginnt, welchen er Beschreibt, kann der Output von `godoc` mit `grep` gut ausgewertet werden. Wenn man beispielsweise den Namen der Funktion "Compile" vergessen hat, aber nach der _parsing_ Funktion sucht, so kann man folgenden Befehl verwenden:

```
$ godoc regexp | grep -i parse
```

Wenn alle doc Kommentare mit "This function..." anfangen würden, wäre `grep` nicht in der Lage den Namen zu finden. Da aber jeder doc Kommentar mit dem gesuchten Namen beginnt, kann der Output wie folgt aussehen, welches gleich das gesuchte Wort beinhaltet:

```
$ godoc regexp | grep parse
    Compile parses a regular expression and returns, if successful, a Regexp
    parsed. It simplifies safe initialization of global variables holding
    cannot be parsed. It simplifies safe initialization of global variables
$
```

Die Go Syntax erlaubt es Deklarationen zu gruppieren. Ein doc Kommentar kann so eine Gruppe von Variablen oder Konstanten beschreiben. Da hierdurch mehrere Deklarationen beschrieben werden ist so ein Kommentar eher oberflächlich.

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

So eine Gruppierung kann auch eine Beziehung zwischen den einzelnen Elementen ausdrücken, wie z.B. dass mehrere Variablen durch ein mutex geschützt werden.

```go
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```

## Namen

### Paketnamen

### Getters

### Interface Namen

### MixedCaps