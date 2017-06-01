package SE

type CardType byte

var CTYPnull CardType = 0
var CTYPland CardType = 1
var CTYPinstant CardType = 2
var CTYPcreature CardType = 3
var CTYPpermanent CardType = 4

var CardTypes = []*CardType{&CTYPnull, &CTYPland, &CTYPinstant, &CTYPcreature, &CTYPpermanent}
