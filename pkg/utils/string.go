package utils

import "strings"

func ConvertString(input string) string {
	var output strings.Builder

	for i, char := range input {
		if i != 0 && 'A' <= char && char <= 'Z' {
			output.WriteRune('_')
		}
		output.WriteRune(char)
	}

	return removeAORM(ToPlural(strings.ToLower(output.String())))
}

// 将英文单词变成复数形式
func ToPlural(word string) string {
	// 声明一些特殊情况的规则
	irregulars := map[string]string{
		"man":   "men",
		"woman": "women",
		"child": "children",
	}

	// 检查是否有特殊情况的规则
	if plural, ok := irregulars[word]; ok {
		return plural
	}

	// 根据一般规则处理其他情况
	// 处理以辅音字母 + y 结尾的单词，将 y 替换为 ies
	if strings.HasSuffix(word, "y") && !isVowel(word[len(word)-2]) {
		return word[:len(word)-1] + "ies"
	}

	// 处理以 s, x, z, ch, sh 结尾的单词，在末尾加上 es
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "z") || strings.HasSuffix(word, "ch") ||
		strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	// 其他情况，在末尾加上 s
	return word + "s"
}

// 检查字符 c 是否是元音字母
func isVowel(c byte) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, rune(c))
}

func removeAORM(str string) string {
	return strings.ReplaceAll(str, "_aorm", "")
}
