package def
// dream def file parser

import (
	"regexp"
	"toma/dream"
	"toma/parse"
	"strconv"
)


// given a path ro a def file (defpath), return a 
// DBlock object.
func Def(defpath string) dream.DBlock {

	// resulting object
	db := dream.DBlock{}

	// regexp to spilt on - in this case whitespace
	ws_re := regexp.MustCompile(`\s+`)

	// channel of lines with line comments and blank lines removed
	line_ch := make(chan string, 20)

	// channel of tokens split from the lines
	tok_ch := make(chan string, 20)

	go parse.ReadLinesStripped(defpath, "#", line_ch)

	go parse.SplitTokenizer(ws_re, line_ch, tok_ch)

	for token := range tok_ch {

		switch token {
		case "VERSION": parseVersion(&db, tok_ch)
		case "DESIGN": parseDesign(&db, tok_ch)
		case "UNITS": parseUnits(&db, tok_ch)
		case "COMPONENTS": parseComponents(&db, tok_ch)
		case "NETS": parseNets(&db, tok_ch)
		}
	}

	return db
}

// parse the VERSION statement
func parseVersion(db *dream.DBlock, tok_ch <-chan string) {
	frame := parse.TakeN(2, tok_ch)
	db.Version, _ = strconv.ParseFloat(frame[0], 32)
}
	
// parse the DESIGN statement
func parseDesign(db *dream.DBlock, tok_ch <-chan string) {
	frame := parse.TakeN(2, tok_ch)
	db.Design = frame[0]
}
	
// parse the UNITS DISTANCE MICRONS statement
func parseUnits(db *dream.DBlock, tok_ch <-chan string) {
	frame := parse.TakeN(4, tok_ch)
	db.DBU, _ = strconv.Atoi(frame[2])
}
	
// parse the COMPONENTS section
// TODO - refactor for reuse (create a parseSection())
func parseComponents(db *dream.DBlock, tok_ch <-chan string) {
	frame := parse.TakeN(2, tok_ch)
	count, _ := strconv.Atoi(frame[0])
	for i := 0 ; i < count ; i++ {
		frame = parse.TakeUntil(";", tok_ch)
		inst_name := frame[1]
		model_name := frame[2]
		comp := dream.DInst{InstName: inst_name, ModelName: model_name}
		db.Insts[inst_name] = comp
	}
	_ = parse.TakeN(2, tok_ch) // consume the END...
}
	
// parse the NETS section
// TODO - refactor for reuse (create a parseSection())
func parseNets(db *dream.DBlock, tok_ch <-chan string) {
	frame := parse.TakeN(2, tok_ch)
	count, _ := strconv.Atoi(frame[0])
	for i := 0 ; i < count ; i++ {
		frame = parse.TakeUntil(";", tok_ch)
		net_name := frame[1]
		net := dream.DNet{Name: net_name}
		db.Nets[net_name] = net
	}
	_ = parse.TakeN(2, tok_ch) // consume the END...
}

	
