async function fetchAndDisplayMarkers() {
    try {
        // Получаем границы видимой области карты
        const bounds = map.getBounds();
        const sw = bounds.getSouthWest();
        const ne = bounds.getNorthEast();

        const response = await fetch('/api/v1/developments/search/filter', {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                board: { topLeftLon: ne.lng, topLeftLat: ne.lat, bottomRightLon: sw.lng, bottomRightLat: sw.lat }
            }),
        });

        if (!response.ok) throw new Error('Ошибка сети');
        const data = await response.json();

        if (window.markers) {
            window.markers.forEach(marker => marker.remove());
        }
        window.markers = [];

        data.developments.forEach((item) => {
            const { id, name, coords, imageUrl } = item;

            const el = document.createElement('div');
            el.className = 'marker';
            el.style.backgroundImage = 'url("https://cdn-icons-png.flaticon.com/512/684/684908.png")';
            el.style.backgroundSize = 'cover';
            el.style.width = '30px';
            el.style.height = '30px';

            const popup = new mapboxgl.Popup({ offset: 25 }).setHTML(`
                <h4>${name}</h4>
                <img src="${imageUrl}" style="width: 100px; height: auto;">
                <button id="infra-btn-${id}">Показать инфраструктуру</button>
            `);

            const marker = new mapboxgl.Marker(el)
                .setLngLat([coords.lon, coords.lat])
                .setPopup(popup)
                .addTo(map);

            window.markers.push(marker);

            marker.getElement().addEventListener('click', () => {
                selectJk(item);
            });
        });
    } catch (error) {
        console.error('Ошибка:', error);
    }
}

// Функция загрузки и отображения инфраструктурных объектов
async function fetchAndDisplayInfrastructure(developmentId) {
    try {
        // Удаляем старые инфраструктурные маркеры
        if (window.infrastructureMarkers) {
            window.infrastructureMarkers.forEach(marker => marker.remove());
        }
        window.infrastructureMarkers = [];

        const url = new URL('/api/v1/infrastructure/radius', window.location.origin);
        url.searchParams.set('developmentId', developmentId);
        url.searchParams.set('radius', "1000");

        const response = await fetch(url);
        if (!response.ok) {
            throw new Error('Ошибка сети при загрузке инфраструктуры');
        }
        const infraData = await response.json();

        infraData.forEach((infra) => {
            const { lon, lat, name, objType } = infra.coords;

            // Создание кастомного маркера для инфраструктуры
            const el = document.createElement('div');
            el.className = 'infra-marker';
            el.style.backgroundImage = 'url("https://cdn-icons-png.flaticon.com/512/684/684913.png")';
            el.style.backgroundSize = 'cover';
            el.style.width = '20px';
            el.style.height = '20px';
            el.style.borderRadius = '50%';

            const popup = new mapboxgl.Popup({ offset: 15 }).setText(`${name || objType}`);

            const marker = new mapboxgl.Marker(el)
                .setLngLat([lon, lat])
                .setPopup(popup)
                .addTo(map);

            window.infrastructureMarkers.push(marker);
        });

    } catch (error) {
        console.error('Ошибка загрузки инфраструктуры:', error);
    }
}

// Запуск функции обновления маркеров при бездействии карты
// map.on('idle', fetchAndDisplayMarkers);

// Функция получения параметра из URL
function getQueryParam(param) {
    return new URLSearchParams(window.location.search).get(param);
}