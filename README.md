
      
  <div id="readme" class="readme blob instapaper_body">
    <article class="markdown-body entry-content" itemprop="text"><h1><a id="user-content-tech-db-forum" class="anchor" aria-hidden="true" href="#tech-db-forum"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>tech-db-forum</h1>
<p>Тестовое задание для реализации проекта "Форумы" на курсе по базам данных в Технопарке Mail.ru (<a href="https://park.mail.ru" rel="nofollow">https://park.mail.ru</a>).</p>
<p>Суть задания заключается в реализации API к базе данных проекта «Форумы» по документации к этому API.</p>
<p>Таким образом, на входе:</p>
<ul>
<li>документация к API;</li>
</ul>
<p>На выходе:</p>
<ul>
<li>репозиторий, содержащий все необходимое для разворачивания сервиса в Docker-контейнере.</li>
</ul>
<h2><a id="user-content-документация-к-api" class="anchor" aria-hidden="true" href="#документация-к-api"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Документация к API</h2>
<p>Документация к API предоставлена в виде спецификации <a href="https://ru.wikipedia.org/wiki/OpenAPI_%28%D1%81%D0%BF%D0%B5%D1%86%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D1%8F%29" rel="nofollow">OpenAPI</a>: swagger.yml</p>
<p>Документацию можно читать как собственно в файле swagger.yml, так и через Swagger UI (там же есть возможность поиграться с запросами): <a href="https://tech-db-forum.bozaro.ru/" rel="nofollow">https://tech-db-forum.bozaro.ru/</a></p>
<h2><a id="user-content-требования-к-проекту" class="anchor" aria-hidden="true" href="#требования-к-проекту"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Требования к проекту</h2>
<p>Проект должен включать в себя все необходимое для разворачивания сервиса в Docker-контейнере.</p>
<p>При этом:</p>
<ul>
<li>файл для сборки Docker-контейнера должен называться Dockerfile и располагаться в корне репозитория;</li>
<li>реализуемое API должно быть доступно на 5000-ом порту по протоколу http;</li>
<li>допускается использовать любой язык программирования;</li>
<li>крайне не рекомендуется использовать ORM.</li>
</ul>
<p>Контейнер будет собираться из запускаться командами вида:</p>
<pre><code>docker build -t a.navrotskiy https://github.com/bozaro/tech-db-forum-server.git
docker run -p 5000:5000 --name a.navrotskiy -t a.navrotskiy
</code></pre>
<p>В качестве отправной точки можно посмотреть на примеры реализации более простого API на различных языках программирования: <a href="https://github.com/bozaro/tech-db-hello/">https://github.com/bozaro/tech-db-hello/</a></p>
<h2><a id="user-content-функциональное-тестирование" class="anchor" aria-hidden="true" href="#функциональное-тестирование"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Функциональное тестирование</h2>
<p>Корректность API будет проверяться при помощи автоматического функционального тестирования.</p>
<p>Методика тестирования:</p>
<ul>
<li>собирается Docker-контейнер из репозитория;</li>
<li>запускается Docker-контейнер;</li>
<li>запускается скрипт на Go, который будет проводить тестирование;</li>
<li>останавливается Docker-контейнер.</li>
</ul>
<p>Скомпилированные программы для тестирования можно скачать по ссылкам:</p>
<ul>
<li><a href="https://bozaro.github.io/tech-db-forum/darwin_amd64.zip" rel="nofollow">darwin_amd64.zip</a></li>
<li><a href="https://bozaro.github.io/tech-db-forum/linux_386.zip" rel="nofollow">linux_386.zip</a></li>
<li><a href="https://bozaro.github.io/tech-db-forum/linux_amd64.zip" rel="nofollow">linux_amd64.zip</a></li>
<li><a href="https://bozaro.github.io/tech-db-forum/windows_386.zip" rel="nofollow">windows_386.zip</a></li>
<li><a href="https://bozaro.github.io/tech-db-forum/windows_amd64.zip" rel="nofollow">windows_amd64.zip</a></li>
</ul>
<p>Для локальной сборки Go-скрипта достаточно выполнить команду:</p>
<pre><code>go get -u -v github.com/bozaro/tech-db-forum
go build github.com/bozaro/tech-db-forum
</code></pre>
<p>После этого в текущем каталоге будет создан исполняемый файл <code>tech-db-forum</code>.</p>
<h3><a id="user-content-запуск-функционального-тестирования" class="anchor" aria-hidden="true" href="#запуск-функционального-тестирования"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Запуск функционального тестирования</h3>
<p>Для запуска функционального тестирования нужно выполнить команду вида:</p>
<pre><code>./tech-db-forum func -u http://localhost:5000/api -r report.html
</code></pre>
<p>Поддерживаются следующие параметры:</p>
<table>
<thead>
<tr>
<th>Параметр</th>
<th>Описание</th>
</tr>
</thead>
<tbody>
<tr>
<td>-h, --help</td>
<td>Вывод списка поддерживаемых параметров</td>
</tr>
<tr>
<td>-u, --url[=http://localhost:5000/api]</td>
<td>Указание базовой URL тестируемого приложения</td>
</tr>
<tr>
<td>-k, --keep</td>
<td>Продолжить тестирование после первого упавшего теста</td>
</tr>
<tr>
<td>-t, --tests[=.*]</td>
<td>Маска запускаемых тестов (регулярное выражение)</td>
</tr>
<tr>
<td>-r, --report[=report.html]</td>
<td>Имя файла для детального отчета о функциональном тестировании</td>
</tr>
</tbody>
</table>
</article>
  </div>

