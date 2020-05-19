# mclient
不满于只能通过各种语言的脚本来读写memcached服务的情况，用gomemcache和cobra撸了一个命令行客户端，现在只实现了两个命令：get和set

## get
`mclient -H 127.0.0.1 -P 11211 -K foo`

## set
`mclient -H 127.0.0.1 -P 11211 -K foo -V bar`
