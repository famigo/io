package io

//Register is the channel to communicate between the console hardware and game software
type Register chan byte

var (
	PPUCTRL   = make(Register, 0x2000)
	PPUMASK   = make(Register, 0x2001)
	PPUSTATUS = make(Register, 0x2002)
	OAMADDR   = make(Register, 0x2003)
	OAMDATA   = make(Register, 0x2004)
	PPUSCROLL = make(Register, 0x2005)
	PPUADDR   = make(Register, 0x2006)
	PPUDATA   = make(Register, 0x2007)
	OAMDMA    = make(Register, 0x4014)
)
