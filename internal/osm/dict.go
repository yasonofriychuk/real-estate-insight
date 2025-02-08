package osm

type ObjType string

const (
	Hospital     ObjType = "hospital"
	Sport        ObjType = "sport"
	Shops        ObjType = "shops"
	Kindergarten ObjType = "kindergarten"
	BusStop      ObjType = "bus_stop"
)

var (
	//nolint: unused
	tagsMap = map[string]struct{}{
		// Образование
		"amenity=school":       {}, // Школы
		"amenity=kindergarten": {}, // Детские садики
		"amenity=university":   {}, // Университеты и колледжи

		// Здравоохранение
		"amenity=hospital": {}, // Больницы
		"amenity=clinic":   {}, // Поликлиники
		"amenity=pharmacy": {}, // Аптеки
		"amenity=doctors":  {}, // Врачебные кабинеты, медицинские офисы
		"amenity=dentist":  {}, // Стоматологические кабинеты

		// Магазины и бытовые услуги
		"shop=supermarket":      {}, // Супермаркеты
		"shop=convenience":      {}, // Продуктовые магазины
		"shop=department_store": {}, // Универмаги
		"shop=hardware":         {}, // Хозяйственные магазины
		"shop=mall":             {}, // Торговые центры

		// Спортивные объекты и зоны отдыха
		"leisure=sports_centre":  {}, // Спортивные центры
		"leisure=stadium":        {}, // Стадионы
		"leisure=pitch":          {}, // Спортивные площадки
		"leisure=fitness_centre": {}, // Фитнес-центры
		"leisure=playground":     {}, // Детские площадки
		"leisure=park":           {}, // Парки, зоны отдыха

		// Общественный транспорт
		"public_transport=stop_position": {}, // Остановки общественного транспорта
		"public_transport=station":       {}, // Станции поездов, метро
		"amenity=bus_station":            {}, // Автобусные станции
		"railway=subway_entrance":        {}, // Входы в метро

		// Культурные объекты
		"amenity=library": {}, // Библиотеки
		"amenity=theatre": {}, // Театры
		"amenity=cinema":  {}, // Кинотеатры
		"amenity=museum":  {}, // Музеи

		// Прочие важные объекты
		"amenity=bank":         {}, // Банки
		"amenity=atm":          {}, // Банкоматы
		"amenity=post_office":  {}, // Почтовые отделения
		"amenity=police":       {}, // Отделения полиции
		"amenity=fire_station": {}, // Пожарные станции
		"amenity=parking":      {}, // Парковки
	}
)
