package vii

//go:generate go-jenny -f=account/cache.go -p=account -t=Cache -k=string -v=*T

//go:generate go-jenny -f=chat/cache.go -p=chat -t=Cache -k=string -v=*Room

//go:generate go-jenny -f=match/cache.go -p=match -t=Cache -k=string -v=*Queue

//go:generate go-jenny -f=game/cache.go -p=game -t=Cache -k=string -v=*T
