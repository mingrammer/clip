<br><br>

<h1 align="center">cfmt</h1>

<p align="center">  [![img](https://img.shields.io/badge/license-MIT-blue.svg)]()    [![img](https://godoc.org/github.com/mingrammer/cfmt?status.svg)](%EE%80%96https://godoc.org/github.com/mingrammer/cfmt%EE%80%97)  [![img](https://travis-ci.org/mingrammer/pyreportcard.svg?branch=master)](%EE%80%96https://travis-ci.org/mingrammer/pyreportcard%EE%80%97)  </p>

<p align="center"> Contextual fmt </p>

<br><br><br>

It provides contextual formatting functions that have nearly identical usage of the fmt package. The ideas were borrowed from bootstrap's contextual color classes.

## Installation

```
go get github.com/mingrammer/cfmt
```

## Usage

```
import (
	"github.com/mingrammer/cfmt"
)

func main() {
	cfmt.Success("User was created successfully")
	cfmt.Infoln("Here are some candidates")
	cfmt.Warningf("%s is not valid integer value", "123a")
	log.Fatal(cfmt.Serror("Only numeric is allowed, got %s", "123.456a"))
}
```

![cfmt output](images/output.png)

## Contextual Functions

> Note: `cfmt.Errorf` function does not do same things to `fmt.Errorf`, but just same as `fmt.Printf` with red colored text.

- Success (Green)
  - Fsuccess, Fsuccessf, Fsuccessln
  - Success, Successf, Successln
  - Ssuccess, Ssuccessf, Ssuccessln
- Info (Cyan)
  - Finfo, Finfof, Finfoln
  - Info, Infof, Infoln
  - Sinfo, Sinfof, Sinfoln
- Warning (Yellow)
  - Fwarning, Fwarningf, Fwarningln
  - Warning, Warningf, Warningln
  - Swarning, Swarningf, Swarningln
- Error (Red)
  - Ferror, Ferrorf, Ferrorln
  - Error, Errorf, Errorln
  - Serror, Serrorf, Serrorln
