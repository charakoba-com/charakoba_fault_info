<!doctype html>
<html lang="ja">
  <head>
    <meta charset="UTF-8"/>
    <title>Information::charakoba.com</title>
    {% include "include.tpl" %}
  </head>
  <body>
    {% include "navbar.tpl" %}
    <main class="container">
      <div class="row">
        <div class="col s10 offset-s1">
          <div class="card">
            <div class="card-content green-text text-darken-4">
              <span class="card-title">{{ info.type }} :: {{ info.service.decode('utf-8') }}</span>
              <p>
                {{ info.begin }} - {{ info.end or '' }}
              </p>
              <p>
                {{ info.detail.decode('utf-8') }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </main>
  </body>
</html>
