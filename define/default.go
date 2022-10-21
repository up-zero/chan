package define

var DefaultChanConf = ChanConf{
	Server: []Server{
		{
			Listen:     ":80",
			ServerName: "localhost",
		},
	},
}
