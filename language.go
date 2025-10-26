/*
 * Copyright © 2021-present Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lingua

// Language is the type used for enumerating the 10 languages which can
// be detected by Lingua.
//
//go:generate stringer -type=Language
type Language int

const (
	Chinese Language = iota
	English
	French
	German
	Italian
	Japanese
	Korean
	Portuguese
	Russian
	Spanish
	Unknown
)

// AllLanguages returns a sorted slice of all currently supported languages.
func AllLanguages() []Language {
	languages := make([]Language, amountOfSupportedLanguages())
	for i := 0; i < amountOfSupportedLanguages(); i++ {
		languages[i] = Language(i)
	}
	return languages
}

// AllSpokenLanguages returns a sorted slice of all supported languages
// that are not extinct but still spoken.
func AllSpokenLanguages() (languages []Language) {
	// All currently supported languages are spoken languages
	return AllLanguages()
}

// AllLanguagesWithArabicScript returns a sorted slice of all built-in
// languages supporting the Arabic script.
func AllLanguagesWithArabicScript() []Language {
	return allLanguagesWithScript(arabic)
}

// AllLanguagesWithCyrillicScript returns a sorted slice of all built-in
// languages supporting the Cyrillic script.
func AllLanguagesWithCyrillicScript() []Language {
	return allLanguagesWithScript(cyrillic)
}

// AllLanguagesWithDevanagariScript returns a sorted slice of all built-in
// languages supporting the Devanagari script.
func AllLanguagesWithDevanagariScript() []Language {
	return allLanguagesWithScript(devanagari)
}

// AllLanguagesWithLatinScript returns a sorted slice of all built-in
// languages supporting the Latin script.
func AllLanguagesWithLatinScript() []Language {
	return allLanguagesWithScript(latin)
}

// GetLanguageFromIsoCode639_1 returns the language for the given
// ISO 639-1 code enum value.
func GetLanguageFromIsoCode639_1(isoCode IsoCode639_1) Language {
	for _, language := range AllLanguages() {
		if language.IsoCode639_1() == isoCode {
			return language
		}
	}
	return Unknown
}

// GetLanguageFromIsoCode639_3 returns the language for the given
// ISO 639-3 code enum value.
func GetLanguageFromIsoCode639_3(isoCode IsoCode639_3) Language {
	for _, language := range AllLanguages() {
		if language.IsoCode639_3() == isoCode {
			return language
		}
	}
	return Unknown
}

func allLanguagesWithScript(script alphabet) (languages []Language) {
	for _, language := range AllLanguages() {
		if language.alphabets()[0] == script {
			languages = append(languages, language)
		}
	}
	return
}

func amountOfSupportedLanguages() int {
	return int(Spanish + 1)
}

// IsoCode639_1 returns a language's ISO 639-1 code.
func (language Language) IsoCode639_1() IsoCode639_1 {
	switch language {
	case Chinese:
		return ZH
	case English:
		return EN
	case French:
		return FR
	case German:
		return DE
	case Italian:
		return IT
	case Japanese:
		return JA
	case Korean:
		return KO
	case Portuguese:
		return PT
	case Russian:
		return RU
	case Spanish:
		return ES
	case Unknown:
		return UnknownIsoCode639_1
	default:
		return -1
	}
}

// IsoCode639_3 returns a language's ISO 639-3 code.
func (language Language) IsoCode639_3() IsoCode639_3 {
	switch language {
	case Chinese:
		return ZHO
	case English:
		return ENG
	case French:
		return FRA
	case German:
		return DEU
	case Italian:
		return ITA
	case Japanese:
		return JPN
	case Korean:
		return KOR
	case Portuguese:
		return POR
	case Russian:
		return RUS
	case Spanish:
		return SPA
	case Unknown:
		return UnknownIsoCode639_3
	default:
		return -1
	}
}

func (language Language) alphabets() []alphabet {
	switch language {
	case English, French, German, Italian, Portuguese, Spanish:
		return []alphabet{latin}
	case Russian:
		return []alphabet{cyrillic}
	case Chinese:
		return []alphabet{han}
	case Japanese:
		return []alphabet{hiragana, katakana, han}
	case Korean:
		return []alphabet{hangul}
	default:
		return []alphabet{}
	}
}

func (language Language) uniqueCharacters() string {
	switch language {
	case German:
		return "ß"
	case Spanish:
		return "¿¡"
	default:
		return ""
	}
}
