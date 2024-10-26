package pbf_scanner

import (
	"fmt"
	"io"

	"github.com/qedus/osmpbf"
)

var (
	allowedGroupTags = map[string]struct{}{
		// Общие
		"amenity":          {}, // Важные объекты инфраструктуры: больницы, школы, детские сады, кафе и т.д.
		"shop":             {}, // Магазины, торговые центры
		"public_transport": {}, // Остановки общественного транспорта, метро
		"leisure":          {}, // Зоны отдыха: парки, спортивные площадки
		"healthcare":       {}, // Медицинские учреждения, клиники
		"railway":          {}, // Метро
	}

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

type FeatureScanner struct {
	decoder    *osmpbf.Decoder
	nodeFilter func(*osmpbf.Node) bool
}

func New(osmPbf io.Reader, countGoroutines int) (*FeatureScanner, error) {
	decoder := osmpbf.NewDecoder(osmPbf)
	if err := decoder.Start(countGoroutines); err != nil {
		return nil, fmt.Errorf("could not start osmpbf decoder: %w", err)
	}

	return &FeatureScanner{
		decoder: decoder,
		nodeFilter: func(node *osmpbf.Node) bool {
			for _, tags := range []map[string]struct{}{allowedGroupTags, tagsMap} {
				for t, _ := range tags {
					if _, ok := node.Tags[t]; ok {
						return true
					}
				}
			}
			return false
		},
	}, nil
}

func (s *FeatureScanner) Next() (osmpbf.Node, error) {
	for {
		node, err := s.decoder.Decode()
		if err != nil {
			return osmpbf.Node{}, io.EOF
		}

		n, ok := node.(*osmpbf.Node)
		if !ok || !s.nodeFilter(n) {
			continue
		}

		return *n, nil
	}
}
