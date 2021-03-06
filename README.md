# Geolocation WEB API 🌍
Работа с открытым браузерным интерфейсом Web Geolocation API, которое позволяет определить географическое положение и пользователя.

- [Презентация с вебинара на SkillFactory](https://docs.google.com/presentation/d/1Y9Gfzxn_IwuJC3u2c-xPgOZ90jT6Py8hIaTBkrMlILY/edit?usp=sharing)
- Код работы с Geolocation API располагается в index.html. Запущенный скрипт позволит отобразить искаженные (специально) геоданные или обработать исключения.

## Что такое WEB API? ⬇️

WEB API расширяет возможности браузера, стандартизирует и упрощает ваш код.

Приставка API говорит о том, что это открытый интерфейс, а WEB говорит о том, что это интерфейс для веб приложений. Реализовать доступный открытый веб интерфейс можно на стороне клиента и сервера.  

В терминологии WEB API, весь веб интерфейс подразделяется на Browser API , Server API, Third Party API (сторонний интерфейс). Мы, как frontend разработчики, говорим о Browser API.

WEB API на клиенте расширяет возможности веб-браузера или другой HTTP клиент. Изначально интерфейс был в форме встроенных расширений браузера, однако большинство нацелены на стандартизацию и привязку к JavaScript.

Немного о возможностях открытого интерфейса. К примеру возьмем API немалоизвестной платформы для рекрутинга - headhunter. Мы можем произвести запрос, указав в качестве параметров регион, готовность к переезду, тип занятости и получить ответ с данными поиска в Базе. Но тут мы не строим аналитическое SQL - хранилище.

Все браузеры имеют комплекс интерфейсов, которые помогут решить сложные задачи на более высоком уровне. С помощью Geolocation API мы сможем получить координаты широты и долготы местоположения нашего браузера. Смотрите код проекта.

## Что такое Geolocation Web API? ⬇️

Geolocation API используется для определения географического положения пользователя. Однако с точки зрения безопасности мы не сможем автоматически отследить положение человека, который посетил ваш сайт. Обязательно появится модальное окно с запросом на вашу геолокацию. (вспоминайте любую карту на вашем мобильном телефоне или в браузере)

Геолокация работает гораздо точнее на устройствах с GPS, к примеру на смартфоне. Поддерживается многими популярными браузерами, такими как Google Chrome, Safari, Mozilla и другими.

Начиная с версии Chrome 50 API, запросы геолокации не будут работать, если сайт размещен в незащищенном источнике. Только в безопасном контексте, таком как HTTPS.
