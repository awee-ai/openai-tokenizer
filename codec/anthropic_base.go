package codec

import "github.com/dlclark/regexp2"

func NewAnthropicBase() *Codec {
	anthropicVocabOnce.Do(anthropicVocabInit)

	// Pattern from anthropic.json
	splitRegexp := regexp2.MustCompile(`'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`, regexp2.None)

	return &Codec{
		name:        "anthropic",
		vocabulary:  anthropicVocab,
		splitRegexp: splitRegexp,
		specialTokens: map[string]uint{
			"<EOT>":        0,
			"<META>":       1,
			"<META_START>": 2,
			"<META_END>":   3,
			"<SOS>":        4,
		},
	}
}
