package operadora

import (
	"strconv"

	"bitbucket.org/novatrixbr/cursodego/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora que se estar a usar
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
