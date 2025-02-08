// Функция загрузки ЖК
async function fetchJks(searchQuery = "") {
    try {
        const response = await fetch('/api/v1/developments/search/filter', {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                searchQuery,
                pagination: { page: currentPage, perPage },
            }),
        });

        if (!response.ok) throw new Error("Ошибка загрузки ЖК");

        const data = await response.json();
        allJks = data.developments;
        totalJks = data.meta.total;

        renderJks();
    } catch (error) {
        console.error(error);
    }
}

// Обновление пагинации
function updatePagination() {
    pageInfo.innerText = `Страница ${currentPage}`;
    prevPageBtn.disabled = currentPage === 1;
    nextPageBtn.disabled = currentPage * perPage >= totalJks;
}

// Отрисовка ЖК
function renderJks() {
    jkList.innerHTML = "";
    allJks.forEach(jk => {
        const card = document.createElement("div");
        card.className = "jk-card";
        card.innerHTML = `
            <img src="${jk.imageUrl}" alt="${jk.name}">
            <h4>${jk.name}</h4>
            <p>${jk.description || "Нет описания"}</p>
        `;
        card.addEventListener("click", () => selectJk(jk));
        jkList.appendChild(card);
    });

    updatePagination();
}

// Выбор ЖК (убирает остальные ЖК, зумирует, загружает инфраструктуру)
function selectJk(jk) {
    jkList.style.display = "none";
    pagination.style.display = "none";
    jkDetails.style.display = "block";
    jkDetailsContent.innerHTML = `
        <h2>${jk.name}</h2>
        <img src="${jk.imageUrl}" style="width: 100%;">
        <p>${jk.description || "Нет описания"}</p>
    `;

    map.flyTo({ center: [jk.coords.lon, jk.coords.lat], zoom: 15 });
    fetchAndDisplayInfrastructure(jk.id);
}

// Функция показа детальной информации о ЖК
function showJkDetails(jk) {
    jkDetailsContent.innerHTML = `
        <img src="${jk.imageUrl}" alt="${jk.name}" style="width: 100%; border-radius: 5px;">
        <h2>${jk.name}</h2>
        <p>${jk.description || "Нет описания"}</p>
    `;

    jkList.style.display = "none";
    pagination.style.display = "none";
    jkDetails.style.display = "block";
}
