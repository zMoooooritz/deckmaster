package main

import (
	"strconv"
	"strings"
)

var keycodes = map[string]uint8{
	"Esc":              1,
	"Num1":             2,
	"Num2":             3,
	"Num3":             4,
	"Num4":             5,
	"Num5":             6,
	"Num6":             7,
	"Num7":             8,
	"Num8":             9,
	"Num9":             10,
	"Num0":             11,
	"Minus":            12,
	"Equal":            13,
	"Backspace":        14,
	"Tab":              15,
	"Q":                16,
	"W":                17,
	"E":                18,
	"R":                19,
	"T":                20,
	"Y":                21,
	"U":                22,
	"I":                23,
	"O":                24,
	"P":                25,
	"Leftbrace":        26,
	"Rightbrace":       27,
	"Enter":            28,
	"Leftctrl":         29,
	"A":                30,
	"S":                31,
	"D":                32,
	"F":                33,
	"G":                34,
	"H":                35,
	"J":                36,
	"K":                37,
	"L":                38,
	"Semicolon":        39,
	"Apostrophe":       40,
	"Grave":            41,
	"Leftshift":        42,
	"Backslash":        43,
	"Z":                44,
	"X":                45,
	"C":                46,
	"V":                47,
	"B":                48,
	"N":                49,
	"M":                50,
	"Comma":            51,
	"Dot":              52,
	"Slash":            53,
	"Rightshift":       54,
	"Kpasterisk":       55,
	"Leftalt":          56,
	"Space":            57,
	"Capslock":         58,
	"F1":               59,
	"F2":               60,
	"F3":               61,
	"F4":               62,
	"F5":               63,
	"F6":               64,
	"F7":               65,
	"F8":               66,
	"F9":               67,
	"F10":              68,
	"Numlock":          69,
	"Scrolllock":       70,
	"Kp7":              71,
	"Kp8":              72,
	"Kp9":              73,
	"Kpminus":          74,
	"Kp4":              75,
	"Kp5":              76,
	"Kp6":              77,
	"Kpplus":           78,
	"Kp1":              79,
	"Kp2":              80,
	"Kp3":              81,
	"Kp0":              82,
	"Kpdot":            83,
	"Zenkakuhankaku":   85,
	"102Nd":            86,
	"F11":              87,
	"F12":              88,
	"Ro":               89,
	"Katakana":         90,
	"Hiragana":         91,
	"Henkan":           92,
	"Katakanahiragana": 93,
	"Muhenkan":         94,
	"Kpjpcomma":        95,
	"Kpenter":          96,
	"Rightctrl":        97,
	"Kpslash":          98,
	"Sysrq":            99,
	"Rightalt":         100,
	"Linefeed":         101,
	"Home":             102,
	"Up":               103,
	"Pageup":           104,
	"Left":             105,
	"Right":            106,
	"End":              107,
	"Down":             108,
	"Pagedown":         109,
	"Insert":           110,
	"Delete":           111,
	"Macro":            112,
	"Mute":             113,
	"Volumedown":       114,
	"Volumeup":         115,
	"Power":            116,
	"Kpequal":          117,
	"Kpplusminus":      118,
	"Pause":            119,
	"Scale":            120,
	"Kpcomma":          121,
	"Hangeul":          122,
	"Hanja":            123,
	"Yen":              124,
	"Leftmeta":         125,
	"Rightmeta":        126,
	"Compose":          127,
	"Stop":             128,
	"Again":            129,
	"Props":            130,
	"Undo":             131,
	"Front":            132,
	"Copy":             133,
	"Open":             134,
	"Paste":            135,
	"Find":             136,
	"Cut":              137,
	"Help":             138,
	"Menu":             139,
	"Calc":             140,
	"Setup":            141,
	"Sleep":            142,
	"Wakeup":           143,
	"File":             144,
	"Sendfile":         145,
	"Deletefile":       146,
	"Xfer":             147,
	"Prog1":            148,
	"Prog2":            149,
	"Www":              150,
	"Msdos":            151,
	"Coffee":           152,
	"Direction":        153,
	"Cyclewindows":     154,
	"Mail":             155,
	"Bookmarks":        156,
	"Computer":         157,
	"Back":             158,
	"Forward":          159,
	"Closecd":          160,
	"Ejectcd":          161,
	"Ejectclosecd":     162,
	"Nextsong":         163,
	"Playpause":        164,
	"Previoussong":     165,
	"Stopcd":           166,
	"Record":           167,
	"Rewind":           168,
	"Phone":            169,
	"Iso":              170,
	"Config":           171,
	"Homepage":         172,
	"Refresh":          173,
	"Exit":             174,
	"Move":             175,
	"Edit":             176,
	"Scrollup":         177,
	"Scrolldown":       178,
	"Kpleftparen":      179,
	"Kprightparen":     180,
	"New":              181,
	"Redo":             182,
	"F13":              183,
	"F14":              184,
	"F15":              185,
	"F16":              186,
	"F17":              187,
	"F18":              188,
	"F19":              189,
	"F20":              190,
	"F21":              191,
	"F22":              192,
	"F23":              193,
	"F24":              194,
	"Playcd":           200,
	"Pausecd":          201,
	"Prog3":            202,
	"Prog4":            203,
	"Dashboard":        204,
	"Suspend":          205,
	"Close":            206,
	"Play":             207,
	"Fastforward":      208,
	"Bassboost":        209,
	"Print":            210,
	"Hp":               211,
	"Camera":           212,
	"Sound":            213,
	"Question":         214,
	"Email":            215,
	"Chat":             216,
	"Search":           217,
	"Connect":          218,
	"Finance":          219,
	"Sport":            220,
	"Shop":             221,
	"Alterase":         222,
	"Cancel":           223,
	"Brightnessdown":   224,
	"Brightnessup":     225,
	"Media":            226,
	"Switchvideomode":  227,
	"Kbdillumtoggle":   228,
	"Kbdillumdown":     229,
	"Kbdillumup":       230,
	"Send":             231,
	"Reply":            232,
	"Forwardmail":      233,
	"Save":             234,
	"Documents":        235,
	"Battery":          236,
	"Bluetooth":        237,
	"Wlan":             238,
	"Uwb":              239,
	"Unknown":          240,
	"VideoNext":        241,
	"VideoPrev":        242,
	"BrightnessCycle":  243,
	"BrightnessZero":   244,
	"DisplayOff":       245,
}

func formatKeycodes(keycode string) string {
	for k, v := range keycodes {
		if strings.ToLower(keycode) == strings.ToLower(k) {
			keycode = strconv.Itoa(int(v))
			break
		}
	}

	return keycode
}
