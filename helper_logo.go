package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CreateInitial(name string) []byte {

	rand.NewSource(time.Now().UnixNano())
	bg := Colors[rand.Intn(len(Colors))]

	words := strings.Fields(name)

	initials := ""
	for _, w := range words {
		if len(w) > 0 {
			initials += strings.ToUpper(string(w[0]))
		}
	}

	if len(initials) > 2 {
		initials = initials[:2]
	}

	svg := fmt.Sprintf(
		`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 100 100">
			<rect width="100" height="100" rx="12" fill="#%s"/>
			<text x="50" y="50"
				text-anchor="middle"
				dominant-baseline="central"
				font-family="Arial, Helvetica, sans-serif"
				font-size="40"
				font-weight="600"
				fill="#ffffff">%s</text>
		</svg>`,
		128, 128, bg, initials,
	)

	return []byte(svg)
}
