package main

func Hello(name, lang string) string {
	prefix := "Hello"
	if name == "" {
		name = "world"
	}
	if lang == "Thai" {
		prefix = "สวัสดี"
	}
	return prefix + " " + name
}
