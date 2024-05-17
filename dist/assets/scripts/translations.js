// Cookie retrieval by name
function getCookie(name) {
    let match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
    if (match) return match[2];
}

// Cookie setting by name, value and lifetime
function setCookie(name, value, days) {
    let expires = "";
    if (days) {
        let date = new Date();
        date.setTime(date.getTime() + (days*24*60*60*1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + value + expires + "; SameSite=strict; path=/";
}

// Language setting based on user's choice
function setLanguage() {
    const selectedLanguage = document.getElementById('language').value;
    setCookie('language', selectedLanguage, 7);
    loadTranslations(selectedLanguage);
}

// Change translation file
function loadTranslations(langCode) {
    fetch(`assets/translations/${langCode}.yaml`)
        .then(response => response.text())
        .then(data => {
            const translations = jsyaml.load(data);
            document.getElementsByTagName('html')[0].setAttribute('lang', translations.lang);
            document.getElementById('title').textContent = translations.title;
            document.getElementsByTagName('h1')[0].textContent = translations.title;
            document.getElementById('language-selector').getElementsByTagName('label')[0].textContent = translations.choose_language;
        })
        .catch(error => console.error('Error loading translation:', error));
}

// Load default language or the one defined in the cookie
document.addEventListener('DOMContentLoaded', (event) => {
    const savedLanguage = getCookie('language') || (navigator.language || navigator.userLanguage).split('-')[0];
    document.getElementById('language').value = savedLanguage;
    loadTranslations(savedLanguage);
});
