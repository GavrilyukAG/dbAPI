
      
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

    </div>

  

  <details class="details-reset details-overlay details-overlay-dark">
    <summary data-hotkey="l" aria-label="Jump to line"></summary>
    <details-dialog class="Box Box--overlay d-flex flex-column anim-fade-in fast linejump" aria-label="Jump to line">
      <!-- '"` --><!-- </textarea></xmp> --></option></form><form class="js-jump-to-line-form Box-body d-flex" action="" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
        <input class="form-control flex-auto mr-3 linejump-input js-jump-to-line-field" type="text" placeholder="Jump to line&hellip;" aria-label="Jump to line" autofocus>
        <button type="submit" class="btn" data-close-dialog>Go</button>
</form>    </details-dialog>
  </details>


  </div>
  <div class="modal-backdrop js-touch-events"></div>
</div>

    </div>
  </div>

  </div>

        
<div class="footer container-lg px-3" role="contentinfo">
  <div class="position-relative d-flex flex-justify-between pt-6 pb-2 mt-6 f6 text-gray border-top border-gray-light ">
    <ul class="list-style-none d-flex flex-wrap ">
      <li class="mr-3">&copy; 2018 <span title="0.13012s from unicorn-546f974758-sxvmx">GitHub</span>, Inc.</li>
        <li class="mr-3"><a data-ga-click="Footer, go to terms, text:terms" href="https://github.com/site/terms">Terms</a></li>
        <li class="mr-3"><a data-ga-click="Footer, go to privacy, text:privacy" href="https://github.com/site/privacy">Privacy</a></li>
        <li class="mr-3"><a href="https://help.github.com/articles/github-security/" data-ga-click="Footer, go to security, text:security">Security</a></li>
        <li class="mr-3"><a href="https://status.github.com/" data-ga-click="Footer, go to status, text:status">Status</a></li>
        <li><a data-ga-click="Footer, go to help, text:help" href="https://help.github.com">Help</a></li>
    </ul>

    <a aria-label="Homepage" title="GitHub" class="footer-octicon mr-lg-4" href="https://github.com">
      <svg height="24" class="octicon octicon-mark-github" viewBox="0 0 16 16" version="1.1" width="24" aria-hidden="true"><path fill-rule="evenodd" d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"/></svg>
</a>
   <ul class="list-style-none d-flex flex-wrap ">
        <li class="mr-3"><a data-ga-click="Footer, go to contact, text:contact" href="https://github.com/contact">Contact GitHub</a></li>
        <li class="mr-3"><a href="https://github.com/pricing" data-ga-click="Footer, go to Pricing, text:Pricing">Pricing</a></li>
      <li class="mr-3"><a href="https://developer.github.com" data-ga-click="Footer, go to api, text:api">API</a></li>
      <li class="mr-3"><a href="https://training.github.com" data-ga-click="Footer, go to training, text:training">Training</a></li>
        <li class="mr-3"><a href="https://blog.github.com" data-ga-click="Footer, go to blog, text:blog">Blog</a></li>
        <li><a data-ga-click="Footer, go to about, text:about" href="https://github.com/about">About</a></li>

    </ul>
  </div>
  <div class="d-flex flex-justify-center pb-6">
    <span class="f6 text-gray-light"></span>
  </div>
</div>



  <div id="ajax-error-message" class="ajax-error-message flash flash-error">
    <svg class="octicon octicon-alert" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"/></svg>
    <button type="button" class="flash-close js-ajax-error-dismiss" aria-label="Dismiss error">
      <svg class="octicon octicon-x" viewBox="0 0 12 16" version="1.1" width="12" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z"/></svg>
    </button>
    You can’t perform that action at this time.
  </div>


    <script crossorigin="anonymous" integrity="sha512-BlCeoZU+kjn7xucWZBcl0n4Bn0P8dE19/sUfLHOxySQnsoy3ufEzapurMbZWSlwab5KGfnp1X5ipJvUDMLroqw==" type="application/javascript" src="https://assets-cdn.github.com/assets/compat-0f98b12d09f3ba331eef956ab02996e3.js"></script>
    <script crossorigin="anonymous" integrity="sha512-9E3sHSwrJGUiIDmf7IURDLz9Ob2E/W2JRykVfAAaL0cvLrwnjoUYAeLu4XQ4j9RTPLclLDLAeRhbGR9bwxTiyg==" type="application/javascript" src="https://assets-cdn.github.com/assets/frameworks-8b43903482b2c45904929cdac0c62a0e.js"></script>
    
    <script crossorigin="anonymous" async="async" integrity="sha512-Uv7ykUpMT7bI4W390Z67EVhYKVaMZz1C4ADfiPVrtrdF6PCquKN7AHRzrrTaikkaL3jDnY/pYalw4ZF+5Bj90Q==" type="application/javascript" src="https://assets-cdn.github.com/assets/github-b8e8624c21ec3f2091551b4c77d879e0.js"></script>
    
    
    
  <div class="js-stale-session-flash stale-session-flash flash flash-warn flash-banner d-none">
    <svg class="octicon octicon-alert" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"/></svg>
    <span class="signed-in-tab-flash">You signed in with another tab or window. <a href="">Reload</a> to refresh your session.</span>
    <span class="signed-out-tab-flash">You signed out in another tab or window. <a href="">Reload</a> to refresh your session.</span>
  </div>
  <div class="facebox" id="facebox" style="display:none;">
  <div class="facebox-popup">
    <div class="facebox-content" role="dialog" aria-labelledby="facebox-header" aria-describedby="facebox-description">
    </div>
    <button type="button" class="facebox-close js-facebox-close" aria-label="Close modal">
      <svg class="octicon octicon-x" viewBox="0 0 12 16" version="1.1" width="12" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z"/></svg>
    </button>
  </div>
</div>

  <template id="site-details-dialog">
  <details class="details-reset details-overlay details-overlay-dark lh-default text-gray-dark" open>
    <summary aria-haspopup="dialog" aria-label="Close dialog"></summary>
    <details-dialog class="Box Box--overlay d-flex flex-column anim-fade-in fast">
      <button class="Box-btn-octicon m-0 btn-octicon position-absolute right-0 top-0" type="button" aria-label="Close dialog" data-close-dialog>
        <svg class="octicon octicon-x" viewBox="0 0 12 16" version="1.1" width="12" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z"/></svg>
      </button>
      <div class="octocat-spinner my-6 js-details-dialog-spinner"></div>
    </details-dialog>
  </details>
</template>

  <div class="Popover js-hovercard-content position-absolute" style="display: none; outline: none;" tabindex="0">
  <div class="Popover-message Popover-message--bottom-left Popover-message--large Box box-shadow-large" style="width:360px;">
  </div>
</div>

<div id="hovercard-aria-description" class="sr-only">
  Press h to open a hovercard with more details.
</div>


  </body>
</html>

