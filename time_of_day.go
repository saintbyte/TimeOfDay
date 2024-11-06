/*
Модуль занимается исключительно тем что принимает текущие время и делает чисто алгоритмически вывод какое время суток
День/Ночь AM/PM День/Вечер/Ночь/Утро. Делается исключительно для того чтобы подмешивать в промпты GPТ кроме текущего
времени еще и честь. Для того чтоб например ассистент говорил "Добрый вечер" не напрягая своих мозгов искючительно из
system prompt
*/
package TimeOfDay

import "time"

type TimeOfDay struct {
	PartOfDay1 string // День или ночь
	AmOrPM     string // AM или PM
	PartOfDay2 string // День или ночь или после полудня, вечер , ночь
}

func Parse(inputTime time.Time) TimeOfDay {
	// Функция получает на вход время, стоит позабоиться чтоб подкрутить там часовой пояс если надо и выдает структуру с
	// Результатом в котором проставлены соотвествующие флаги.
	hour := inputTime.Hour()
	var dayNight string
	var partOfDay2 string
	var ampm string

	switch {
	case hour >= 5 && hour < 12:
		dayNight = Day
		partOfDay2 = Morning
	case hour >= 12 && hour < 18:
		dayNight = Day
		partOfDay2 = Afternoon
	case hour >= 18 && hour < 21:
		dayNight = Day
		partOfDay2 = Evening
	default:
		dayNight = Night
		partOfDay2 = Night
	}
	if hour >= 12 {
		ampm = PM
	} else {
		ampm = AM
	}
	return TimeOfDay{
		dayNight,
		ampm,
		partOfDay2,
	}
}
