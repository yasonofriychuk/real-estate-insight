mapboxgl.accessToken = 'pk.eyJ1Ijoic3ZjLW9rdGEtbWFwYm94LXN0YWZmLWFjY2VzcyIsImEiOiJjbG5sMnExa3kxNTJtMmtsODJld24yNGJlIn0.RQ4CHchAYPJQZSiUJ0O3VQ';

let currentPage = 1;
const perPage = 10;
let allJks = [];

const jkList = document.getElementById("jks-list");
const prevPageBtn = document.getElementById("prev-page");
const nextPageBtn = document.getElementById("next-page");
const pageInfo = document.getElementById("page-info");
const jkDetails = document.getElementById("jk-details");
const jkDetailsContent = document.getElementById("jk-details-content");
const backBtn = document.getElementById("back-btn");

const searchInput = document.getElementById("search-input");
const filterBtn = document.getElementById("filter-btn");
const filterModal = document.getElementById("filter-modal");
const closeFilterBtn = document.getElementById("close-filter-btn");

const map = new mapboxgl.Map({
    container: 'map',
    center: [73.423043, 61.258726],
    zoom: 12,
    style: 'mapbox://styles/mapbox/standard',
    hash: true,
});

document.addEventListener("DOMContentLoaded", () => {
    // Возврат к списку ЖК
    backBtn.addEventListener("click", () => {
        jkList.style.display = "block";
        pagination.style.display = "flex";
        jkDetails.style.display = "none";

        if (window.infrastructureMarkers) {
            window.infrastructureMarkers.forEach(marker => marker.remove());
            window.infrastructureMarkers = [];
        }

        fetchAndDisplayMarkers();
    });

    // Пагинация
    prevPageBtn.addEventListener("click", () => {
        if (currentPage > 1) {
            currentPage--;
            renderJks();
        }
    });

    nextPageBtn.addEventListener("click", () => {
        if (currentPage * perPage < allJks.length) {
            currentPage++;
            renderJks();
        }
    });

    searchInput.addEventListener("input", () => fetchJks(searchInput.value));

    filterBtn.addEventListener("click", () => {
        filterModal.classList.remove("hidden");
    });

    closeFilterBtn.addEventListener("click", () => {
        filterModal.classList.add("hidden");
    });
});

// Запуск при загрузке страницы
window.onload = function () {
    const lon = getQueryParam('lon');
    const lat = getQueryParam('lat');

    if (lon && lat) {
        const position = [parseFloat(lon), parseFloat(lat)];

        map.flyTo({
            center: position,
            zoom: 15,
            speed: 0.1,
            duration: 2000,
        });

        new mapboxgl.Marker()
            .setLngLat(position)
            .addTo(map);
    }

    fetchAndDisplayMarkers();
    fetchJks("");
};

map.on('moveend', () => {
    fetchAndDisplayMarkers();
});