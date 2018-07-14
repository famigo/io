package ppu

import (
	"github.com/famigo/asm"
	"github.com/famigo/io"
)

// Nametable addresses
const (
	//NameTableTopLeft is the address of the name table localized at top left
	NameTableTopLeft = uint16(0x2000 + 0x0400*iota)
	NameTableTopRight
	NameTableBottomLeft
	NameTableBottomRight
)

// Screen dimensions
const (
	//ScreenWidthTiles is the number of tiles of a screen
	ScreenWidthTiles = 32

	//ScreenHeightTiles is the number of tiles of a screen
	ScreenHeightTiles = 30
)

// CTRL flags
const (
	SelectNameTableAtTopRight = byte(1 << iota)
	SelectNameTableAtBottomLeft
	IncrementVramBy32GoingDown
	SelectRightPatternTableFor8x8Sprites
	SelectRightPatternTableForBackground
	Enable8x16Sprites
	EnableOutputColorOnEXT
	EnableNMI

	SelectNameTebleAtBottomRight = byte(3)
	SelectNameTableAtTopLeft     = byte(0)
	IncrementVramBy1GoingAcross
	SelectLeftPatternTableFor8x8Sprites
	SelectLeftPatternTableForBackground
	Enable8x8Sprites
	DisableNMI
)

// MASK flags
const (
	EnableGreyscale = byte(1 << iota)
	ShowBackgroundInLefmostColumn
	ShowSpritesInLeftmostColumn
	ShowBackground
	ShowSprites
	EmphasizeRed
	EmphasizeGreen
	EmphasizeBlue
	DisableGreyscale = byte(0)
	HideBackgroundInLefmostColumn
	HideSpritesInLeftmostColumn
	HideBackground
	HideSprites
)

// Registers
var (
	/*
		CTRL is the PPUCTRL register

			7  bit  0
			---- ----
			VPHB SINN
			|||| ||||
			|||| ||++- Base nametable address
			|||| ||    (0 = $2000; 1 = $2400; 2 = $2800; 3 = $2C00)
			|||| |+--- VRAM address increment per CPU read/write of PPUDATA
			|||| |     (0: add 1, going across; 1: add 32, going down)
			|||| +---- Sprite pattern table address for 8x8 sprites
			||||       (0: $0000; 1: $1000; ignored in 8x16 mode)
			|||+------ Background pattern table address (0: $0000; 1: $1000)
			||+------- Sprite size (0: 8x8; 1: 8x16)
			|+-------- PPU master/slave select
			|          (0: read backdrop from EXT pins; 1: output color on EXT pins)
			+--------- Generate a NMI at the start of the vertical blanking interval (0: off; 1: on)
	*/
	CTRL = io.PPUCTRL

	/*
		MASK blablabla
	*/
	MASK = io.PPUMASK

	/*
		STATUS blablabla
	*/
	STATUS = io.PPUSTATUS

	/*
		SCROLL blablabla
	*/
	SCROLL = io.PPUSCROLL

	/*
		ADDR blablabla
	*/
	ADDR = io.PPUADDR

	/*
		DATA blablabla
	*/
	DATA = io.PPUDATA
)

//SetNameTableTile sets the value of one cell of a name table
func SetNameTableTile(nametable uint16, row byte, col byte, tile byte) {
	asm.Printfln("  LDA %v", STATUS)
	offset := int16(row*32 + col)
	asm.Printfln("  CLC")
	asm.Printfln("  LDA %v + 0", offset)
	asm.Printfln("  ADC %v + 0", nametable)
	asm.Printfln("  STA %v + 0", nametable)
	asm.Printfln("  LDA %v + 1", offset)
	asm.Printfln("  ADC %v + 1", nametable)
	asm.Printfln("  STA %v", ADDR)
	asm.Printfln("  LDA %v + 0", nametable)
	asm.Printfln("  STA %v", ADDR)
	asm.Printfln("  LDA %v", tile)
	asm.Printfln("  STA %v", DATA)
}

//SetBackgroundPallete loads one of the background palettes
func SetBackgroundPallete(index byte, pallete [4]byte) {
	asm.Printfln("  LDA #>$3F00")
	asm.Printfln("  STA %v", ADDR)
	asm.Printfln("  LDA %v", index)
	asm.Printfln("  ASL A")
	asm.Printfln("  ASL A")
	asm.Printfln("  CLC")
	asm.Printfln("  ADC #<$3F00")

	asm.Printfln("  LDY #0")
	asm.Printfln("-loop:")
	asm.Printfln("  LDA (%v), Y", pallete)
	asm.Printfln("  STA %v", DATA)
	asm.Printfln("  INY")
	asm.Printfln("  CPY #4")
	asm.Printfln("  BNE -loop")
}
