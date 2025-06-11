package main

import "testing"

type TestCases struct {
	Input, Output string
}

var (
	Tests = []TestCases{
		{
			Input:  "echo hello world",
			Output: "echo hello world",
		},
		{
			Input:  "I should stop SHOUTING (low)",
			Output: "I should stop shouting",
		},
		{
			Input:  "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			Output: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			Input:  "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			Output: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			Input:  "There is no greater agony than bearing a untold story inside you.",
			Output: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			Input:  "Punctuation tests are ... kinda boring ,what do you think ?",
			Output: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			Input:  "1E (hex) files were added",
			Output: "30 files were added",
		},
		{
			Input:  "It has been 10 (bin) years",
			Output: "It has been 2 years",
		},
		{
			Input:  "Ready, set, go (up) !",
			Output: "Ready, set, GO!",
		},
		{
			Input:  "I should stop SHOUTING (low)",
			Output: "I should stop shouting",
		},
		{
			Input:  "Welcome to the Brooklyn bridge (cap)",
			Output: "Welcome to the Brooklyn Bridge",
		},
		{
			Input:  "This is so exciting . (up, 2)",
			Output: "This is SO EXCITING.",
		},
		{
			Input:  "I was sitting over there ,and then BAMM ! ! ",
			Output: "I was sitting over there, and then BAMM!!",
		},
		{
			Input:  "I was thinking . . . You were right",
			Output: "I was thinking... You were right",
		},
		{
			Input:  "I am exactly how they describe me: ' awesome '",
			Output: "I am exactly how they describe me: 'awesome'",
		},
		{
			Input:  "As Elton John said: ' I am the most well-known homosexual in the world '",
			Output: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			Input:  "There it was. A amazing rock!",
			Output: "There it was. An amazing rock!",
		},
		{
			Input:  "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			Output: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			Input:  "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			Output: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			Input:  "Don not be sad ,because sad backwards is das . And das not good",
			Output: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			Input:  "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			Output: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			Input:  "a apple",
			Output: "an apple",
		},
		{
			Input:  "A APPLE",
			Output: "An APPLE",
		},
		{
			Input:  "A Apple",
			Output: "An Apple",
		},
		{
			Input:  "A a apple",
			Output: "An an apple",
		},
		{
			Input:  "a A apple",
			Output: "an An apple",
		},
		{
			Input:  "a hour",
			Output: "an hour",
		},
		{
			Input:  "a or b",
			Output: "an or b",
		},
		{
			Input:  "a and the",
			Output: "an and the",
		},
		{
			Input:  "I am a optimist, but a optimist",
			Output: "I am an optimist, but an optimist",
		},
		{
			Input:  "Good morning, mr harold wilson (cap, 3)",
			Output: "Good morning, Mr Harold Wilson",
		},
		{
			Input:  "one(low)two(cap)three(up)",
			Output: "one Two THREE",
		},
		{
			Input:  "one (up) two (up) three (cap, 3)",
			Output: "One Two Three",
		},
		{
			Input:  "Abc (L(low)O(low)W(low))",
			Output: "Abc (l o w)",
		},
		{
			Input:  "hello (up) world (cap)",
			Output: "HELLO World",
		},
		{
			Input:  "done             (          up     , 1     )",
			Output: "DONE",
		},
		{
			Input:  "one two (cap, a (hex) (bin))",
			Output: "One Two",
		},
		{
			Input:  "me (up, 10(bin))",
			Output: "ME",
		},
		{
			Input:  "WORD (LOW(low))",
			Output: "word",
		},
		{
			Input:  "it (cap) was a 'amazing' (cap) experience   ?!",
			Output: "It was an 'Amazing' experience?!",
		},
		{
			Input:  "a 'amazing'",
			Output: "an 'amazing'",
		},
		{
			Input:  "a (up) alem (up ,2)",
			Output: "AN ALEM",
		},
		{
			Input:  "A (low) ALEM (low)",
			Output: "an alem",
		},
		{
			Input:  "a alem (up,2)",
			Output: "AN ALEM",
		},
		{
			Input:  "A ALEM (low, 2)",
			Output: "an alem",
		},
		{
			Input:  "Hello:world.How:are you?",
			Output: "Hello: world. How: are you?",
		},
		{
			Input:  "Elton       John",
			Output: "Elton John",
		},
		{
			Input:  " wanna      chose",
			Output: "wanna chose",
		},
		{
			Input:  "word!     !!!!!!!!!!!  ........   ,,,,,,,,,  ??????????     :::::::::;;;;;;;;;;;;;       :",
			Output: "word!!!!!!!!!!!!........,,,,,,,,,??????????:::::::::;;;;;;;;;;;;;:",
		},
		{
			Input:  "a (hex) (bin)",
			Output: "2",
		},
		{
			Input:  "a (hex)",
			Output: "10",
		},
		{
			Input:  "AB (Hex)(low)",
			Output: "171",
		},
		{
			Input:  "ab (up)(HeX)(low)",
			Output: "171",
		},
		{
			Input:  "010101101001000 (bin)",
			Output: "11080",
		},
		{
			Input:  "hi 'hi' hi",
			Output: "hi 'hi' hi",
		},
		{
			Input:  "hi' hi' hi",
			Output: "hi 'hi' hi",
		},
		{
			Input:  "hi'hi' hi",
			Output: "hi 'hi' hi",
		},
		{
			Input:  "hi 'hi ' hi",
			Output: "hi 'hi' hi",
		},
		{
			Input:  "hi ' hi ' hi",
			Output: "hi 'hi' hi",
		},
		{
			Input:  "hi 'hi",
			Output: "hi 'hi",
		},
		{
			Input:  "  ' hi '   hi'",
			Output: "'hi' hi'",
		},
		{
			Input:  "one'two' three'",
			Output: "one 'two' three'",
		},
		{
			Input:  "we ' ll i ' m",
			Output: "we 'll i' m",
		},
		{
			Input:  "'one''two''''four''five''''seven''eight''''ten'",
			Output: "'one' 'two' '' 'four' 'five' '' 'seven' 'eight' '' 'ten'",
		},
		{
			Input:  "' awesome '",
			Output: "'awesome'",
		},
		{
			Input:  "'Transform inside (up)'",
			Output: "'Transform INSIDE'",
		},
		{
			Input:  "ONE, TWO, THREE... (low, 9999999999999999999)",
			Output: "one, two, three...",
		},
	}
)

func TestGoReloaded(t *testing.T) {
	for i, test := range Tests {
		if output := procedures(test.Input); output != test.Output {
			t.Errorf("Output in id %d %s not equal to expected %s", i, output, test.Output)
		}
	}
}
