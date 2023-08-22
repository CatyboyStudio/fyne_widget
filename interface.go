package fyne_widget

var GetI18n func(key string) string = SimpleI18nGet

var simple_i18n map[string]string

func SimpleI18nGet(key string) string {
	if simple_i18n != nil {
		if v, ok := simple_i18n[key]; ok {
			return v
		}
	}
	return key
}

func SimpleI18nSet(key string, v string) {
	if simple_i18n == nil {
		simple_i18n = make(map[string]string)
	}
	simple_i18n[key] = v
}
