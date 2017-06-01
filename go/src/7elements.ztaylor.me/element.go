package SE

type Element byte

var ELEMnull Element = 0
var ELEMwhite Element = 1
var ELEMblue Element = 2
var ELEMgreen Element = 3
var ELEMgold Element = 4
var ELEMred Element = 5
var ELEMindigo Element = 6
var ELEMblack Element = 7

var Elements = []*Element{&ELEMnull, &ELEMwhite, &ELEMblue, &ELEMgreen, &ELEMgold, &ELEMred, &ELEMindigo, &ELEMblack}
