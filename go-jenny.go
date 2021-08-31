package vii

//go:generate go-jenny -f=account/cache.go -p=account -t=Cache -k=string -v=*T

//go:generate go-jenny -f=chat/cache.go -p=chat -t=Cache -k=string -v=*Room

//go:generate go-jenny -f=gameserver/cache.go -p=gameserver -i=github.com/zachtaylor/7elements/game, -t=Cache -k=string -v=*game.T
