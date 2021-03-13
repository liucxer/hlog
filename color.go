package hlog

import "fmt"

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

func black(str string) string {
	return textColor(textBlack, str)
}

func red(str string) string {
	return textColor(textRed, str)
}
func yellow(str string) string {
	return textColor(textYellow, str)
}
func green(str string) string {
	return textColor(textGreen, str)
}
func cyan(str string) string {
	return textColor(textCyan, str)
}
func blue(str string) string {
	return textColor(textBlue, str)
}
func purple(str string) string {
	return textColor(textPurple, str)
}
func white(str string) string {
	return textColor(textWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}
