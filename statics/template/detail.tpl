<!doctype html>
<html lang="ja">
  <head>
    <meta charset="UTF-8"/>
    <title>Information::charakoba.com</title>
    <!--Import Google Icon Font-->
    <link href="http://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!--Import materialize.css-->
    <link type="text/css" rel="stylesheet" href="static_file/css/materialize.min.css"  media="screen,projection"/>
    <!--Import jQuery before materialize.js-->
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script type="text/javascript" src="statics/js/materialize.min.js"></script>

    <!--Let browser know website is optimized for mobile-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  </head>
  <body>
    <nav class="nav-wrapper green darken-3">
      <a class="brand-logo" href="">Information::charakoba.com</a>
      <ul id="nav-mobile" class="right">
        <li><a href="admin"></a></li>
      </ul>
    </nav>
    <main class="container">
      <div class="row">
        <div class="col s10 offset-s1">
          <div class="card">
            <div class="card-content green-text text-darken-4">
              <span class="card-title">{{ info.type }} :: {{ info.service }}</span>
              <p>
                {{ info.begin }} - {{ info.end or '' }}
              </p>
              <p>
                {{ info.detail }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </main>
  </body>
</html>
