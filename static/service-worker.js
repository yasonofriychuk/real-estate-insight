if ('serviceWorker' in navigator) {
navigator.serviceWorker.register('/service-worker.js')
    .then(function(registration) {
        console.log('ServiceWorker зарегистрирован:', registration);
    })
    .catch(function(error) {
        console.log('Ошибка регистрации ServiceWorker:', error);
    });
}