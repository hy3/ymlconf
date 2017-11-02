# ymlconf

## Description
Get and set properties in YAML file, by like "git config"


## Installation
This tool can be installed with the `go get` command.

    go get github.com/hy3/ymlconf


## Usage
Show all properties in YAML file.

    ymlconf -f filePath -l
    ymlconf -f filePath -list

Show value of the property.

    ymlconf -f filePath path.to.property

Set value to the perperty.

    ymlconf -f filepath path.to.property value


## Author
Takahiro Honda (a.k.a hy3)
