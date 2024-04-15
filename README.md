<h1>Ozon Link reducer</h1>
<p>Решение тестового задания на позицию Golang стажёр</p>
<hr />
<h2>Задание</h2>
<p><img src="![image](https://github.com/Serveresty/OZONTestCaseLinks/assets/99687697/d8484c38-7b99-430b-bd16-b2439b8cd6d1)
" /></p>
<hr />
<h2>Для начала</h2>
<p>Для запуска сервиса напрямую(не docker образ): 
<br>
<pre>
```bash
go run ./cmd/main.go -storage="in-memory"
```
</pre>
<br>
P.S. В параметр storage можно записать "postgresql" или "in-memory", иначе выбрасываем.
</p>
<p>Для запуска сервиса через docker:
<br>
```bash
docker-compose --env-file configs/.env run -p 8080:8080 app ./build/main --STORAGE in-memory
```
<br>
P.S. Указывается .env файл и порты сервиса, т.к. какой-то трабл с этим...
</p>
<hr />
<h2>Описание методов</h2>
<p>Все методы для работы с базой данных и кешем не учитывают вхождения одних и тех же данных(это сервисы только для доступа к базе). Всё это учитывается методом database.ReduceLink() и database.OriginalLink.</p>
<p>На весь сервис ОДИН обработчик, который валидирует "GET" и "POST" запросы. Чтобы создать сокращённую ссылку необходимо отправить JSON с ссылкой(БЕЗ КАКИХ ЛИБО СКОБОК!!!!!)
<br>
Пример отправки POST запроса:
<br>
<img src="![image](https://github.com/Serveresty/OZONTestCaseLinks/assets/99687697/08bc31d9-e3be-4e83-a97d-c41011340a10)
" />
<br>
Для получения изначальной ссылки отправляется запрос через укороченную ссылку.
</p>
<hr />
<h2>Пример</h2>
<p>Отправляем POST запрос с нашей ссылкой в теле: "https://ozon.ru/" и получаем ссылку вида: "host:port/AfZGd_ZTNg"
<br>
Для получения исходной ссылки отправляем GET запрос: "host:port/AfZGd_ZTNg" и получаем "https://ozon.ru/"
</p>
